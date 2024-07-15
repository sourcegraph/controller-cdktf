package googleclouddeploycustomtargettype


type GoogleClouddeployCustomTargetTypeCustomActions struct {
	// The Skaffold custom action responsible for deploy operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#deploy_action GoogleClouddeployCustomTargetType#deploy_action}
	DeployAction *string `field:"required" json:"deployAction" yaml:"deployAction"`
	// include_skaffold_modules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#include_skaffold_modules GoogleClouddeployCustomTargetType#include_skaffold_modules}
	IncludeSkaffoldModules interface{} `field:"optional" json:"includeSkaffoldModules" yaml:"includeSkaffoldModules"`
	// The Skaffold custom action responsible for render operations.
	//
	// If not provided then Cloud Deploy will perform the render operations via 'skaffold render'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#render_action GoogleClouddeployCustomTargetType#render_action}
	RenderAction *string `field:"optional" json:"renderAction" yaml:"renderAction"`
}

