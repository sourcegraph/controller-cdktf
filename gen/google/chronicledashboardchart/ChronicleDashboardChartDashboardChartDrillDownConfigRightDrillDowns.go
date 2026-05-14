package chronicledashboardchart


type ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDowns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#display_name ChronicleDashboardChart#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#id ChronicleDashboardChart#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// custom_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#custom_settings ChronicleDashboardChart#custom_settings}
	CustomSettings *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings `field:"optional" json:"customSettings" yaml:"customSettings"`
	// default_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#default_settings ChronicleDashboardChart#default_settings}
	DefaultSettings *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsDefaultSettings `field:"optional" json:"defaultSettings" yaml:"defaultSettings"`
}

