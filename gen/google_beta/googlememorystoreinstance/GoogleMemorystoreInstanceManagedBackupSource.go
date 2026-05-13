package googlememorystoreinstance


type GoogleMemorystoreInstanceManagedBackupSource struct {
	// Example: 'projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_memorystore_instance#backup GoogleMemorystoreInstance#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

