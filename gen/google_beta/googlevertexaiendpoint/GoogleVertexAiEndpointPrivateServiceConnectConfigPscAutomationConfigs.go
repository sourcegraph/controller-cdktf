package googlevertexaiendpoint


type GoogleVertexAiEndpointPrivateServiceConnectConfigPscAutomationConfigs struct {
	// The full name of the Google Compute Engine [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks). [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/get): projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint#network GoogleVertexAiEndpoint#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Project id used to create forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint#project_id GoogleVertexAiEndpoint#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

