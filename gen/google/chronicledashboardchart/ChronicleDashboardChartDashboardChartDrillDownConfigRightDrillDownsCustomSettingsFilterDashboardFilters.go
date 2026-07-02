package chronicledashboardchart

type ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#dashboard_filter_id ChronicleDashboardChart#dashboard_filter_id}.
	DashboardFilterId *string `field:"required" json:"dashboardFilterId" yaml:"dashboardFilterId"`
	// filter_operator_and_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#filter_operator_and_values ChronicleDashboardChart#filter_operator_and_values}
	FilterOperatorAndValues any `field:"required" json:"filterOperatorAndValues" yaml:"filterOperatorAndValues"`
}
