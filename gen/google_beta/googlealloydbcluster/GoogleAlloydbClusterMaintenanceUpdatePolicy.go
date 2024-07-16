package googlealloydbcluster


type GoogleAlloydbClusterMaintenanceUpdatePolicy struct {
	// maintenance_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#maintenance_windows GoogleAlloydbCluster#maintenance_windows}
	MaintenanceWindows interface{} `field:"optional" json:"maintenanceWindows" yaml:"maintenanceWindows"`
}

