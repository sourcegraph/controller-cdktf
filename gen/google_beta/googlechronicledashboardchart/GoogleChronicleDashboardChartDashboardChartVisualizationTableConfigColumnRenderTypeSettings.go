package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationTableConfigColumnRenderTypeSettings struct {
	// Possible values: ["RENDER_TYPE_UNSPECIFIED", "RENDER_TYPE_TEXT", "RENDER_TYPE_ICON", "RENDER_TYPE_ICON_AND_TEXT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#column_render_type GoogleChronicleDashboardChart#column_render_type}
	ColumnRenderType *string `field:"optional" json:"columnRenderType" yaml:"columnRenderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#field GoogleChronicleDashboardChart#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

