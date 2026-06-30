package chronicledashboardchart

type ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#new_tab ChronicleDashboardChart#new_tab}.
	NewTab any `field:"required" json:"newTab" yaml:"newTab"`
	// external_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#external_link ChronicleDashboardChart#external_link}
	ExternalLink *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLink `field:"optional" json:"externalLink" yaml:"externalLink"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#filter ChronicleDashboardChart#filter}
	Filter *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#left_click_column ChronicleDashboardChart#left_click_column}.
	LeftClickColumn *string `field:"optional" json:"leftClickColumn" yaml:"leftClickColumn"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#query ChronicleDashboardChart#query}
	Query *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQuery `field:"optional" json:"query" yaml:"query"`
}
