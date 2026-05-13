package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig struct {
	// dns_peering_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#dns_peering_configs GoogleVertexAiReasoningEngine#dns_peering_configs}
	DnsPeeringConfigs interface{} `field:"optional" json:"dnsPeeringConfigs" yaml:"dnsPeeringConfigs"`
	// Optional.
	//
	// The name of the Compute Engine network attachment
	// to attach to the resource within the region and user project.
	// To specify this field, you must have already created a network attachment.
	// This field is only used for resources using PSC-Interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#network_attachment GoogleVertexAiReasoningEngine#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
}

