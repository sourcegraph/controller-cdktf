package googlerediscluster


type GoogleRedisClusterZoneDistributionConfig struct {
	// Immutable.
	//
	// The mode for zone distribution for Memorystore Redis cluster.
	// If not provided, MULTI_ZONE will be used as default Possible values: ["MULTI_ZONE", "SINGLE_ZONE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_redis_cluster#mode GoogleRedisCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Immutable. The zone for single zone Memorystore Redis cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_redis_cluster#zone GoogleRedisCluster#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

