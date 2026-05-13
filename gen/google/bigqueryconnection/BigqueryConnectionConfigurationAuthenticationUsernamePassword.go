package bigqueryconnection


type BigqueryConnectionConfigurationAuthenticationUsernamePassword struct {
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#password BigqueryConnection#password}
	Password *BigqueryConnectionConfigurationAuthenticationUsernamePasswordPassword `field:"required" json:"password" yaml:"password"`
	// Username for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#username BigqueryConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

