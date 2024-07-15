package googlecomputeregionautoscaler


type GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControl struct {
	// max_scaled_down_replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_autoscaler#max_scaled_down_replicas GoogleComputeRegionAutoscaler#max_scaled_down_replicas}
	MaxScaledDownReplicas *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas `field:"optional" json:"maxScaledDownReplicas" yaml:"maxScaledDownReplicas"`
	// How long back autoscaling should look when computing recommendations to include directives regarding slower scale down, as described above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_autoscaler#time_window_sec GoogleComputeRegionAutoscaler#time_window_sec}
	TimeWindowSec *float64 `field:"optional" json:"timeWindowSec" yaml:"timeWindowSec"`
}

