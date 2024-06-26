package edgecontainercluster


type EdgecontainerClusterMaintenancePolicy struct {
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/edgecontainer_cluster#window EdgecontainerCluster#window}
	Window *EdgecontainerClusterMaintenancePolicyWindow `field:"required" json:"window" yaml:"window"`
}

