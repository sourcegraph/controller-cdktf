package googlecontainercluster


type GoogleContainerClusterMaintenancePolicy struct {
	// daily_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#daily_maintenance_window GoogleContainerCluster#daily_maintenance_window}
	DailyMaintenanceWindow *GoogleContainerClusterMaintenancePolicyDailyMaintenanceWindow `field:"optional" json:"dailyMaintenanceWindow" yaml:"dailyMaintenanceWindow"`
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#disruption_budget GoogleContainerCluster#disruption_budget}
	DisruptionBudget *GoogleContainerClusterMaintenancePolicyDisruptionBudget `field:"optional" json:"disruptionBudget" yaml:"disruptionBudget"`
	// maintenance_exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#maintenance_exclusion GoogleContainerCluster#maintenance_exclusion}
	MaintenanceExclusion interface{} `field:"optional" json:"maintenanceExclusion" yaml:"maintenanceExclusion"`
	// recurring_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#recurring_window GoogleContainerCluster#recurring_window}
	RecurringWindow *GoogleContainerClusterMaintenancePolicyRecurringWindow `field:"optional" json:"recurringWindow" yaml:"recurringWindow"`
}

