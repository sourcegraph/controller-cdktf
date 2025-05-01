package googlerediscluster


type GoogleRedisClusterPersistenceConfigAofConfig struct {
	// Optional. Available fsync modes.
	//
	// - NO - Do not explicitly call fsync(). Rely on OS defaults.
	// - EVERYSEC - Call fsync() once per second in a background thread. A balance between performance and durability.
	// - ALWAYS - Call fsync() for earch write command. Possible values: ["APPEND_FSYNC_UNSPECIFIED", "NO", "EVERYSEC", "ALWAYS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_redis_cluster#append_fsync GoogleRedisCluster#append_fsync}
	AppendFsync *string `field:"optional" json:"appendFsync" yaml:"appendFsync"`
}

