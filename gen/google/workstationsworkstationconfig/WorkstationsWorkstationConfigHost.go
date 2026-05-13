package workstationsworkstationconfig


type WorkstationsWorkstationConfigHost struct {
	// gce_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#gce_instance WorkstationsWorkstationConfigA#gce_instance}
	GceInstance *WorkstationsWorkstationConfigHostGceInstance `field:"optional" json:"gceInstance" yaml:"gceInstance"`
}

