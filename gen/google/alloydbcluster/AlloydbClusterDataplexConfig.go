package alloydbcluster


type AlloydbClusterDataplexConfig struct {
	// Indicates whether Dataplex integration is enabled for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/alloydb_cluster#enabled AlloydbCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

