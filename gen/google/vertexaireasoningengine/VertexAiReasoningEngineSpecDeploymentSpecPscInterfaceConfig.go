package vertexaireasoningengine


type VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig struct {
	// dns_peering_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#dns_peering_configs VertexAiReasoningEngine#dns_peering_configs}
	DnsPeeringConfigs interface{} `field:"optional" json:"dnsPeeringConfigs" yaml:"dnsPeeringConfigs"`
	// Optional.
	//
	// The name of the Compute Engine network attachment
	// to attach to the resource within the region and user project.
	// To specify this field, you must have already created a network attachment.
	// This field is only used for resources using PSC-Interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#network_attachment VertexAiReasoningEngine#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
}

