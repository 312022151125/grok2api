package web

import (
	"strings"
	"testing"

	"github.com/chenyme/grok2api/backend/internal/domain/account"
)

func TestParseImportedCredentialsAcceptsOneSSOTokenPerLine(t *testing.T) {
	adapter := &Adapter{}
	values, err := adapter.ParseImportedCredentials([]byte("token-one\nsso=token-two; other=drop\n\ntoken-one\n"))
	if err != nil {
		t.Fatal(err)
	}
	if len(values) != 2 {
		t.Fatalf("credentials = %#v", values)
	}
	if values[0].AccessToken != "token-one" || values[1].AccessToken != "token-two" {
		t.Fatalf("tokens = %q, %q", values[0].AccessToken, values[1].AccessToken)
	}
	for _, value := range values {
		if value.Provider != account.ProviderWeb || value.AuthType != account.AuthTypeSSO || value.WebTier != account.WebTierAuto {
			t.Fatalf("credential = %#v", value)
		}
	}
}

func TestParseImportedCredentialsRejectsOversizedPlainToken(t *testing.T) {
	adapter := &Adapter{}
	_, err := adapter.ParseImportedCredentials([]byte(strings.Repeat("x", maxSSOTokenBytes+1)))
	if err == nil {
		t.Fatal("expected oversized token error")
	}
}

func TestWebCredentialJSONUsesCurrentDocumentShape(t *testing.T) {
	adapter := &Adapter{}
	values, err := adapter.ParseImportedCredentials([]byte(`{"provider":"grok_web","accounts":[{"name":"primary","sso_token":"token-one","tier":"super","cloudflare_cookies":"cf_clearance=abc; sso=drop"}]}`))
	if err != nil || len(values) != 1 || values[0].Name != "primary" || values[0].AccessToken != "token-one" || values[0].WebTier != account.WebTierSuper || values[0].CloudflareCookies != "cf_clearance=abc; sso=drop" {
		t.Fatalf("credentials = %#v, err = %v", values, err)
	}
	data, err := adapter.MarshalCredentials(values)
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	if !strings.Contains(text, `"provider": "grok_web"`) || !strings.Contains(text, `"cloudflare_cookies": "cf_clearance=abc; sso=drop"`) {
		t.Fatalf("export lost current document fields: %s", data)
	}
	if strings.Contains(text, `"token":`) || strings.Contains(text, `"version"`) {
		t.Fatalf("export contains legacy metadata: %s", data)
	}
	roundTrip, err := adapter.ParseImportedCredentials(data)
	if err != nil || len(roundTrip) != 1 || roundTrip[0].Name != "primary" || roundTrip[0].AccessToken != "token-one" || roundTrip[0].WebTier != account.WebTierSuper || roundTrip[0].CloudflareCookies != "cf_clearance=abc; sso=drop" {
		t.Fatalf("round-trip credentials = %#v, err = %v", roundTrip, err)
	}
	if _, err := adapter.ParseImportedCredentials([]byte(`{"basic":["token-one"]}`)); err == nil {
		t.Fatal("legacy tier pools were accepted")
	}
}
