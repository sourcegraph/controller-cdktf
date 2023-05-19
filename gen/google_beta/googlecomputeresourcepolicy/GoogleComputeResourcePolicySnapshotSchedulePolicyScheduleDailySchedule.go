package googlecomputeresourcepolicy


type GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule struct {
	// The number of days between snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_resource_policy#days_in_cycle GoogleComputeResourcePolicy#days_in_cycle}
	DaysInCycle *float64 `field:"required" json:"daysInCycle" yaml:"daysInCycle"`
	// This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00.
	//
	// For example,
	// both 13:00-5 and 08:00 are valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_resource_policy#start_time GoogleComputeResourcePolicy#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

