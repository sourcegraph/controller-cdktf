package googlerediscluster


type GoogleRedisClusterPersistenceConfig struct {
	// aof_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster#aof_config GoogleRedisCluster#aof_config}
	AofConfig *GoogleRedisClusterPersistenceConfigAofConfig `field:"optional" json:"aofConfig" yaml:"aofConfig"`
	// Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
	//
	// - DISABLED: 	Persistence (both backup and restore) is disabled for the cluster.
	// - RDB: RDB based Persistence is enabled.
	// - AOF: AOF based Persistence is enabled. Possible values: ["PERSISTENCE_MODE_UNSPECIFIED", "DISABLED", "RDB", "AOF"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster#mode GoogleRedisCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// rdb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster#rdb_config GoogleRedisCluster#rdb_config}
	RdbConfig *GoogleRedisClusterPersistenceConfigRdbConfig `field:"optional" json:"rdbConfig" yaml:"rdbConfig"`
}

