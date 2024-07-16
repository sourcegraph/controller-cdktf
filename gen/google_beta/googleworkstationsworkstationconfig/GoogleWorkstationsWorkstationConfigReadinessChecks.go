package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigReadinessChecks struct {
	// Path to which the request should be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#path GoogleWorkstationsWorkstationConfigA#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Port to which the request should be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#port GoogleWorkstationsWorkstationConfigA#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

