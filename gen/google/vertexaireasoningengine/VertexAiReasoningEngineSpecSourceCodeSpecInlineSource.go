package vertexaireasoningengine


type VertexAiReasoningEngineSpecSourceCodeSpecInlineSource struct {
	// Required. Input only. The application source code archive, provided as a compressed tarball (.tar.gz) file. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#source_archive VertexAiReasoningEngine#source_archive}
	SourceArchive *string `field:"optional" json:"sourceArchive" yaml:"sourceArchive"`
}

