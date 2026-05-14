package googlegkehubrolloutsequence


type GoogleGkeHubRolloutSequenceAutoUpgradeConfigRolloutCreationScope struct {
	// The list of enabled upgrade types. Current valid values are 'CONTROL_PLANE_MINOR', 'CONTROL_PLANE_PATCH', 'NODE_MINOR', and 'NODE_PATCH'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_gke_hub_rollout_sequence#upgrade_types GoogleGkeHubRolloutSequence#upgrade_types}
	UpgradeTypes *[]*string `field:"optional" json:"upgradeTypes" yaml:"upgradeTypes"`
}

