package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#new_tab GoogleChronicleDashboardChart#new_tab}.
	NewTab interface{} `field:"required" json:"newTab" yaml:"newTab"`
	// external_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#external_link GoogleChronicleDashboardChart#external_link}
	ExternalLink *GoogleChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLink `field:"optional" json:"externalLink" yaml:"externalLink"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#filter GoogleChronicleDashboardChart#filter}
	Filter *GoogleChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilter `field:"optional" json:"filter" yaml:"filter"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#query GoogleChronicleDashboardChart#query}
	Query *GoogleChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQuery `field:"optional" json:"query" yaml:"query"`
}

