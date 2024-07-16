package edgecontainernodepool


type EdgecontainerNodePoolNodeConfig struct {
	// "The Kubernetes node labels".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/edgecontainer_node_pool#labels EdgecontainerNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

