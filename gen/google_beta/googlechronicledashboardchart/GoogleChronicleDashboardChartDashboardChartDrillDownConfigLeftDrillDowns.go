package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDowns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#display_name GoogleChronicleDashboardChart#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#id GoogleChronicleDashboardChart#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// custom_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#custom_settings GoogleChronicleDashboardChart#custom_settings}
	CustomSettings *GoogleChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettings `field:"optional" json:"customSettings" yaml:"customSettings"`
	// default_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#default_settings GoogleChronicleDashboardChart#default_settings}
	DefaultSettings *GoogleChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsDefaultSettings `field:"optional" json:"defaultSettings" yaml:"defaultSettings"`
}

