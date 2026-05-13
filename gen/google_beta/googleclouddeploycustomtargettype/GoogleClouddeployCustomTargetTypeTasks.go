package googleclouddeploycustomtargettype


type GoogleClouddeployCustomTargetTypeTasks struct {
	// deploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_custom_target_type#deploy GoogleClouddeployCustomTargetType#deploy}
	Deploy *GoogleClouddeployCustomTargetTypeTasksDeploy `field:"required" json:"deploy" yaml:"deploy"`
	// render block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_custom_target_type#render GoogleClouddeployCustomTargetType#render}
	Render *GoogleClouddeployCustomTargetTypeTasksRender `field:"optional" json:"render" yaml:"render"`
}

