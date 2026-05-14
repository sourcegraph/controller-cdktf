package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig struct {
	// data_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#data_settings ChronicleDashboardChart#data_settings}
	DataSettings *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings `field:"optional" json:"dataSettings" yaml:"dataSettings"`
	// map_position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#map_position ChronicleDashboardChart#map_position}
	MapPosition *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition `field:"optional" json:"mapPosition" yaml:"mapPosition"`
	// Possible values: ["PLOT_MODE_UNSPECIFIED", "PLOT_MODE_POINTS", "PLOT_MODE_HEATMAP", "PLOT_MODE_BOTH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#plot_mode ChronicleDashboardChart#plot_mode}
	PlotMode *string `field:"optional" json:"plotMode" yaml:"plotMode"`
	// point_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#point_settings ChronicleDashboardChart#point_settings}
	PointSettings *ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigPointSettings `field:"optional" json:"pointSettings" yaml:"pointSettings"`
}

