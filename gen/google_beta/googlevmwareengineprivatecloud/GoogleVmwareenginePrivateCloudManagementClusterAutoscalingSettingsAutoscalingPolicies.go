package googlevmwareengineprivatecloud


type GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_private_cloud#autoscale_policy_id GoogleVmwareenginePrivateCloud#autoscale_policy_id}.
	AutoscalePolicyId *string `field:"required" json:"autoscalePolicyId" yaml:"autoscalePolicyId"`
	// The canonical identifier of the node type to add or remove.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_private_cloud#node_type_id GoogleVmwareenginePrivateCloud#node_type_id}
	NodeTypeId *string `field:"required" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Number of nodes to add to a cluster during a scale-out operation. Must be divisible by 2 for stretched clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_private_cloud#scale_out_size GoogleVmwareenginePrivateCloud#scale_out_size}
	ScaleOutSize *float64 `field:"required" json:"scaleOutSize" yaml:"scaleOutSize"`
	// consumed_memory_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_private_cloud#consumed_memory_thresholds GoogleVmwareenginePrivateCloud#consumed_memory_thresholds}
	ConsumedMemoryThresholds *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds `field:"optional" json:"consumedMemoryThresholds" yaml:"consumedMemoryThresholds"`
	// cpu_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_private_cloud#cpu_thresholds GoogleVmwareenginePrivateCloud#cpu_thresholds}
	CpuThresholds *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds `field:"optional" json:"cpuThresholds" yaml:"cpuThresholds"`
	// storage_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vmwareengine_private_cloud#storage_thresholds GoogleVmwareenginePrivateCloud#storage_thresholds}
	StorageThresholds *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds `field:"optional" json:"storageThresholds" yaml:"storageThresholds"`
}

