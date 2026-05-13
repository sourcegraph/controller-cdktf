package vertexaireasoningengine


type VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfigDnsPeeringConfigs struct {
	// Required. The DNS name suffix of the zone being peered to, e.g., "my-internal-domain.corp.". Must end with a dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#domain VertexAiReasoningEngine#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Required. The VPC network name in the targetProject where the DNS zone specified by 'domain' is visible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#target_network VertexAiReasoningEngine#target_network}
	TargetNetwork *string `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
	// Required.
	//
	// The project id hosting the Cloud DNS managed
	// zone that contains the 'domain'.
	// The Vertex AI service Agent requires the dns.peer role
	// on this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#target_project VertexAiReasoningEngine#target_project}
	TargetProject *string `field:"required" json:"targetProject" yaml:"targetProject"`
}

