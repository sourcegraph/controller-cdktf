package containercluster


type ContainerClusterMaintenancePolicy struct {
	// daily_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#daily_maintenance_window ContainerCluster#daily_maintenance_window}
	DailyMaintenanceWindow *ContainerClusterMaintenancePolicyDailyMaintenanceWindow `field:"optional" json:"dailyMaintenanceWindow" yaml:"dailyMaintenanceWindow"`
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#disruption_budget ContainerCluster#disruption_budget}
	DisruptionBudget *ContainerClusterMaintenancePolicyDisruptionBudget `field:"optional" json:"disruptionBudget" yaml:"disruptionBudget"`
	// maintenance_exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#maintenance_exclusion ContainerCluster#maintenance_exclusion}
	MaintenanceExclusion interface{} `field:"optional" json:"maintenanceExclusion" yaml:"maintenanceExclusion"`
	// recurring_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#recurring_window ContainerCluster#recurring_window}
	RecurringWindow *ContainerClusterMaintenancePolicyRecurringWindow `field:"optional" json:"recurringWindow" yaml:"recurringWindow"`
}

