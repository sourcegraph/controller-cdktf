package googlecomputeimage


type GoogleComputeImageGuestOsFeatures struct {
	// The type of supported feature.
	//
	// Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: ["MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image#type GoogleComputeImage#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

