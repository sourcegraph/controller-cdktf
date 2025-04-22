package googleclouddeploytarget


type GoogleClouddeployTargetAssociatedEntitiesAnthosClusters struct {
	// Optional. Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_target#membership GoogleClouddeployTarget#membership}
	Membership *string `field:"optional" json:"membership" yaml:"membership"`
}

