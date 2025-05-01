package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigAllowedPorts struct {
	// Starting port number for the current range of ports.
	//
	// Valid ports are 22, 80, and ports within the range 1024-65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_workstations_workstation_config#first GoogleWorkstationsWorkstationConfigA#first}
	First *float64 `field:"optional" json:"first" yaml:"first"`
	// Ending port number for the current range of ports.
	//
	// Valid ports are 22, 80, and ports within the range 1024-65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_workstations_workstation_config#last GoogleWorkstationsWorkstationConfigA#last}
	Last *float64 `field:"optional" json:"last" yaml:"last"`
}

