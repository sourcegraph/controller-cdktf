package googlecontainercluster


type GoogleContainerClusterNodePoolManagement struct {
	// Whether the nodes will be automatically repaired.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#auto_repair GoogleContainerCluster#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
	// Whether the nodes will be automatically upgraded.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#auto_upgrade GoogleContainerCluster#auto_upgrade}
	AutoUpgrade interface{} `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
}

