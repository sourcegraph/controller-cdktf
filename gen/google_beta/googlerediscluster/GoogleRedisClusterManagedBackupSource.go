package googlerediscluster


type GoogleRedisClusterManagedBackupSource struct {
	// Example: //redis.googleapis.com/projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup} A shorter version (without the prefix) of the backup name is also supported, like projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backupId}. In this case, it assumes the backup is under redis.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_redis_cluster#backup GoogleRedisCluster#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

