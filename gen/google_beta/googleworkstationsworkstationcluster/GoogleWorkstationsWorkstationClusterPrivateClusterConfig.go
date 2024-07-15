package googleworkstationsworkstationcluster


type GoogleWorkstationsWorkstationClusterPrivateClusterConfig struct {
	// Whether Workstations endpoint is private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#enable_private_endpoint GoogleWorkstationsWorkstationCluster#enable_private_endpoint}
	EnablePrivateEndpoint interface{} `field:"required" json:"enablePrivateEndpoint" yaml:"enablePrivateEndpoint"`
	// Additional project IDs that are allowed to attach to the workstation cluster's service attachment.
	//
	// By default, the workstation cluster's project and the VPC host project (if different) are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#allowed_projects GoogleWorkstationsWorkstationCluster#allowed_projects}
	AllowedProjects *[]*string `field:"optional" json:"allowedProjects" yaml:"allowedProjects"`
}

