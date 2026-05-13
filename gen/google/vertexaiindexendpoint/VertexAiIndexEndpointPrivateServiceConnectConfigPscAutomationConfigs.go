package vertexaiindexendpoint


type VertexAiIndexEndpointPrivateServiceConnectConfigPscAutomationConfigs struct {
	// The full name of the Google Compute Engine [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks). [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/get): projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_index_endpoint#network VertexAiIndexEndpoint#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Project id used to create forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_index_endpoint#project_id VertexAiIndexEndpoint#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

