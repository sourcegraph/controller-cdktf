package clouddeploycustomtargettype


type ClouddeployCustomTargetTypeTasksDeployContainer struct {
	// Image is the container image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_custom_target_type#image ClouddeployCustomTargetType#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Args is the container arguments to use. This overrides the default arguments defined in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_custom_target_type#args ClouddeployCustomTargetType#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Command is the container entrypoint to use. This overrides the default entrypoint defined in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_custom_target_type#command ClouddeployCustomTargetType#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Environment variables that are set in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/clouddeploy_custom_target_type#env ClouddeployCustomTargetType#env}
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

