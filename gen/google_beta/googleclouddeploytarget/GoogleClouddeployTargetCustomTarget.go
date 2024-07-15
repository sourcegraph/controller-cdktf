package googleclouddeploytarget


type GoogleClouddeployTargetCustomTarget struct {
	// Required. The name of the CustomTargetType. Format must be `projects/{project}/locations/{location}/customTargetTypes/{custom_target_type}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#custom_target_type GoogleClouddeployTarget#custom_target_type}
	CustomTargetType *string `field:"required" json:"customTargetType" yaml:"customTargetType"`
}

