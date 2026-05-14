package chronicledashboardchart


type ChronicleDashboardChartChartLayout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#span_x ChronicleDashboardChart#span_x}.
	SpanX *float64 `field:"required" json:"spanX" yaml:"spanX"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#span_y ChronicleDashboardChart#span_y}.
	SpanY *float64 `field:"required" json:"spanY" yaml:"spanY"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#start_x ChronicleDashboardChart#start_x}.
	StartX *float64 `field:"optional" json:"startX" yaml:"startX"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_dashboard_chart#start_y ChronicleDashboardChart#start_y}.
	StartY *float64 `field:"optional" json:"startY" yaml:"startY"`
}

