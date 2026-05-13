package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardQuery struct {
	// The raw query string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#query GoogleChronicleDashboardChart#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#input GoogleChronicleDashboardChart#input}
	Input *GoogleChronicleDashboardChartDashboardQueryInput `field:"optional" json:"input" yaml:"input"`
}

