package googleassuredworkloadsworkload


type GoogleAssuredWorkloadsWorkloadKmsSettings struct {
	// Required.
	//
	// Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload#next_rotation_time GoogleAssuredWorkloadsWorkload#next_rotation_time}
	NextRotationTime *string `field:"required" json:"nextRotationTime" yaml:"nextRotationTime"`
	// Required.
	//
	// Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload#rotation_period GoogleAssuredWorkloadsWorkload#rotation_period}
	RotationPeriod *string `field:"required" json:"rotationPeriod" yaml:"rotationPeriod"`
}

