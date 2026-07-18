package cli

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode"
)

const (
	cliproxyAuthType     = "xai"
	cliproxyAuthKind     = "oauth"
	cliproxyTokenType    = "Bearer"
	cliproxyBaseURL      = "https://api.x.ai/v1"
	cliproxyFilenamePref = "xai-"
)

// CLIProxyAuthFile is one auth-dir entry ready for single JSON or multi ZIP packaging.
type CLIProxyAuthFile struct {
	Name string
	Body []byte
}

// cliproxyAuthDocument matches CLIProxyAPI auth-dir xai-*.json wire format.
// ExpiresIn is *int64 so 0 is emitted when expiry is set but already past.
type cliproxyAuthDocument struct {
	AccessToken   string `json:"access_token"`
	AuthKind      string `json:"auth_kind"`
	BaseURL       string `json:"base_url"`
	Disabled      bool   `json:"disabled"`
	Email         string `json:"email,omitempty"`
	Expired       string `json:"expired,omitempty"`
	ExpiresIn     *int64 `json:"expires_in,omitempty"`
	IDToken       string `json:"id_token,omitempty"`
	LastRefresh   string `json:"last_refresh"`
	Priority      int    `json:"priority"`
	RedirectURI   string `json:"redirect_uri,omitempty"`
	RefreshToken  string `json:"refresh_token,omitempty"`
	Sub           string `json:"sub,omitempty"`
	TokenEndpoint string `json:"token_endpoint"`
	TokenType     string `json:"token_type"`
	Type          string `json:"type"`
}

func buildCLIProxyAuthDocument(accessToken, refreshToken, email, userID string, enabled bool, priority int, expiresAt time.Time, lastRefreshAt *time.Time, now time.Time) cliproxyAuthDocument {
	nowUTC := now.UTC()
	doc := cliproxyAuthDocument{
		AccessToken:   accessToken,
		AuthKind:      cliproxyAuthKind,
		BaseURL:       cliproxyBaseURL,
		Disabled:      !enabled,
		Email:         strings.TrimSpace(email),
		LastRefresh:   nowUTC.Format(time.RFC3339),
		Priority:      priority,
		RefreshToken:  refreshToken,
		Sub:           strings.TrimSpace(userID),
		TokenEndpoint: defaultTokenURL,
		TokenType:     cliproxyTokenType,
		Type:          cliproxyAuthType,
	}
	if lastRefreshAt != nil && !lastRefreshAt.IsZero() {
		doc.LastRefresh = lastRefreshAt.UTC().Format(time.RFC3339)
	}
	if !expiresAt.IsZero() {
		expUTC := expiresAt.UTC()
		doc.Expired = expUTC.Format(time.RFC3339)
		seconds := int64(expUTC.Sub(nowUTC).Seconds())
		if seconds < 0 {
			seconds = 0
		}
		doc.ExpiresIn = &seconds
	}
	return doc
}

// MarshalCLIProxyAuthFile builds one CLIProxy auth JSON body + default filename for an account.
func MarshalCLIProxyAuthFile(
	accessToken, refreshToken, email, userID string,
	enabled bool, priority int,
	expiresAt time.Time, lastRefreshAt *time.Time, now time.Time,
	accountID uint64,
) (body []byte, filename string, err error) {
	doc := buildCLIProxyAuthDocument(accessToken, refreshToken, email, userID, enabled, priority, expiresAt, lastRefreshAt, now)
	body, err = marshalCLIProxyAuthDocument(doc)
	if err != nil {
		return nil, "", err
	}
	return body, cliproxyAuthFilename(email, userID, accountID), nil
}

func cliproxyAuthFilename(email, userID string, accountID uint64) string {
	identity := strings.TrimSpace(email)
	if identity == "" {
		identity = strings.TrimSpace(userID)
	}
	if identity == "" {
		return fmt.Sprintf("%saccount-%d.json", cliproxyFilenamePref, accountID)
	}
	return cliproxyFilenamePref + sanitizeCLIProxyFilenameSegment(identity) + ".json"
}

func sanitizeCLIProxyFilenameSegment(value string) string {
	var b strings.Builder
	b.Grow(len(value))
	for _, r := range value {
		switch {
		case r == '/' || r == '\\' || r == '"' || r == 0:
			continue
		case unicode.IsControl(r):
			continue
		default:
			b.WriteRune(r)
		}
	}
	if b.Len() == 0 {
		return "account"
	}
	return b.String()
}

// DisambiguateCLIProxyFilenames renames duplicate basenames using account IDs.
func DisambiguateCLIProxyFilenames(names []string, accountIDs []uint64) []string {
	if len(names) == 0 {
		return nil
	}
	out := make([]string, len(names))
	seen := make(map[string]int, len(names))
	for i, name := range names {
		count := seen[name]
		seen[name] = count + 1
		if count == 0 {
			out[i] = name
			continue
		}
		id := uint64(0)
		if i < len(accountIDs) {
			id = accountIDs[i]
		}
		base := strings.TrimSuffix(name, ".json")
		out[i] = fmt.Sprintf("%s__%d.json", base, id)
	}
	// Second pass if disambiguated names still collide (rare).
	seen = make(map[string]int, len(out))
	for i, name := range out {
		count := seen[name]
		seen[name] = count + 1
		if count == 0 {
			continue
		}
		id := uint64(0)
		if i < len(accountIDs) {
			id = accountIDs[i]
		}
		base := strings.TrimSuffix(name, ".json")
		out[i] = fmt.Sprintf("%s__%d.json", base, id)
	}
	return out
}

func marshalCLIProxyAuthDocument(doc cliproxyAuthDocument) ([]byte, error) {
	data, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}
	data = append(data, '\n')
	return data, nil
}

// BuildCLIProxyAuthExport packages one JSON file or a multi-account ZIP.
func BuildCLIProxyAuthExport(entries []CLIProxyAuthFile) (data []byte, filename string, contentType string, err error) {
	if len(entries) == 0 {
		return nil, "", "", fmt.Errorf("no CLIProxy auth entries")
	}
	if len(entries) == 1 {
		return entries[0].Body, entries[0].Name, "application/json; charset=utf-8", nil
	}
	var buf bytes.Buffer
	writer := zip.NewWriter(&buf)
	for _, entry := range entries {
		file, createErr := writer.Create(entry.Name)
		if createErr != nil {
			_ = writer.Close()
			return nil, "", "", createErr
		}
		if _, writeErr := file.Write(entry.Body); writeErr != nil {
			_ = writer.Close()
			return nil, "", "", writeErr
		}
	}
	if closeErr := writer.Close(); closeErr != nil {
		return nil, "", "", closeErr
	}
	filename = "cliproxyapi-xai-auth-" + time.Now().UTC().Format("20060102T150405Z") + ".zip"
	return buf.Bytes(), filename, "application/zip", nil
}
