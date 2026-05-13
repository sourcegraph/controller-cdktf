package observabilitytracescope


type ObservabilityTraceScopeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#create ObservabilityTraceScope#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#delete ObservabilityTraceScope#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/observability_trace_scope#update ObservabilityTraceScope#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

