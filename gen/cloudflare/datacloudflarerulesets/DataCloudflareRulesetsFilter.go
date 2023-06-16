package datacloudflarerulesets


type DataCloudflareRulesetsFilter struct {
	// The ID of the Ruleset to target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/rulesets#id DataCloudflareRulesets#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Type of Ruleset to create. Available values: `custom`, `managed`, `root`, `schema`, `zone`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/rulesets#kind DataCloudflareRulesets#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Name of the ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/rulesets#name DataCloudflareRulesets#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Point in the request/response lifecycle where the ruleset will be created.
	//
	// Available values: `ddos_l4`, `ddos_l7`, `http_custom_errors`, `http_log_custom_fields`, `http_request_cache_settings`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_late_transform_managed`, `http_request_main`, `http_request_origin`, `http_request_dynamic_redirect`, `http_request_redirect`, `http_request_sanitize`, `http_request_transform`, `http_response_firewall_managed`, `http_response_headers_transform`, `http_response_headers_transform_managed`, `magic_transit`, `http_ratelimit`, `http_request_sbfm`, `http_config_settings`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/rulesets#phase DataCloudflareRulesets#phase}
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	// Version of the ruleset to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/rulesets#version DataCloudflareRulesets#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

