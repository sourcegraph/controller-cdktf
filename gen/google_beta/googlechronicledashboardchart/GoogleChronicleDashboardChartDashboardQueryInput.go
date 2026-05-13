package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardQueryInput struct {
	// relative_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#relative_time GoogleChronicleDashboardChart#relative_time}
	RelativeTime *GoogleChronicleDashboardChartDashboardQueryInputRelativeTime `field:"optional" json:"relativeTime" yaml:"relativeTime"`
	// time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#time_window GoogleChronicleDashboardChart#time_window}
	TimeWindow *GoogleChronicleDashboardChartDashboardQueryInputTimeWindow `field:"optional" json:"timeWindow" yaml:"timeWindow"`
}

