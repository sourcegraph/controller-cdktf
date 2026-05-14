package chronicledashboardchart


type ChronicleDashboardChartDashboardQueryInput struct {
	// relative_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#relative_time ChronicleDashboardChart#relative_time}
	RelativeTime *ChronicleDashboardChartDashboardQueryInputRelativeTime `field:"optional" json:"relativeTime" yaml:"relativeTime"`
	// time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#time_window ChronicleDashboardChart#time_window}
	TimeWindow *ChronicleDashboardChartDashboardQueryInputTimeWindow `field:"optional" json:"timeWindow" yaml:"timeWindow"`
}

