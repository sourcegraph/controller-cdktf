package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig struct {
	// dedicated_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#dedicated_resources GoogleVertexAiEndpointWithModelGardenDeployment#dedicated_resources}
	DedicatedResources *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources `field:"optional" json:"dedicatedResources" yaml:"dedicatedResources"`
	// If true, enable the QMT fast tryout feature for this model if possible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#fast_tryout_enabled GoogleVertexAiEndpointWithModelGardenDeployment#fast_tryout_enabled}
	FastTryoutEnabled interface{} `field:"optional" json:"fastTryoutEnabled" yaml:"fastTryoutEnabled"`
	// System labels for Model Garden deployments. These labels are managed by Google and for tracking purposes only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#system_labels GoogleVertexAiEndpointWithModelGardenDeployment#system_labels}
	SystemLabels *map[string]*string `field:"optional" json:"systemLabels" yaml:"systemLabels"`
}

