package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationTableConfigColumnTooltipSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#field GoogleChronicleDashboardChart#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#cell_tooltip_text GoogleChronicleDashboardChart#cell_tooltip_text}.
	CellTooltipText *string `field:"optional" json:"cellTooltipText" yaml:"cellTooltipText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#header_tooltip_text GoogleChronicleDashboardChart#header_tooltip_text}.
	HeaderTooltipText *string `field:"optional" json:"headerTooltipText" yaml:"headerTooltipText"`
}

