package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterDashboardFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#dashboard_filter_id GoogleChronicleDashboardChart#dashboard_filter_id}.
	DashboardFilterId *string `field:"required" json:"dashboardFilterId" yaml:"dashboardFilterId"`
	// filter_operator_and_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#filter_operator_and_values GoogleChronicleDashboardChart#filter_operator_and_values}
	FilterOperatorAndValues interface{} `field:"required" json:"filterOperatorAndValues" yaml:"filterOperatorAndValues"`
}

