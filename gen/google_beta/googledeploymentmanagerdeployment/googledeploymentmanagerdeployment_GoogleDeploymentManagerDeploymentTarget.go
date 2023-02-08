package googledeploymentmanagerdeployment


type GoogleDeploymentManagerDeploymentTarget struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_deployment_manager_deployment#config GoogleDeploymentManagerDeployment#config}
	Config *GoogleDeploymentManagerDeploymentTargetConfig `field:"required" json:"config" yaml:"config"`
	// imports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_deployment_manager_deployment#imports GoogleDeploymentManagerDeployment#imports}
	Imports interface{} `field:"optional" json:"imports" yaml:"imports"`
}

