package googleclouddeploycustomtargettype


type GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit struct {
	// Git repository the package should be cloned from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#repo GoogleClouddeployCustomTargetType#repo}
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// Relative path from the repository root to the Skaffold file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#path GoogleClouddeployCustomTargetType#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Git ref the package should be cloned from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#ref GoogleClouddeployCustomTargetType#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

