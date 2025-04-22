package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileSalesforceProfile struct {
	// Domain for the Salesforce Org.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#domain GoogleDatastreamConnectionProfile#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// oauth2_client_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#oauth2_client_credentials GoogleDatastreamConnectionProfile#oauth2_client_credentials}
	Oauth2ClientCredentials *GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentials `field:"optional" json:"oauth2ClientCredentials" yaml:"oauth2ClientCredentials"`
	// user_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_datastream_connection_profile#user_credentials GoogleDatastreamConnectionProfile#user_credentials}
	UserCredentials *GoogleDatastreamConnectionProfileSalesforceProfileUserCredentials `field:"optional" json:"userCredentials" yaml:"userCredentials"`
}

