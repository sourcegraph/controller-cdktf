package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationButtonProperties struct {
	// Possible values: ["BUTTON_STYLE_UNSPECIFIED", "BUTTON_STYLE_FILLED", "BUTTON_STYLE_OUTLINED", "BUTTON_STYLE_TRANSPARENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#button_style GoogleChronicleDashboardChart#button_style}
	ButtonStyle *string `field:"optional" json:"buttonStyle" yaml:"buttonStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#color GoogleChronicleDashboardChart#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
}

