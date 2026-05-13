package googlegkehubrolloutsequence


type GoogleGkeHubRolloutSequenceStages struct {
	// List of Fleet projects to select the clusters from. Expected format: projects/{project}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#fleet_projects GoogleGkeHubRolloutSequence#fleet_projects}
	FleetProjects *[]*string `field:"required" json:"fleetProjects" yaml:"fleetProjects"`
	// cluster_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#cluster_selector GoogleGkeHubRolloutSequence#cluster_selector}
	ClusterSelector *GoogleGkeHubRolloutSequenceStagesClusterSelector `field:"optional" json:"clusterSelector" yaml:"clusterSelector"`
	// Soak time after upgrading all the clusters in the stage, specified in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#soak_duration GoogleGkeHubRolloutSequence#soak_duration}
	SoakDuration *string `field:"optional" json:"soakDuration" yaml:"soakDuration"`
}

