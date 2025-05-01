package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentials struct {
	// Client ID to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_connection_profile#client_id GoogleDatastreamConnectionProfile#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_connection_profile#client_secret GoogleDatastreamConnectionProfile#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A reference to a Secret Manager resource name storing the client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_connection_profile#secret_manager_stored_client_secret GoogleDatastreamConnectionProfile#secret_manager_stored_client_secret}
	SecretManagerStoredClientSecret *string `field:"optional" json:"secretManagerStoredClientSecret" yaml:"secretManagerStoredClientSecret"`
}

