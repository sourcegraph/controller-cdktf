package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationYAxes struct {
	// Possible values: ["VALUE", "CATEGORY", "TIME", "LOG"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#axis_type ChronicleDashboardChart#axis_type}
	AxisType *string `field:"optional" json:"axisType" yaml:"axisType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#display_name ChronicleDashboardChart#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#max ChronicleDashboardChart#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#min ChronicleDashboardChart#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

