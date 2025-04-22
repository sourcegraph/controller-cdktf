package googlerediscluster


type GoogleRedisClusterCrossClusterReplicationConfigPrimaryCluster struct {
	// The full resource path of the primary cluster in the format: projects/{project}/locations/{region}/clusters/{cluster-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster#cluster GoogleRedisCluster#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
}

