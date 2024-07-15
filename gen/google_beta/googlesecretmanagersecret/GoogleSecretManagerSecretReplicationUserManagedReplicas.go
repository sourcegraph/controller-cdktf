package googlesecretmanagersecret


type GoogleSecretManagerSecretReplicationUserManagedReplicas struct {
	// The canonical IDs of the location to replicate data. For example: "us-east1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#location GoogleSecretManagerSecret#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// customer_managed_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#customer_managed_encryption GoogleSecretManagerSecret#customer_managed_encryption}
	CustomerManagedEncryption *GoogleSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption `field:"optional" json:"customerManagedEncryption" yaml:"customerManagedEncryption"`
}

