package googlevertexaifeatureonlinestore


type GoogleVertexAiFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig struct {
	// If set to true, customers will use private service connection to send request.
	//
	// Otherwise, the connection will set to public endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#enable_private_service_connect GoogleVertexAiFeatureOnlineStore#enable_private_service_connect}
	EnablePrivateServiceConnect interface{} `field:"required" json:"enablePrivateServiceConnect" yaml:"enablePrivateServiceConnect"`
	// A list of Projects from which the forwarding rule will target the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store#project_allowlist GoogleVertexAiFeatureOnlineStore#project_allowlist}
	ProjectAllowlist *[]*string `field:"optional" json:"projectAllowlist" yaml:"projectAllowlist"`
}

