package googlealloydbinstance


type GoogleAlloydbInstanceReadPoolConfig struct {
	// Read capacity, i.e. number of nodes in a read pool instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_instance#node_count GoogleAlloydbInstance#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
}

