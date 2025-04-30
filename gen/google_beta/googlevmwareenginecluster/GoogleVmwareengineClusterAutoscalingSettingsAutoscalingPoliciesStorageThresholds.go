package googlevmwareenginecluster


type GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds struct {
	// The utilization triggering the scale-in operation in percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#scale_in GoogleVmwareengineCluster#scale_in}
	ScaleIn *float64 `field:"required" json:"scaleIn" yaml:"scaleIn"`
	// The utilization triggering the scale-out operation in percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#scale_out GoogleVmwareengineCluster#scale_out}
	ScaleOut *float64 `field:"required" json:"scaleOut" yaml:"scaleOut"`
}

