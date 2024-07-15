package googleedgecontainercluster


type GoogleEdgecontainerClusterMaintenancePolicy struct {
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#window GoogleEdgecontainerCluster#window}
	Window *GoogleEdgecontainerClusterMaintenancePolicyWindow `field:"required" json:"window" yaml:"window"`
	// maintenance_exclusions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#maintenance_exclusions GoogleEdgecontainerCluster#maintenance_exclusions}
	MaintenanceExclusions interface{} `field:"optional" json:"maintenanceExclusions" yaml:"maintenanceExclusions"`
}

