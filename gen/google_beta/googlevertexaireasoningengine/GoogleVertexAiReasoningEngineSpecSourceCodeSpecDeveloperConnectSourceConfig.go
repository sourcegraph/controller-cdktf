package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecSourceCodeSpecDeveloperConnectSourceConfig struct {
	// Directory, relative to the source root, in which to run the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#dir GoogleVertexAiReasoningEngine#dir}
	Dir *string `field:"required" json:"dir" yaml:"dir"`
	// The Developer Connect Git repository link, formatted as projects/* /locations/* /connections/* /gitRepositoryLink/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#git_repository_link GoogleVertexAiReasoningEngine#git_repository_link}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	GitRepositoryLink *string `field:"required" json:"gitRepositoryLink" yaml:"gitRepositoryLink"`
	// The revision to fetch from the Git repository such as a branch, a tag, a commit SHA, or any Git ref.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#revision GoogleVertexAiReasoningEngine#revision}
	Revision *string `field:"required" json:"revision" yaml:"revision"`
}

