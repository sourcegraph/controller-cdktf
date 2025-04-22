package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageServiceAccount struct {
	// A list of service scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#scopes GoogleComputeInstanceFromMachineImage#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account e-mail address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image#email GoogleComputeInstanceFromMachineImage#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

