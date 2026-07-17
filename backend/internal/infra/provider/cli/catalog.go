package cli

import (
	"github.com/chenyme/grok2api/backend/internal/domain/account"
	modeldomain "github.com/chenyme/grok2api/backend/internal/domain/model"
)

type ModelSpec struct {
	PublicID      string
	UpstreamModel string
}

var catalog = []ModelSpec{
	{PublicID: "grok-build-0.1", UpstreamModel: "grok-build-0.1"},
	{PublicID: "grok-4.5", UpstreamModel: "grok-4.5"},
	{PublicID: "grok-4.3", UpstreamModel: "grok-4.3"},
	{PublicID: "grok-4.20-0309-reasoning", UpstreamModel: "grok-4.20-0309-reasoning"},
	{PublicID: "grok-4.20-0309-non-reasoning", UpstreamModel: "grok-4.20-0309-non-reasoning"},
	{PublicID: "grok-4.20-multi-agent-0309", UpstreamModel: "grok-4.20-multi-agent-0309"},
	{PublicID: "grok-3-mini", UpstreamModel: "grok-3-mini"},
	{PublicID: "grok-3-mini-fast", UpstreamModel: "grok-3-mini-fast"},
}

func Catalog() []ModelSpec { return append([]ModelSpec(nil), catalog...) }

func Routes() []modeldomain.Route {
	values := make([]modeldomain.Route, 0, len(catalog))
	for _, spec := range catalog {
		publicID, _ := modeldomain.NormalizePublicID(account.ProviderBuild, spec.PublicID)
		values = append(values, modeldomain.Route{
			PublicID: publicID, Provider: account.ProviderBuild, UpstreamModel: spec.UpstreamModel,
			Capability: modeldomain.CapabilityResponses, Origin: modeldomain.OriginCatalog, Enabled: true,
		})
	}
	return values
}

func Resolve(upstreamModel string) (ModelSpec, bool) {
	for _, spec := range catalog {
		if spec.UpstreamModel == upstreamModel {
			return spec, true
		}
	}
	return ModelSpec{}, false
}
