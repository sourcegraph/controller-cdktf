package googlecontainercluster


type GoogleContainerClusterDatabaseEncryption struct {
	// ENCRYPTED, ALL_OBJECTS_ENCRYPTION_ENABLED or DECRYPTED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#state GoogleContainerCluster#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The key to use to encrypt/decrypt secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#key_name GoogleContainerCluster#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
}

