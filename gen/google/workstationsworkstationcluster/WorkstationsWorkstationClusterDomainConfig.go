package workstationsworkstationcluster


type WorkstationsWorkstationClusterDomainConfig struct {
	// Domain used by Workstations for HTTP ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_cluster#domain WorkstationsWorkstationCluster#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

