package googledeploymentmanagerdeployment


type GoogleDeploymentManagerDeploymentTargetConfig struct {
	// The full YAML contents of your configuration file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_deployment_manager_deployment#content GoogleDeploymentManagerDeployment#content}
	Content *string `field:"required" json:"content" yaml:"content"`
}

