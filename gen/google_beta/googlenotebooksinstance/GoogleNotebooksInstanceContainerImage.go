package googlenotebooksinstance


type GoogleNotebooksInstanceContainerImage struct {
	// The path to the container image repository. For example: gcr.io/{project_id}/{imageName}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#repository GoogleNotebooksInstance#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The tag of the container image. If not specified, this defaults to the latest tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#tag GoogleNotebooksInstance#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

