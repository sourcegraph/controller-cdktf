package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionConfigVariable struct {
	// Key for the configVariable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#key GoogleIntegrationConnectorsConnection#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Boolean Value of configVariable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#boolean_value GoogleIntegrationConnectorsConnection#boolean_value}
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// encryption_key_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#encryption_key_value GoogleIntegrationConnectorsConnection#encryption_key_value}
	EncryptionKeyValue *GoogleIntegrationConnectorsConnectionConfigVariableEncryptionKeyValue `field:"optional" json:"encryptionKeyValue" yaml:"encryptionKeyValue"`
	// Integer Value of configVariable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#integer_value GoogleIntegrationConnectorsConnection#integer_value}
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// secret_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#secret_value GoogleIntegrationConnectorsConnection#secret_value}
	SecretValue *GoogleIntegrationConnectorsConnectionConfigVariableSecretValue `field:"optional" json:"secretValue" yaml:"secretValue"`
	// String Value of configVariabley.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#string_value GoogleIntegrationConnectorsConnection#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

