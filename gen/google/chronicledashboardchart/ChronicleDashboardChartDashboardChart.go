package chronicledashboardchart


type ChronicleDashboardChartDashboardChart struct {
	// Display name/Title of the dashboardChart visible to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#display_name ChronicleDashboardChart#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// visualization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#visualization ChronicleDashboardChart#visualization}
	Visualization *ChronicleDashboardChartDashboardChartVisualization `field:"required" json:"visualization" yaml:"visualization"`
	// chart_datasource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#chart_datasource ChronicleDashboardChart#chart_datasource}
	ChartDatasource *ChronicleDashboardChartDashboardChartChartDatasource `field:"optional" json:"chartDatasource" yaml:"chartDatasource"`
	// Description of the dashboardChart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#description ChronicleDashboardChart#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// drill_down_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#drill_down_config ChronicleDashboardChart#drill_down_config}
	DrillDownConfig *ChronicleDashboardChartDashboardChartDrillDownConfig `field:"optional" json:"drillDownConfig" yaml:"drillDownConfig"`
	// Type of tile (e.g., visualization, button, markdown). Possible values: ["TILE_TYPE_UNSPECIFIED", "TILE_TYPE_VISUALIZATION", "TILE_TYPE_BUTTON", "TILE_TYPE_MARKDOWN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#tile_type ChronicleDashboardChart#tile_type}
	TileType *string `field:"optional" json:"tileType" yaml:"tileType"`
}

