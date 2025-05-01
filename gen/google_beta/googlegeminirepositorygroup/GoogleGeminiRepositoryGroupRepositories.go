package googlegeminirepositorygroup


type GoogleGeminiRepositoryGroupRepositories struct {
	// Required. The Git branch pattern used for indexing in RE2 syntax. See https://github.com/google/re2/wiki/syntax for syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#branch_pattern GoogleGeminiRepositoryGroup#branch_pattern}
	BranchPattern *string `field:"required" json:"branchPattern" yaml:"branchPattern"`
	// Required. The DeveloperConnect repository full resource name, relative resource name or resource URL to be indexed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group#resource GoogleGeminiRepositoryGroup#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
}

