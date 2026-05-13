package googlecomputeinterconnect


type GoogleComputeInterconnectApplicationAwareInterconnectShapeAveragePercentage struct {
	// Bandwidth percentage for a specific traffic class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect#percentage GoogleComputeInterconnect#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
	// Enum representing the various traffic classes offered by AAI.
	//
	// Default value: "TC_UNSPECIFIED" Possible values: ["TC_UNSPECIFIED", "TC1", "TC2", "TC3", "TC4", "TC5", "TC6"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect#traffic_class GoogleComputeInterconnect#traffic_class}
	TrafficClass *string `field:"optional" json:"trafficClass" yaml:"trafficClass"`
}

