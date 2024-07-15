package alloydbcluster


type AlloydbClusterMaintenanceUpdatePolicy struct {
	// maintenance_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/alloydb_cluster#maintenance_windows AlloydbCluster#maintenance_windows}
	MaintenanceWindows interface{} `field:"optional" json:"maintenanceWindows" yaml:"maintenanceWindows"`
}

