package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigPointSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#color GoogleChronicleDashboardChart#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Possible values: ["POINT_SIZE_TYPE_UNSPECIFIED", "POINT_SIZE_TYPE_FIXED", "POINT_SIZE_TYPE_PROPORTIONAL_TO_SIZE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#point_size_type GoogleChronicleDashboardChart#point_size_type}
	PointSizeType *string `field:"optional" json:"pointSizeType" yaml:"pointSizeType"`
}

