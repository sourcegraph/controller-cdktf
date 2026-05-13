package workstationsworkstationconfig


type WorkstationsWorkstationConfigAllowedPorts struct {
	// Starting port number for the current range of ports.
	//
	// Valid ports are 22, 80, and ports within the range 1024-65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#first WorkstationsWorkstationConfigA#first}
	First *float64 `field:"optional" json:"first" yaml:"first"`
	// Ending port number for the current range of ports.
	//
	// Valid ports are 22, 80, and ports within the range 1024-65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#last WorkstationsWorkstationConfigA#last}
	Last *float64 `field:"optional" json:"last" yaml:"last"`
}

