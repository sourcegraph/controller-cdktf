package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigBootDisk struct {
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#disk_type GoogleContainerNodePool#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Configured IOPs provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#provisioned_iops GoogleContainerNodePool#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Configured throughput provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#provisioned_throughput GoogleContainerNodePool#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#size_gb GoogleContainerNodePool#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

