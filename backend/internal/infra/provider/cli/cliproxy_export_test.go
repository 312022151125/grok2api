package cli

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestBuildCLIProxyAuthDocumentFields(t *testing.T) {
	now := time.Date(2026, 7, 18, 0, 30, 55, 0, time.UTC)
	expiresAt := time.Date(2026, 7, 18, 6, 30, 55, 0, time.UTC)
	lastRefresh := time.Date(2026, 7, 18, 0, 30, 55, 0, time.UTC)
	doc := buildCLIProxyAuthDocument(
		"access", "refresh", "user@example.com", "user-1",
		true, 0, expiresAt, &lastRefresh, now,
	)
	if doc.Type != "xai" || doc.AuthKind != "oauth" || doc.TokenType != "Bearer" {
		t.Fatalf("auth markers = %#v", doc)
	}
	if doc.BaseURL != "https://api.x.ai/v1" || doc.TokenEndpoint != defaultTokenURL {
		t.Fatalf("endpoints = %#v", doc)
	}
	if doc.Disabled || doc.Priority != 0 || doc.Email != "user@example.com" || doc.Sub != "user-1" {
		t.Fatalf("identity/priority = %#v", doc)
	}
	if doc.Expired != "2026-07-18T06:30:55Z" || doc.ExpiresIn == nil || *doc.ExpiresIn != 21600 {
		t.Fatalf("expiry = expired=%q expires_in=%v", doc.Expired, doc.ExpiresIn)
	}
	if doc.LastRefresh != "2026-07-18T00:30:55Z" {
		t.Fatalf("last_refresh = %q", doc.LastRefresh)
	}
	if doc.IDToken != "" || doc.RedirectURI != "" {
		t.Fatalf("optional empty fields set: %#v", doc)
	}

	// Already expired: expires_in must be present as 0, not omitted.
	past := time.Date(2026, 7, 17, 0, 0, 0, 0, time.UTC)
	expiredDoc := buildCLIProxyAuthDocument("a", "r", "e@x.com", "sub", true, 0, past, nil, now)
	body, err := marshalCLIProxyAuthDocument(expiredDoc)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(body, []byte(`"expires_in":0`)) {
		t.Fatalf("expired account must emit expires_in:0, body=%s", body)
	}

	disabled := buildCLIProxyAuthDocument("a", "r", "", "sub", false, 3, time.Time{}, nil, now)
	if !disabled.Disabled || disabled.Priority != 3 || disabled.Expired != "" || disabled.ExpiresIn != nil {
		t.Fatalf("disabled/no-expiry = %#v", disabled)
	}
	if disabled.LastRefresh != now.Format(time.RFC3339) {
		t.Fatalf("fallback last_refresh = %q", disabled.LastRefresh)
	}
}

