package googlecomputeresourcepolicy


type GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks struct {
	// The day of the week to create the snapshot.
	//
	// e.g. MONDAY Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#day GoogleComputeResourcePolicy#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// Time within the window to start the operations.
	//
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#start_time GoogleComputeResourcePolicy#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

