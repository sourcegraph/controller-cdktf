package chronicledashboardchart


type ChronicleDashboardChartDashboardQuery struct {
	// The raw query string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#query ChronicleDashboardChart#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#input ChronicleDashboardChart#input}
	Input *ChronicleDashboardChartDashboardQueryInput `field:"optional" json:"input" yaml:"input"`
}

