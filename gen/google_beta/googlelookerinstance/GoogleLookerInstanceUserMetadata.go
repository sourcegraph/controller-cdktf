package googlelookerinstance


type GoogleLookerInstanceUserMetadata struct {
	// Number of additional Developer Users to allocate to the Looker Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#additional_developer_user_count GoogleLookerInstance#additional_developer_user_count}
	AdditionalDeveloperUserCount *float64 `field:"optional" json:"additionalDeveloperUserCount" yaml:"additionalDeveloperUserCount"`
	// Number of additional Standard Users to allocate to the Looker Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#additional_standard_user_count GoogleLookerInstance#additional_standard_user_count}
	AdditionalStandardUserCount *float64 `field:"optional" json:"additionalStandardUserCount" yaml:"additionalStandardUserCount"`
	// Number of additional Viewer Users to allocate to the Looker Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#additional_viewer_user_count GoogleLookerInstance#additional_viewer_user_count}
	AdditionalViewerUserCount *float64 `field:"optional" json:"additionalViewerUserCount" yaml:"additionalViewerUserCount"`
}

