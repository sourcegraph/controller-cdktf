package lustreinstance


type LustreInstanceMaintenancePolicy struct {
	// weekly_maintenance_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#weekly_maintenance_windows LustreInstance#weekly_maintenance_windows}
	WeeklyMaintenanceWindows *LustreInstanceMaintenancePolicyWeeklyMaintenanceWindows `field:"required" json:"weeklyMaintenanceWindows" yaml:"weeklyMaintenanceWindows"`
	// maintenance_exclusion_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#maintenance_exclusion_window LustreInstance#maintenance_exclusion_window}
	MaintenanceExclusionWindow *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow `field:"optional" json:"maintenanceExclusionWindow" yaml:"maintenanceExclusionWindow"`
}

