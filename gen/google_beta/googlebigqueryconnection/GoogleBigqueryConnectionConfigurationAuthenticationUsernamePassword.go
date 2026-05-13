package googlebigqueryconnection


type GoogleBigqueryConnectionConfigurationAuthenticationUsernamePassword struct {
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#password GoogleBigqueryConnection#password}
	Password *GoogleBigqueryConnectionConfigurationAuthenticationUsernamePasswordPassword `field:"required" json:"password" yaml:"password"`
	// Username for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#username GoogleBigqueryConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

