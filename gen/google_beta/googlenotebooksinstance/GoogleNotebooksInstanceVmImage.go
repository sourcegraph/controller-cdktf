package googlenotebooksinstance


type GoogleNotebooksInstanceVmImage struct {
	// The name of the Google Cloud project that this VM image belongs to. Format: projects/{project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#project GoogleNotebooksInstance#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Use this VM image family to find the image; the newest image in this family will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#image_family GoogleNotebooksInstance#image_family}
	ImageFamily *string `field:"optional" json:"imageFamily" yaml:"imageFamily"`
	// Use VM image name to find the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#image_name GoogleNotebooksInstance#image_name}
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
}

