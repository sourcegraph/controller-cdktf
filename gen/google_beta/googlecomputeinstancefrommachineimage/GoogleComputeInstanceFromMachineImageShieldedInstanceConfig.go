package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageShieldedInstanceConfig struct {
	// Whether integrity monitoring is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#enable_integrity_monitoring GoogleComputeInstanceFromMachineImage#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Whether secure boot is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#enable_secure_boot GoogleComputeInstanceFromMachineImage#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Whether the instance uses vTPM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#enable_vtpm GoogleComputeInstanceFromMachineImage#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

