package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits struct {
	// The maximum number of nodes for this specific replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_instance#max_nodes GoogleSpannerInstance#max_nodes}
	MaxNodes *float64 `field:"required" json:"maxNodes" yaml:"maxNodes"`
	// The minimum number of nodes for this specific replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_instance#min_nodes GoogleSpannerInstance#min_nodes}
	MinNodes *float64 `field:"required" json:"minNodes" yaml:"minNodes"`
}

