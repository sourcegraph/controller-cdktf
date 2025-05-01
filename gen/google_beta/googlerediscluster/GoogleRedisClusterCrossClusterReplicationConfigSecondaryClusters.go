package googlerediscluster


type GoogleRedisClusterCrossClusterReplicationConfigSecondaryClusters struct {
	// The full resource path of the secondary cluster in the format: projects/{project}/locations/{region}/clusters/{cluster-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_redis_cluster#cluster GoogleRedisCluster#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
}

