package googleclouddeploycustomtargettype


type GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage struct {
	// Cloud Storage source paths to copy recursively.
	//
	// For example, providing 'gs://my-bucket/dir/configs/*' will result in Skaffold copying all files within the 'dir/configs' directory in the bucket 'my-bucket'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#source GoogleClouddeployCustomTargetType#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Relative path from the source to the Skaffold file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type#path GoogleClouddeployCustomTargetType#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

