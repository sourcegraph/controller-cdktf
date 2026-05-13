package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfigDnsPeeringConfigs struct {
	// Required. The DNS name suffix of the zone being peered to, e.g., "my-internal-domain.corp.". Must end with a dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#domain GoogleVertexAiReasoningEngine#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Required. The VPC network name in the targetProject where the DNS zone specified by 'domain' is visible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#target_network GoogleVertexAiReasoningEngine#target_network}
	TargetNetwork *string `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
	// Required.
	//
	// The project id hosting the Cloud DNS managed
	// zone that contains the 'domain'.
	// The Vertex AI service Agent requires the dns.peer role
	// on this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#target_project GoogleVertexAiReasoningEngine#target_project}
	TargetProject *string `field:"required" json:"targetProject" yaml:"targetProject"`
}

