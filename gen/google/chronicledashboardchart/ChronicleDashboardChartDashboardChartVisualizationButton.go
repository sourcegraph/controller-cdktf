package chronicledashboardchart

type ChronicleDashboardChartDashboardChartVisualizationButton struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#hyperlink ChronicleDashboardChart#hyperlink}.
	Hyperlink *string `field:"required" json:"hyperlink" yaml:"hyperlink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#label ChronicleDashboardChart#label}.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#description ChronicleDashboardChart#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#new_tab ChronicleDashboardChart#new_tab}.
	NewTab any `field:"optional" json:"newTab" yaml:"newTab"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#properties ChronicleDashboardChart#properties}
	Properties *ChronicleDashboardChartDashboardChartVisualizationButtonProperties `field:"optional" json:"properties" yaml:"properties"`
}
