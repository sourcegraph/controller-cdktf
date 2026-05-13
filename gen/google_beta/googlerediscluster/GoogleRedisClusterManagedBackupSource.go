package googlerediscluster


type GoogleRedisClusterManagedBackupSource struct {
	// Example: 'projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_redis_cluster#backup GoogleRedisCluster#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

