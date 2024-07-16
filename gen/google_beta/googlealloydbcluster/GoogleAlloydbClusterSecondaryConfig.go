package googlealloydbcluster


type GoogleAlloydbClusterSecondaryConfig struct {
	// Name of the primary cluster must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#primary_cluster_name GoogleAlloydbCluster#primary_cluster_name}
	PrimaryClusterName *string `field:"required" json:"primaryClusterName" yaml:"primaryClusterName"`
}

