package googlecomputedisk


type GoogleComputeDiskGuestOsFeatures struct {
	// The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_disk#type GoogleComputeDisk#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

