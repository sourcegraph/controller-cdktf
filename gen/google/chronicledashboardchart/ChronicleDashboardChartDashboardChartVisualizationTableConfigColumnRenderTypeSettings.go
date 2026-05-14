package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationTableConfigColumnRenderTypeSettings struct {
	// Possible values: ["RENDER_TYPE_UNSPECIFIED", "RENDER_TYPE_TEXT", "RENDER_TYPE_ICON", "RENDER_TYPE_ICON_AND_TEXT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#column_render_type ChronicleDashboardChart#column_render_type}
	ColumnRenderType *string `field:"optional" json:"columnRenderType" yaml:"columnRenderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#field ChronicleDashboardChart#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

