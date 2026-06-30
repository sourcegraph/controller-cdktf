package chronicledashboardchart

type ChronicleDashboardChartDashboardChartVisualizationVisualMaps struct {
	// pieces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#pieces ChronicleDashboardChart#pieces}
	Pieces any `field:"optional" json:"pieces" yaml:"pieces"`
	// Possible values: ["VISUAL_MAP_TYPE_UNSPECIFIED", "CONTINUOUS", "PIECEWISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#visual_map_type ChronicleDashboardChart#visual_map_type}
	VisualMapType *string `field:"optional" json:"visualMapType" yaml:"visualMapType"`
}
