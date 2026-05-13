package googlecomputefuturereservation


type GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties struct {
	// guest_accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#guest_accelerators GoogleComputeFutureReservation#guest_accelerators}
	GuestAccelerators interface{} `field:"optional" json:"guestAccelerators" yaml:"guestAccelerators"`
	// local_ssds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#local_ssds GoogleComputeFutureReservation#local_ssds}
	LocalSsds interface{} `field:"optional" json:"localSsds" yaml:"localSsds"`
	// An opaque location hint used to place the allocation close to other resources.
	//
	// This field is for use by internal tools that use the public API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#location_hint GoogleComputeFutureReservation#location_hint}
	LocationHint *string `field:"optional" json:"locationHint" yaml:"locationHint"`
	// Specifies type of machine (name only) which has fixed number of vCPUs and fixed amount of memory.
	//
	// This also includes specifying custom machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#machine_type GoogleComputeFutureReservation#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Specifies the number of hours after reservation creation where instances using the reservation won't be scheduled for maintenance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#maintenance_freeze_duration_hours GoogleComputeFutureReservation#maintenance_freeze_duration_hours}
	MaintenanceFreezeDurationHours *float64 `field:"optional" json:"maintenanceFreezeDurationHours" yaml:"maintenanceFreezeDurationHours"`
	// Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC Possible values: ["PERIODIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#maintenance_interval GoogleComputeFutureReservation#maintenance_interval}
	MaintenanceInterval *string `field:"optional" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// Minimum cpu platform the reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#min_cpu_platform GoogleComputeFutureReservation#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
}

