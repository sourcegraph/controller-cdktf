package googlealloydbcluster


type GoogleAlloydbClusterMaintenanceUpdatePolicyMaintenanceWindows struct {
	// Preferred day of the week for maintenance, e.g. MONDAY, TUESDAY, etc. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#day GoogleAlloydbCluster#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#start_time GoogleAlloydbCluster#start_time}
	StartTime *GoogleAlloydbClusterMaintenanceUpdatePolicyMaintenanceWindowsStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

