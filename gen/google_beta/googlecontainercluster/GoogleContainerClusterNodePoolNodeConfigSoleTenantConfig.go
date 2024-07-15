package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigSoleTenantConfig struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#node_affinity GoogleContainerCluster#node_affinity}
	NodeAffinity interface{} `field:"required" json:"nodeAffinity" yaml:"nodeAffinity"`
}

