package googledeploymentmanagerdeployment


type GoogleDeploymentManagerDeploymentTargetImports struct {
	// The full contents of the template that you want to import.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_deployment_manager_deployment#content GoogleDeploymentManagerDeployment#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The name of the template to import, as declared in the YAML configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_deployment_manager_deployment#name GoogleDeploymentManagerDeployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

