package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationTableConfig struct {
	// column_render_type_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#column_render_type_settings GoogleChronicleDashboardChart#column_render_type_settings}
	ColumnRenderTypeSettings interface{} `field:"optional" json:"columnRenderTypeSettings" yaml:"columnRenderTypeSettings"`
	// column_tooltip_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#column_tooltip_settings GoogleChronicleDashboardChart#column_tooltip_settings}
	ColumnTooltipSettings interface{} `field:"optional" json:"columnTooltipSettings" yaml:"columnTooltipSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#enable_text_wrap GoogleChronicleDashboardChart#enable_text_wrap}.
	EnableTextWrap interface{} `field:"optional" json:"enableTextWrap" yaml:"enableTextWrap"`
}

