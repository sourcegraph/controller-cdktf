package googleclouddeploycustomtargettype


type GoogleClouddeployCustomTargetTypeTasksDeployContainer struct {
	// Image is the container image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_custom_target_type#image GoogleClouddeployCustomTargetType#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Args is the container arguments to use. This overrides the default arguments defined in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_custom_target_type#args GoogleClouddeployCustomTargetType#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Command is the container entrypoint to use. This overrides the default entrypoint defined in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_custom_target_type#command GoogleClouddeployCustomTargetType#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Environment variables that are set in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_custom_target_type#env GoogleClouddeployCustomTargetType#env}
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

