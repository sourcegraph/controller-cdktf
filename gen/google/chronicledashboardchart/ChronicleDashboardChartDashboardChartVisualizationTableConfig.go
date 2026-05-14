package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationTableConfig struct {
	// column_render_type_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#column_render_type_settings ChronicleDashboardChart#column_render_type_settings}
	ColumnRenderTypeSettings interface{} `field:"optional" json:"columnRenderTypeSettings" yaml:"columnRenderTypeSettings"`
	// column_tooltip_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#column_tooltip_settings ChronicleDashboardChart#column_tooltip_settings}
	ColumnTooltipSettings interface{} `field:"optional" json:"columnTooltipSettings" yaml:"columnTooltipSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#enable_text_wrap ChronicleDashboardChart#enable_text_wrap}.
	EnableTextWrap interface{} `field:"optional" json:"enableTextWrap" yaml:"enableTextWrap"`
}

