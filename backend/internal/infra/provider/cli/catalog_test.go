package cli

import (
	"testing"

	"github.com/chenyme/grok2api/backend/internal/domain/account"
	modeldomain "github.com/chenyme/grok2api/backend/internal/domain/model"
)

func TestBuildCatalogRoutes(t *testing.T) {
	routes := Routes()
	if len(routes) != 8 {
		t.Fatalf("routes size = %d", len(routes))
	}
	wantUpstream := map[string]struct{}{
		"grok-build-0.1": {}, "grok-4.5": {}, "grok-4.3": {},
		"grok-4.20-0309-reasoning": {}, "grok-4.20-0309-non-reasoning": {}, "grok-4.20-multi-agent-0309": {},
		"grok-3-mini": {}, "grok-3-mini-fast": {},
	}
	seen := make(map[string]struct{}, len(routes))
	for _, route := range routes {
		if route.Provider != account.ProviderBuild {
			t.Fatalf("provider = %s", route.Provider)
		}
		if route.Capability != modeldomain.CapabilityResponses {
			t.Fatalf("capability = %s for %s", route.Capability, route.PublicID)
		}
		if route.Origin != modeldomain.OriginCatalog {
			t.Fatalf("origin = %s for %s", route.Origin, route.PublicID)
		}
		if !route.Enabled {
			t.Fatalf("route disabled: %s", route.PublicID)
		}
		wantPublic, ok := modeldomain.NormalizePublicID(account.ProviderBuild, route.UpstreamModel)
		if !ok || route.PublicID != wantPublic {
			t.Fatalf("public id = %q, want Build/%s", route.PublicID, route.UpstreamModel)
		}
		if _, exists := wantUpstream[route.UpstreamModel]; !exists {
			t.Fatalf("unexpected upstream %s", route.UpstreamModel)
		}
		if _, exists := seen[route.UpstreamModel]; exists {
			t.Fatalf("duplicate upstream %s", route.UpstreamModel)
		}
		seen[route.UpstreamModel] = struct{}{}
	}
	if len(seen) != len(wantUpstream) {
		t.Fatalf("upstream set size = %d, want %d", len(seen), len(wantUpstream))
	}
}
