package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChart struct {
	// Display name/Title of the dashboardChart visible to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#display_name GoogleChronicleDashboardChart#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// visualization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#visualization GoogleChronicleDashboardChart#visualization}
	Visualization *GoogleChronicleDashboardChartDashboardChartVisualization `field:"required" json:"visualization" yaml:"visualization"`
	// chart_datasource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#chart_datasource GoogleChronicleDashboardChart#chart_datasource}
	ChartDatasource *GoogleChronicleDashboardChartDashboardChartChartDatasource `field:"optional" json:"chartDatasource" yaml:"chartDatasource"`
	// Description of the dashboardChart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#description GoogleChronicleDashboardChart#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// drill_down_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#drill_down_config GoogleChronicleDashboardChart#drill_down_config}
	DrillDownConfig *GoogleChronicleDashboardChartDashboardChartDrillDownConfig `field:"optional" json:"drillDownConfig" yaml:"drillDownConfig"`
	// Type of tile (e.g., visualization, button, markdown). Possible values: ["TILE_TYPE_UNSPECIFIED", "TILE_TYPE_VISUALIZATION", "TILE_TYPE_BUTTON", "TILE_TYPE_MARKDOWN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#tile_type GoogleChronicleDashboardChart#tile_type}
	TileType *string `field:"optional" json:"tileType" yaml:"tileType"`
}

