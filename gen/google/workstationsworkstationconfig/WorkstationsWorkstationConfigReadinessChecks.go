package workstationsworkstationconfig


type WorkstationsWorkstationConfigReadinessChecks struct {
	// Path to which the request should be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#path WorkstationsWorkstationConfigA#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Port to which the request should be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#port WorkstationsWorkstationConfigA#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

