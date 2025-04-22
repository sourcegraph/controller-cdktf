package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileSalesforceProfileUserCredentials struct {
	// Password of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#password GoogleDatastreamConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// A reference to a Secret Manager resource name storing the user's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#secret_manager_stored_password GoogleDatastreamConnectionProfile#secret_manager_stored_password}
	SecretManagerStoredPassword *string `field:"optional" json:"secretManagerStoredPassword" yaml:"secretManagerStoredPassword"`
	// A reference to a Secret Manager resource name storing the user's security token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#secret_manager_stored_security_token GoogleDatastreamConnectionProfile#secret_manager_stored_security_token}
	SecretManagerStoredSecurityToken *string `field:"optional" json:"secretManagerStoredSecurityToken" yaml:"secretManagerStoredSecurityToken"`
	// Security token of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#security_token GoogleDatastreamConnectionProfile#security_token}
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	// Username to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#username GoogleDatastreamConnectionProfile#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

