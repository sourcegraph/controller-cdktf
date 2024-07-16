package googleworkstationsworkstationcluster


type GoogleWorkstationsWorkstationClusterDomainConfig struct {
	// Domain used by Workstations for HTTP ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#domain GoogleWorkstationsWorkstationCluster#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

