package account

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"testing"
	"time"

	accountdomain "github.com/chenyme/grok2api/backend/internal/domain/account"
	"github.com/chenyme/grok2api/backend/internal/infra/persistence/relational"
	"github.com/chenyme/grok2api/backend/internal/infra/provider"
	cliprovider "github.com/chenyme/grok2api/backend/internal/infra/provider/cli"
	"github.com/chenyme/grok2api/backend/internal/infra/provider/web"
	"github.com/chenyme/grok2api/backend/internal/infra/security"
)

func TestExportCredentialsRoundTripsImportFormat(t *testing.T) {
	ctx := context.Background()
	database, err := relational.OpenSQLite(ctx, filepath.Join(t.TempDir(), "export.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	if err := database.InitializeSchema(ctx); err != nil {
		t.Fatal(err)
	}
	cipher, err := security.NewCipher(base64.StdEncoding.EncodeToString(make([]byte, 32)))
	if err != nil {
		t.Fatal(err)
	}
	accessToken, err := cipher.Encrypt("access-token")
	if err != nil {
		t.Fatal(err)
	}
	refreshToken, err := cipher.Encrypt("refresh-token")
	if err != nil {
		t.Fatal(err)
	}
	expiresAt := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
	repository := relational.NewAccountRepository(database)
	if _, _, err := repository.UpsertByIdentity(ctx, accountdomain.Credential{
		Provider: accountdomain.ProviderBuild, Name: "primary", Email: "user@example.com", UserID: "user-1",
		SourceKey: "export-test", OIDCClientID: "client-1", EncryptedAccessToken: accessToken,
		EncryptedRefreshToken: refreshToken, ExpiresAt: expiresAt, Enabled: false,
		AuthStatus: accountdomain.AuthStatusActive, Priority: 1, MaxConcurrent: 8,
	}); err != nil {
		t.Fatal(err)
	}
	adapter := cliprovider.NewAdapter(cliprovider.Config{}, cipher)
	service := NewService(repository, nil, nil, nil, provider.NewRegistry(adapter), cipher, nil)

	result, err := service.ExportCredentials(ctx)
	if err != nil {
		t.Fatal(err)
	}
	values, err := adapter.ParseImportedCredentials(result.Data)
	if err != nil {
		t.Fatal(err)
	}
	if result.Count != 1 || len(values) != 1 {
		t.Fatalf("export count = %d, imported values = %d", result.Count, len(values))
	}
	value := values[0]
	if value.Name != "primary" || value.Email != "user@example.com" || value.UserID != "user-1" || value.OIDCClientID != "client-1" || value.AccessToken != "access-token" || value.RefreshToken != "refresh-token" || !value.ExpiresAt.Equal(expiresAt) {
		t.Fatalf("round-trip credential = %#v", value)
	}
	progress := make([][2]int, 0, 2)
	if _, err := service.ImportCredentialsWithProgress(ctx, result.Data, nil, func(completed, total int) error {
		progress = append(progress, [2]int{completed, total})
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if len(progress) != 2 || progress[0] != [2]int{0, 1} || progress[1] != [2]int{1, 1} {
		t.Fatalf("import progress = %#v", progress)
	}

	multiProgress := make([][2]int, 0, 3)
	multiResult, err := service.ImportCredentialDocumentsWithProgress(ctx, [][]byte{
		result.Data,
		result.Data,
		[]byte(`{"provider":"grok_build","name":"secondary","access_token":"second-access","refresh_token":"second-refresh","user_id":"user-2"}`),
	}, nil, func(completed, total int) error {
		multiProgress = append(multiProgress, [2]int{completed, total})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if multiResult.Created != 1 || multiResult.Updated != 1 {
		t.Fatalf("multi-file import result = %#v", multiResult)
	}
	if len(multiProgress) != 3 || multiProgress[0] != [2]int{0, 2} || multiProgress[2] != [2]int{2, 2} {
		t.Fatalf("multi-file import progress = %#v", multiProgress)
	}
}

func TestExportWebCredentialsRoundTripsImportFormat(t *testing.T) {
	ctx := context.Background()
	sourceDB, err := relational.OpenSQLite(ctx, filepath.Join(t.TempDir(), "source.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer sourceDB.Close()
	targetDB, err := relational.OpenSQLite(ctx, filepath.Join(t.TempDir(), "target.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer targetDB.Close()
	for _, database := range []*relational.Database{sourceDB, targetDB} {
		if err := database.InitializeSchema(ctx); err != nil {
			t.Fatal(err)
		}
	}
	cipher, err := security.NewCipher(base64.StdEncoding.EncodeToString(make([]byte, 32)))
	if err != nil {
		t.Fatal(err)
	}
	ssoCiphertext, err := cipher.Encrypt("sso-token")
	if err != nil {
		t.Fatal(err)
	}
	cookieCiphertext, err := cipher.Encrypt("cf_clearance=clear; __cf_bm=bm")
	if err != nil {
		t.Fatal(err)
	}
	sourceRepository := relational.NewAccountRepository(sourceDB)
	if _, _, err := sourceRepository.UpsertByIdentity(ctx, accountdomain.Credential{
		Provider: accountdomain.ProviderWeb, AuthType: accountdomain.AuthTypeSSO,
		Name: "web primary", SourceKey: "sso:web-export-test", WebTier: accountdomain.WebTierSuper,
		EncryptedAccessToken: ssoCiphertext, EncryptedCloudflareCookie: cookieCiphertext,
		Enabled: true, AuthStatus: accountdomain.AuthStatusActive,
	}); err != nil {
		t.Fatal(err)
	}
	sourceService := NewService(sourceRepository, nil, nil, nil, provider.NewRegistry(&web.Adapter{}), cipher, nil)
	targetRepository := relational.NewAccountRepository(targetDB)
	targetService := NewService(targetRepository, nil, nil, nil, provider.NewRegistry(&web.Adapter{}), cipher, nil)

	result, err := sourceService.ExportWebCredentials(ctx)
	if err != nil {
		t.Fatal(err)
	}
	targetResult, err := targetService.ImportWebCredentialDocumentsWithProgress(ctx, [][]byte{result.Data}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Count != 1 || targetResult.Created != 1 || len(targetResult.AccountIDs) != 1 {
		t.Fatalf("export = %#v, import = %#v", result, targetResult)
	}
	value, err := targetRepository.Get(ctx, targetResult.AccountIDs[0])
	if err != nil {
		t.Fatal(err)
	}
	ssoToken, err := cipher.Decrypt(value.EncryptedAccessToken)
	if err != nil {
		t.Fatal(err)
	}
	cookies, err := cipher.Decrypt(value.EncryptedCloudflareCookie)
	if err != nil {
		t.Fatal(err)
	}
	if value.Name != "web primary" || value.WebTier != accountdomain.WebTierSuper || ssoToken != "sso-token" || cookies != "cf_clearance=clear; __cf_bm=bm" {
		t.Fatalf("round-trip credential = %#v, sso = %q, cookies = %q", value, ssoToken, cookies)
	}
}

func TestExportWebCredentialsEmpty(t *testing.T) {
	ctx := context.Background()
	database, err := relational.OpenSQLite(ctx, filepath.Join(t.TempDir(), "empty-web-export.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	if err := database.InitializeSchema(ctx); err != nil {
		t.Fatal(err)
	}
	cipher, err := security.NewCipher(base64.StdEncoding.EncodeToString(make([]byte, 32)))
	if err != nil {
		t.Fatal(err)
	}
	service := NewService(relational.NewAccountRepository(database), nil, nil, nil, provider.NewRegistry(&web.Adapter{}), cipher, nil)
	if _, err := service.ExportWebCredentials(ctx); !errors.Is(err, ErrExportEmpty) {
		t.Fatalf("error = %v, want ErrExportEmpty", err)
	}
}

func TestExportCLIProxyCredentialsSingleJSON(t *testing.T) {
	ctx := context.Background()
	database, err := relational.OpenSQLite(ctx, filepath.Join(t.TempDir(), "cliproxy-export.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	if err := database.InitializeSchema(ctx); err != nil {
		t.Fatal(err)
	}
	cipher, err := security.NewCipher(base64.StdEncoding.EncodeToString(make([]byte, 32)))
	if err != nil {
		t.Fatal(err)
	}
	accessToken, err := cipher.Encrypt("access-token")
	if err != nil {
		t.Fatal(err)
	}
	refreshToken, err := cipher.Encrypt("refresh-token")
	if err != nil {
		t.Fatal(err)
	}
	expiresAt := time.Date(2026, 7, 18, 6, 30, 55, 0, time.UTC)
	repository := relational.NewAccountRepository(database)
	if _, _, err := repository.UpsertByIdentity(ctx, accountdomain.Credential{
		Provider: accountdomain.ProviderBuild, Name: "primary", Email: "user@example.com", UserID: "user-1",
		SourceKey: "cliproxy-export-test", OIDCClientID: "client-1", EncryptedAccessToken: accessToken,
		EncryptedRefreshToken: refreshToken, ExpiresAt: expiresAt, Enabled: true,
		AuthStatus: accountdomain.AuthStatusActive, Priority: 0, MaxConcurrent: 8,
	}); err != nil {
		t.Fatal(err)
	}
	adapter := cliprovider.NewAdapter(cliprovider.Config{}, cipher)
	service := NewService(repository, nil, nil, nil, provider.NewRegistry(adapter), cipher, nil)

	result, err := service.ExportCLIProxyCredentials(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if result.Count != 1 {
		t.Fatalf("count = %d", result.Count)
	}
	if result.Filename != "xai-user@example.com.json" {
		t.Fatalf("filename = %q", result.Filename)
	}
	if result.ContentType != "application/json; charset=utf-8" {
		t.Fatalf("content-type = %q", result.ContentType)
	}
	var doc map[string]any
	if err := json.Unmarshal(result.Data, &doc); err != nil {
		t.Fatal(err)
	}
	if doc["type"] != "xai" || doc["auth_kind"] != "oauth" || doc["access_token"] != "access-token" || doc["refresh_token"] != "refresh-token" {
		t.Fatalf("doc markers/tokens = %#v", doc)
	}
	if doc["email"] != "user@example.com" || doc["sub"] != "user-1" {
		t.Fatalf("identity = %#v", doc)
	}
	if doc["base_url"] != "https://api.x.ai/v1" || doc["token_endpoint"] != "https://auth.x.ai/oauth2/token" || doc["token_type"] != "Bearer" {
		t.Fatalf("endpoints = %#v", doc)
	}
	if doc["disabled"] != false {
		t.Fatalf("disabled = %#v", doc["disabled"])
	}
}

func TestExportCLIProxyCredentialsMultiZip(t *testing.T) {
	ctx := context.Background()
	database, err := relational.OpenSQLite(ctx, filepath.Join(t.TempDir(), "cliproxy-export-multi.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	if err := database.InitializeSchema(ctx); err != nil {
		t.Fatal(err)
	}
	cipher, err := security.NewCipher(base64.StdEncoding.EncodeToString(make([]byte, 32)))
	if err != nil {
		t.Fatal(err)
	}
	encrypt := func(value string) string {
		out, err := cipher.Encrypt(value)
		if err != nil {
			t.Fatal(err)
		}
		return out
	}
	repository := relational.NewAccountRepository(database)
	for i, email := range []string{"one@example.com", "two@example.com"} {
		if _, _, err := repository.UpsertByIdentity(ctx, accountdomain.Credential{
			Provider: accountdomain.ProviderBuild, Name: email, Email: email, UserID: "user-" + string(rune('1'+i)),
			SourceKey: "cliproxy-multi-" + email, EncryptedAccessToken: encrypt("access-" + email),
			EncryptedRefreshToken: encrypt("refresh-" + email), Enabled: true,
			AuthStatus: accountdomain.AuthStatusActive, Priority: i, MaxConcurrent: 8,
		}); err != nil {
			t.Fatal(err)
		}
	}
	adapter := cliprovider.NewAdapter(cliprovider.Config{}, cipher)
	service := NewService(repository, nil, nil, nil, provider.NewRegistry(adapter), cipher, nil)

	result, err := service.ExportCLIProxyCredentials(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if result.Count != 2 {
		t.Fatalf("count = %d", result.Count)
	}
	if result.ContentType != "application/zip" {
		t.Fatalf("content-type = %q", result.ContentType)
	}
	if !strings.HasSuffix(result.Filename, ".zip") || !strings.HasPrefix(result.Filename, "cliproxyapi-xai-auth-") {
		t.Fatalf("filename = %q", result.Filename)
	}
	if len(result.Data) < 4 || result.Data[0] != 'P' || result.Data[1] != 'K' {
		t.Fatalf("zip magic missing")
	}
}

func TestExportCLIProxyCredentialsEmpty(t *testing.T) {
	ctx := context.Background()
	database, err := relational.OpenSQLite(ctx, filepath.Join(t.TempDir(), "cliproxy-export-empty.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()
	if err := database.InitializeSchema(ctx); err != nil {
		t.Fatal(err)
	}
	cipher, err := security.NewCipher(base64.StdEncoding.EncodeToString(make([]byte, 32)))
	if err != nil {
		t.Fatal(err)
	}
	adapter := cliprovider.NewAdapter(cliprovider.Config{}, cipher)
	service := NewService(relational.NewAccountRepository(database), nil, nil, nil, provider.NewRegistry(adapter), cipher, nil)
	if _, err := service.ExportCLIProxyCredentials(ctx); !errors.Is(err, ErrExportEmpty) {
		t.Fatalf("error = %v, want ErrExportEmpty", err)
	}
}
