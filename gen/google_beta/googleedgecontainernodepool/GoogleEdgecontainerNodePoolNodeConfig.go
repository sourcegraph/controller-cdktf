package googleedgecontainernodepool


type GoogleEdgecontainerNodePoolNodeConfig struct {
	// "The Kubernetes node labels".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_edgecontainer_node_pool#labels GoogleEdgecontainerNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

