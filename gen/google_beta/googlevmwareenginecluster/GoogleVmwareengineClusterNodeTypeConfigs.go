package googlevmwareenginecluster


type GoogleVmwareengineClusterNodeTypeConfigs struct {
	// The number of nodes of this type in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#node_count GoogleVmwareengineCluster#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#node_type_id GoogleVmwareengineCluster#node_type_id}.
	NodeTypeId *string `field:"required" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Customized number of cores available to each node of the type.
	//
	// This number must always be one of 'nodeType.availableCustomCoreCounts'.
	// If zero is provided max value from 'nodeType.availableCustomCoreCounts' will be used.
	// Once the customer is created then corecount cannot be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#custom_core_count GoogleVmwareengineCluster#custom_core_count}
	CustomCoreCount *float64 `field:"optional" json:"customCoreCount" yaml:"customCoreCount"`
}

