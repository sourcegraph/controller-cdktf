package chronicledashboardchart


type ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#new_tab ChronicleDashboardChart#new_tab}.
	NewTab interface{} `field:"required" json:"newTab" yaml:"newTab"`
	// external_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#external_link ChronicleDashboardChart#external_link}
	ExternalLink *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLink `field:"optional" json:"externalLink" yaml:"externalLink"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#filter ChronicleDashboardChart#filter}
	Filter *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilter `field:"optional" json:"filter" yaml:"filter"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#query ChronicleDashboardChart#query}
	Query *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQuery `field:"optional" json:"query" yaml:"query"`
}

