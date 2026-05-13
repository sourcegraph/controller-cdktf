package googlecloudidentitypolicy


type GoogleCloudIdentityPolicySetting struct {
	// The type of the Setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#type GoogleCloudIdentityPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The value of the Setting as JSON string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#value_json GoogleCloudIdentityPolicy#value_json}
	ValueJson *string `field:"required" json:"valueJson" yaml:"valueJson"`
}