func TestCLIProxyAuthFilename(t *testing.T) {
	if got := cliproxyAuthFilename("user@example.com", "uid", 9); got != "xai-user@example.com.json" {
		t.Fatalf("email filename = %q", got)
	}
	if got := cliproxyAuthFilename("", "user-1", 9); got != "xai-user-1.json" {
		t.Fatalf("userID filename = %q", got)
	}
	if got := cliproxyAuthFilename("", "", 42); got != "xai-account-42.json" {
		t.Fatalf("id filename = %q", got)
	}
	if got := cliproxyAuthFilename(`bad/"\path`, "", 1); strings.ContainsAny(got, `/"\`) {
		t.Fatalf("hostile chars kept: %q", got)
	}
	if got := cliproxyAuthFilename(`natiel_sout@outlook.com`, "", 1); got != "xai-natiel_sout@outlook.com.json" {
		t.Fatalf("sample email filename = %q", got)
	}
}

func TestDisambiguateCLIProxyFilenames(t *testing.T) {
	names := []string{"xai-a@b.com.json", "xai-a@b.com.json", "xai-c@d.com.json"}
	ids := []uint64{1, 2, 3}
	got := DisambiguateCLIProxyFilenames(names, ids)
	if got[0] != "xai-a@b.com.json" || got[1] != "xai-a@b.com__2.json" || got[2] != "xai-c@d.com.json" {
		t.Fatalf("disambiguated = %#v", got)
	}
}

func TestBuildCLIProxyAuthExportSingleAndZip(t *testing.T) {
	doc := buildCLIProxyAuthDocument("access", "refresh", "user@example.com", "user-1", true, 0,
		time.Date(2026, 7, 18, 6, 30, 55, 0, time.UTC), nil, time.Date(2026, 7, 18, 0, 30, 55, 0, time.UTC))
	body, err := marshalCLIProxyAuthDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(body, []byte("\n  ")) {
		t.Fatalf("expected compact JSON, got indented: %s", body)
	}
	var roundTrip cliproxyAuthDocument
	if err := json.Unmarshal(body, &roundTrip); err != nil {
		t.Fatal(err)
	}
	if roundTrip.AccessToken != "access" || roundTrip.RefreshToken != "refresh" {
		t.Fatalf("round-trip tokens = %#v", roundTrip)
	}

	single, name, ctype, err := BuildCLIProxyAuthExport([]CLIProxyAuthFile{{Name: "xai-user@example.com.json", Body: body}})
	if err != nil {
		t.Fatal(err)
	}
	if name != "xai-user@example.com.json" || ctype != "application/json; charset=utf-8" || !bytes.Equal(single, body) {
		t.Fatalf("single export name=%q ctype=%q", name, ctype)
	}

	body2, err := marshalCLIProxyAuthDocument(buildCLIProxyAuthDocument("a2", "r2", "two@example.com", "u2", true, 1, time.Time{}, nil, time.Now().UTC()))
	if err != nil {
		t.Fatal(err)
	}
	zipData, zipName, zipType, err := BuildCLIProxyAuthExport([]CLIProxyAuthFile{
		{Name: "xai-user@example.com.json", Body: body},
		{Name: "xai-two@example.com.json", Body: body2},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(zipName, "cliproxyapi-xai-auth-") || !strings.HasSuffix(zipName, ".zip") {
		t.Fatalf("zip name = %q", zipName)
	}
	if zipType != "application/zip" || len(zipData) < 4 || string(zipData[:2]) != "PK" {
		t.Fatalf("zip payload invalid: type=%q len=%d", zipType, len(zipData))
	}
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		t.Fatal(err)
	}
	if len(reader.File) != 2 {
		t.Fatalf("zip entries = %d", len(reader.File))
	}
	for _, file := range reader.File {
		rc, openErr := file.Open()
		if openErr != nil {
			t.Fatal(openErr)
		}
		var parsed cliproxyAuthDocument
		if decodeErr := json.NewDecoder(rc).Decode(&parsed); decodeErr != nil {
			_ = rc.Close()
			t.Fatal(decodeErr)
		}
		_ = rc.Close()
		if parsed.Type != "xai" || parsed.AccessToken == "" {
			t.Fatalf("zip body for %s = %#v", file.Name, parsed)
		}
	}

	if _, _, _, err := BuildCLIProxyAuthExport(nil); err == nil {
		t.Fatal("empty export should fail")
	}
}

func TestMarshalCLIProxyAuthFile(t *testing.T) {
	now := time.Date(2026, 7, 18, 0, 0, 0, 0, time.UTC)
	body, name, err := MarshalCLIProxyAuthFile("access", "refresh", "user@example.com", "user-1", true, 2, now.Add(time.Hour), nil, now, 7)
	if err != nil {
		t.Fatal(err)
	}
	if name != "xai-user@example.com.json" {
		t.Fatalf("name = %q", name)
	}
	var doc cliproxyAuthDocument
	if err := json.Unmarshal(body, &doc); err != nil {
		t.Fatal(err)
	}
	if doc.AccessToken != "access" || doc.Priority != 2 || doc.Type != "xai" {
		t.Fatalf("doc = %#v", doc)
	}
}
