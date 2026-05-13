package vertexaireasoningengine


type VertexAiReasoningEngineSpecContainerSpec struct {
	// The Artifact Registry Docker image URI (e.g., 'us-central1-docker.pkg.dev/my-project/my-repo/my-image:tag') of the container image that is to be run on each worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#image_uri VertexAiReasoningEngine#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
}

