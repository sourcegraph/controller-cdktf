package googlechronicledashboardchart


type GoogleChronicleDashboardChartChartLayout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#span_x GoogleChronicleDashboardChart#span_x}.
	SpanX *float64 `field:"required" json:"spanX" yaml:"spanX"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#span_y GoogleChronicleDashboardChart#span_y}.
	SpanY *float64 `field:"required" json:"spanY" yaml:"spanY"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#start_x GoogleChronicleDashboardChart#start_x}.
	StartX *float64 `field:"optional" json:"startX" yaml:"startX"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_dashboard_chart#start_y GoogleChronicleDashboardChart#start_y}.
	StartY *float64 `field:"optional" json:"startY" yaml:"startY"`
}

