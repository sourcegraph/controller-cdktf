package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring struct {
	// Specifies the list of backends Policy Controller will export to. Specifying an empty value `[]` disables metrics export.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#backends GoogleGkeHubFeatureMembership#backends}
	Backends *[]*string `field:"optional" json:"backends" yaml:"backends"`
}

