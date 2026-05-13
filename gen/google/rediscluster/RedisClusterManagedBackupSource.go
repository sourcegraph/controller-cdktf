package rediscluster


type RedisClusterManagedBackupSource struct {
	// Example: 'projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/redis_cluster#backup RedisCluster#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

