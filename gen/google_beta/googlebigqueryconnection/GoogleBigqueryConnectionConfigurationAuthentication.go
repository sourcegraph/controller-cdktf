package googlebigqueryconnection


type GoogleBigqueryConnectionConfigurationAuthentication struct {
	// username_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#username_password GoogleBigqueryConnection#username_password}
	UsernamePassword *GoogleBigqueryConnectionConfigurationAuthenticationUsernamePassword `field:"optional" json:"usernamePassword" yaml:"usernamePassword"`
}

