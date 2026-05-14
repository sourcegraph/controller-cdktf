package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationTableConfigColumnTooltipSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#field ChronicleDashboardChart#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#cell_tooltip_text ChronicleDashboardChart#cell_tooltip_text}.
	CellTooltipText *string `field:"optional" json:"cellTooltipText" yaml:"cellTooltipText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#header_tooltip_text ChronicleDashboardChart#header_tooltip_text}.
	HeaderTooltipText *string `field:"optional" json:"headerTooltipText" yaml:"headerTooltipText"`
}

