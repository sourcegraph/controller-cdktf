package memorystoreinstance


type MemorystoreInstanceManagedBackupSource struct {
	// Example: 'projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/memorystore_instance#backup MemorystoreInstance#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

