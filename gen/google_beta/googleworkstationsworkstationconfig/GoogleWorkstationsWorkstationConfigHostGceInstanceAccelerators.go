package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigHostGceInstanceAccelerators struct {
	// Number of accelerator cards exposed to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#count GoogleWorkstationsWorkstationConfigA#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Type of accelerator resource to attach to the instance, for example, "nvidia-tesla-p100".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#type GoogleWorkstationsWorkstationConfigA#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

