package googlecolabnotebookexecution


type GoogleColabNotebookExecutionDataformRepositorySource struct {
	// The resource name of the Dataform Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_notebook_execution#dataform_repository_resource_name GoogleColabNotebookExecution#dataform_repository_resource_name}
	DataformRepositoryResourceName *string `field:"required" json:"dataformRepositoryResourceName" yaml:"dataformRepositoryResourceName"`
	// The commit SHA to read repository with. If unset, the file will be read at HEAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_notebook_execution#commit_sha GoogleColabNotebookExecution#commit_sha}
	CommitSha *string `field:"optional" json:"commitSha" yaml:"commitSha"`
}

