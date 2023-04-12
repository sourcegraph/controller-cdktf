package googleclouddeploytarget


type GoogleClouddeployTargetAnthosCluster struct {
	// Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_target#membership GoogleClouddeployTarget#membership}
	Membership *string `field:"optional" json:"membership" yaml:"membership"`
}

