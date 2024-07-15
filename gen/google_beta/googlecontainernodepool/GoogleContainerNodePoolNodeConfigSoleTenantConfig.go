package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigSoleTenantConfig struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#node_affinity GoogleContainerNodePool#node_affinity}
	NodeAffinity interface{} `field:"required" json:"nodeAffinity" yaml:"nodeAffinity"`
}

