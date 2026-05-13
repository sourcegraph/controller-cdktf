package googlealloydbcluster


type GoogleAlloydbClusterDataplexConfig struct {
	// Indicates whether Dataplex integration is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_alloydb_cluster#enabled GoogleAlloydbCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

