package googlecomputefuturereservation


type GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesLocalSsds struct {
	// Specifies the size of the disk in base-2 GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#disk_size_gb GoogleComputeFutureReservation#disk_size_gb}
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME.
	//
	// The default is SCSI. Possible values: ["SCSI", "NVME"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#interface GoogleComputeFutureReservation#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
}

