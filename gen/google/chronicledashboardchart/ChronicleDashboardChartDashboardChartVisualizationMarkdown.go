package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationMarkdown struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#content ChronicleDashboardChart#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#properties ChronicleDashboardChart#properties}
	Properties *ChronicleDashboardChartDashboardChartVisualizationMarkdownProperties `field:"optional" json:"properties" yaml:"properties"`
}

