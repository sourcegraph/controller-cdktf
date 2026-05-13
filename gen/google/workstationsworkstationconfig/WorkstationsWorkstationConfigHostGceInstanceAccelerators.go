package workstationsworkstationconfig


type WorkstationsWorkstationConfigHostGceInstanceAccelerators struct {
	// Number of accelerator cards exposed to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#count WorkstationsWorkstationConfigA#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Type of accelerator resource to attach to the instance, for example, "nvidia-tesla-p100".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#type WorkstationsWorkstationConfigA#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

