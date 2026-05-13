package googlegkehubrolloutsequence


type GoogleGkeHubRolloutSequenceStagesClusterSelector struct {
	// The label selector must be a valid CEL (Common Expression Language) expression which evaluates resource.labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#label_selector GoogleGkeHubRolloutSequence#label_selector}
	LabelSelector *string `field:"required" json:"labelSelector" yaml:"labelSelector"`
}

