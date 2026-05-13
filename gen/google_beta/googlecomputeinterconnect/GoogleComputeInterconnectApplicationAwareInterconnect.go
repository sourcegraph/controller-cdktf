package googlecomputeinterconnect


type GoogleComputeInterconnectApplicationAwareInterconnect struct {
	// bandwidth_percentage_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect#bandwidth_percentage_policy GoogleComputeInterconnect#bandwidth_percentage_policy}
	BandwidthPercentagePolicy *GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy `field:"optional" json:"bandwidthPercentagePolicy" yaml:"bandwidthPercentagePolicy"`
	// A description for the AAI profile on this interconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect#profile_description GoogleComputeInterconnect#profile_description}
	ProfileDescription *string `field:"optional" json:"profileDescription" yaml:"profileDescription"`
	// shape_average_percentage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect#shape_average_percentage GoogleComputeInterconnect#shape_average_percentage}
	ShapeAveragePercentage interface{} `field:"optional" json:"shapeAveragePercentage" yaml:"shapeAveragePercentage"`
	// strict_priority_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect#strict_priority_policy GoogleComputeInterconnect#strict_priority_policy}
	StrictPriorityPolicy *GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy `field:"optional" json:"strictPriorityPolicy" yaml:"strictPriorityPolicy"`
}

