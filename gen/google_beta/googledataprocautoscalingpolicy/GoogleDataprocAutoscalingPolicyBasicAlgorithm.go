package googledataprocautoscalingpolicy


type GoogleDataprocAutoscalingPolicyBasicAlgorithm struct {
	// yarn_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy#yarn_config GoogleDataprocAutoscalingPolicy#yarn_config}
	YarnConfig *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig `field:"required" json:"yarnConfig" yaml:"yarnConfig"`
	// Duration between scaling events. A scaling period starts after the update operation from the previous event has completed.
	//
	// Bounds: [2m, 1d]. Default: 2m.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy#cooldown_period GoogleDataprocAutoscalingPolicy#cooldown_period}
	CooldownPeriod *string `field:"optional" json:"cooldownPeriod" yaml:"cooldownPeriod"`
}

