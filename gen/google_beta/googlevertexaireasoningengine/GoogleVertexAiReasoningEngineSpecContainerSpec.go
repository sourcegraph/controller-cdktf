package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecContainerSpec struct {
	// The Artifact Registry Docker image URI (e.g., 'us-central1-docker.pkg.dev/my-project/my-repo/my-image:tag') of the container image that is to be run on each worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#image_uri GoogleVertexAiReasoningEngine#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
}

