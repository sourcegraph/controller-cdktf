package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionSslConfigAdditionalVariableEncryptionKeyValue struct {
	// The [KMS key name] with which the content of the Operation is encrypted.
	//
	// The expected
	// format: projects/* /locations/* /keyRings/* /cryptoKeys/*.
	// Will be empty string if google managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#kms_key_name GoogleIntegrationConnectorsConnection#kms_key_name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Type of Encription Key Possible values: ["GOOGLE_MANAGED", "CUSTOMER_MANAGED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#type GoogleIntegrationConnectorsConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

