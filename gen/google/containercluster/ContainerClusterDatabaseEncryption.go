package containercluster


type ContainerClusterDatabaseEncryption struct {
	// ENCRYPTED, ALL_OBJECTS_ENCRYPTION_ENABLED or DECRYPTED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#state ContainerCluster#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The key to use to encrypt/decrypt secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#key_name ContainerCluster#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
}

