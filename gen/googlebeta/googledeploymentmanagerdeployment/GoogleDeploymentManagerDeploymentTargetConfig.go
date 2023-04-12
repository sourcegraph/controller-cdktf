package googledeploymentmanagerdeployment


type GoogleDeploymentManagerDeploymentTargetConfig struct {
	// The full YAML contents of your configuration file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_deployment_manager_deployment#content GoogleDeploymentManagerDeployment#content}
	Content *string `field:"required" json:"content" yaml:"content"`
}

