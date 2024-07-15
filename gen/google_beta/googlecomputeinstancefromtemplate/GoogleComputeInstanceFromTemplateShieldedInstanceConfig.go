package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateShieldedInstanceConfig struct {
	// Whether integrity monitoring is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#enable_integrity_monitoring GoogleComputeInstanceFromTemplate#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Whether secure boot is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#enable_secure_boot GoogleComputeInstanceFromTemplate#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Whether the instance uses vTPM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#enable_vtpm GoogleComputeInstanceFromTemplate#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

