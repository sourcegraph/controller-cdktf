package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfig struct {
	// data_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#data_settings GoogleChronicleDashboardChart#data_settings}
	DataSettings *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigDataSettings `field:"optional" json:"dataSettings" yaml:"dataSettings"`
	// map_position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#map_position GoogleChronicleDashboardChart#map_position}
	MapPosition *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigMapPosition `field:"optional" json:"mapPosition" yaml:"mapPosition"`
	// Possible values: ["PLOT_MODE_UNSPECIFIED", "PLOT_MODE_POINTS", "PLOT_MODE_HEATMAP", "PLOT_MODE_BOTH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#plot_mode GoogleChronicleDashboardChart#plot_mode}
	PlotMode *string `field:"optional" json:"plotMode" yaml:"plotMode"`
	// point_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#point_settings GoogleChronicleDashboardChart#point_settings}
	PointSettings *GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigPointSettings `field:"optional" json:"pointSettings" yaml:"pointSettings"`
}

