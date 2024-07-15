package googleworkbenchinstance


type GoogleWorkbenchInstanceGceSetupContainerImage struct {
	// The path to the container image repository. For example: gcr.io/{project_id}/{imageName}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#repository GoogleWorkbenchInstance#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The tag of the container image. If not specified, this defaults to the latest tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#tag GoogleWorkbenchInstance#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

