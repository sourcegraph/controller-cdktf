package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardChartVisualizationYAxes struct {
	// Possible values: ["VALUE", "CATEGORY", "TIME", "LOG"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#axis_type GoogleChronicleDashboardChart#axis_type}
	AxisType *string `field:"optional" json:"axisType" yaml:"axisType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#display_name GoogleChronicleDashboardChart#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#max GoogleChronicleDashboardChart#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#min GoogleChronicleDashboardChart#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

