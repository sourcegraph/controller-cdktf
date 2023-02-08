package googlecomputeresourcepolicy


type GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule struct {
	// Specifies the frequency for the operation, using the unix-cron format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_resource_policy#schedule GoogleComputeResourcePolicy#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
}

