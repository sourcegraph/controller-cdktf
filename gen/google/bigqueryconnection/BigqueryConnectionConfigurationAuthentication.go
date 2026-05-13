package bigqueryconnection


type BigqueryConnectionConfigurationAuthentication struct {
	// username_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#username_password BigqueryConnection#username_password}
	UsernamePassword *BigqueryConnectionConfigurationAuthenticationUsernamePassword `field:"optional" json:"usernamePassword" yaml:"usernamePassword"`
}

