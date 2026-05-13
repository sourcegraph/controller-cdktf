package googlechronicledashboardchart


type GoogleChronicleDashboardChartDashboardQueryInputRelativeTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#start_time_val GoogleChronicleDashboardChart#start_time_val}.
	StartTimeVal *string `field:"required" json:"startTimeVal" yaml:"startTimeVal"`
	// The time unit for the relative range. Possible values: ["SECOND", "MINUTE", "HOUR", "DAY", "WEEK", "MONTH", "YEAR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#time_unit GoogleChronicleDashboardChart#time_unit}
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
}

