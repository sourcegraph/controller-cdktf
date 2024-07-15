package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig struct {
	// Whether the instance has integrity monitoring enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#enable_integrity_monitoring GoogleWorkstationsWorkstationConfigA#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Whether the instance has Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#enable_secure_boot GoogleWorkstationsWorkstationConfigA#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Whether the instance has the vTPM enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#enable_vtpm GoogleWorkstationsWorkstationConfigA#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

