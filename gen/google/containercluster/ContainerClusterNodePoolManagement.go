package containercluster

type ContainerClusterNodePoolManagement struct {
	// Whether the nodes will be automatically repaired. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_cluster#auto_repair ContainerCluster#auto_repair}
	AutoRepair any `field:"optional" json:"autoRepair" yaml:"autoRepair"`
	// Whether the nodes will be automatically upgraded. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_cluster#auto_upgrade ContainerCluster#auto_upgrade}
	AutoUpgrade any `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
}
