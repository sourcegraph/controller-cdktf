package googlelustreinstance


type GoogleLustreInstanceMaintenancePolicy struct {
	// weekly_maintenance_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#weekly_maintenance_windows GoogleLustreInstance#weekly_maintenance_windows}
	WeeklyMaintenanceWindows *GoogleLustreInstanceMaintenancePolicyWeeklyMaintenanceWindows `field:"required" json:"weeklyMaintenanceWindows" yaml:"weeklyMaintenanceWindows"`
	// maintenance_exclusion_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#maintenance_exclusion_window GoogleLustreInstance#maintenance_exclusion_window}
	MaintenanceExclusionWindow *GoogleLustreInstanceMaintenancePolicyMaintenanceExclusionWindow `field:"optional" json:"maintenanceExclusionWindow" yaml:"maintenanceExclusionWindow"`
}

