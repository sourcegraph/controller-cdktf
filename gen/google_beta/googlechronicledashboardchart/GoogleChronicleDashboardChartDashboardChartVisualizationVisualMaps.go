package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationVisualMaps struct {
	// pieces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#pieces GoogleChronicleDashboardChart#pieces}
	Pieces interface{} `field:"optional" json:"pieces" yaml:"pieces"`
	// Possible values: ["VISUAL_MAP_TYPE_UNSPECIFIED", "CONTINUOUS", "PIECEWISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#visual_map_type GoogleChronicleDashboardChart#visual_map_type}
	VisualMapType *string `field:"optional" json:"visualMapType" yaml:"visualMapType"`
}

