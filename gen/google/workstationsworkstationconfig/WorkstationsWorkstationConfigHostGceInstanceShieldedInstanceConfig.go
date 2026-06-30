package workstationsworkstationconfig

type WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig struct {
	// Whether the instance has integrity monitoring enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config#enable_integrity_monitoring WorkstationsWorkstationConfigA#enable_integrity_monitoring}
	EnableIntegrityMonitoring any `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Whether the instance has Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config#enable_secure_boot WorkstationsWorkstationConfigA#enable_secure_boot}
	EnableSecureBoot any `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Whether the instance has the vTPM enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config#enable_vtpm WorkstationsWorkstationConfigA#enable_vtpm}
	EnableVtpm any `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}
