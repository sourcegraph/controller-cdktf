package googlecomputereservation


type GoogleComputeReservationSpecificReservationInstancePropertiesLocalSsds struct {
	// The size of the disk in base-2 GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#disk_size_gb GoogleComputeReservation#disk_size_gb}
	DiskSizeGb *float64 `field:"required" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The disk interface to use for attaching this disk. Default value: "SCSI" Possible values: ["SCSI", "NVME"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_reservation#interface GoogleComputeReservation#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
}

