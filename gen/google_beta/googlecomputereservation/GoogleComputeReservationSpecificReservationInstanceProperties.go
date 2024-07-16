package googlecomputereservation


type GoogleComputeReservationSpecificReservationInstanceProperties struct {
	// The name of the machine type to reserve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#machine_type GoogleComputeReservation#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// guest_accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#guest_accelerators GoogleComputeReservation#guest_accelerators}
	GuestAccelerators interface{} `field:"optional" json:"guestAccelerators" yaml:"guestAccelerators"`
	// local_ssds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#local_ssds GoogleComputeReservation#local_ssds}
	LocalSsds interface{} `field:"optional" json:"localSsds" yaml:"localSsds"`
	// The minimum CPU platform for the reservation.
	//
	// For example,
	// '"Intel Skylake"'. See
	// the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	// for information on available CPU platforms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#min_cpu_platform GoogleComputeReservation#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
}

