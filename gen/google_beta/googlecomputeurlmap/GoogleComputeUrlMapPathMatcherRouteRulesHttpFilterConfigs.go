package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherRouteRulesHttpFilterConfigs struct {
	// The configuration needed to enable the networkservices.HttpFilter resource. The configuration must be YAML formatted and only contain fields defined in the protobuf identified in configTypeUrl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#config GoogleComputeUrlMap#config}
	Config *string `field:"optional" json:"config" yaml:"config"`
	// The fully qualified versioned proto3 type url of the protobuf that the filter expects for its contextual settings, for example: type.googleapis.com/google.protobuf.Struct.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#config_type_url GoogleComputeUrlMap#config_type_url}
	ConfigTypeUrl *string `field:"optional" json:"configTypeUrl" yaml:"configTypeUrl"`
	// Name of the networkservices.HttpFilter resource this configuration belongs to. This name must be known to the xDS client. Example: envoy.wasm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#filter_name GoogleComputeUrlMap#filter_name}
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

