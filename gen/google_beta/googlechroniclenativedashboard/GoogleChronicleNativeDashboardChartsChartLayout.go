package googlechroniclenativedashboard


type GoogleChronicleNativeDashboardChartsChartLayout struct {
	// The number of columns the chart spans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#span_x GoogleChronicleNativeDashboard#span_x}
	SpanX *float64 `field:"required" json:"spanX" yaml:"spanX"`
	// The number of rows the chart spans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#span_y GoogleChronicleNativeDashboard#span_y}
	SpanY *float64 `field:"required" json:"spanY" yaml:"spanY"`
	// The starting X coordinate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#start_x GoogleChronicleNativeDashboard#start_x}
	StartX *float64 `field:"optional" json:"startX" yaml:"startX"`
	// The starting Y coordinate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_native_dashboard#start_y GoogleChronicleNativeDashboard#start_y}
	StartY *float64 `field:"optional" json:"startY" yaml:"startY"`
}

