package googleworkbenchinstance


type GoogleWorkbenchInstanceGceSetupVmImage struct {
	// Optional. Use this VM image family to find the image; the newest image in this family will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#family GoogleWorkbenchInstance#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// Optional. Use VM image name to find the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#name GoogleWorkbenchInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the Google Cloud project that this VM image belongs to. Format: {project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#project GoogleWorkbenchInstance#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

