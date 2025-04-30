package googlevmwareenginecluster


type GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#autoscale_policy_id GoogleVmwareengineCluster#autoscale_policy_id}.
	AutoscalePolicyId *string `field:"required" json:"autoscalePolicyId" yaml:"autoscalePolicyId"`
	// The canonical identifier of the node type to add or remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#node_type_id GoogleVmwareengineCluster#node_type_id}
	NodeTypeId *string `field:"required" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Number of nodes to add to a cluster during a scale-out operation. Must be divisible by 2 for stretched clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#scale_out_size GoogleVmwareengineCluster#scale_out_size}
	ScaleOutSize *float64 `field:"required" json:"scaleOutSize" yaml:"scaleOutSize"`
	// consumed_memory_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#consumed_memory_thresholds GoogleVmwareengineCluster#consumed_memory_thresholds}
	ConsumedMemoryThresholds *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds `field:"optional" json:"consumedMemoryThresholds" yaml:"consumedMemoryThresholds"`
	// cpu_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#cpu_thresholds GoogleVmwareengineCluster#cpu_thresholds}
	CpuThresholds *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds `field:"optional" json:"cpuThresholds" yaml:"cpuThresholds"`
	// storage_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_cluster#storage_thresholds GoogleVmwareengineCluster#storage_thresholds}
	StorageThresholds *GoogleVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds `field:"optional" json:"storageThresholds" yaml:"storageThresholds"`
}

