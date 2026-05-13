package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationButton struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#hyperlink GoogleChronicleDashboardChart#hyperlink}.
	Hyperlink *string `field:"required" json:"hyperlink" yaml:"hyperlink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#label GoogleChronicleDashboardChart#label}.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#description GoogleChronicleDashboardChart#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#new_tab GoogleChronicleDashboardChart#new_tab}.
	NewTab interface{} `field:"optional" json:"newTab" yaml:"newTab"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#properties GoogleChronicleDashboardChart#properties}
	Properties *GoogleChronicleDashboardChartDashboardChartVisualizationButtonProperties `field:"optional" json:"properties" yaml:"properties"`
}

