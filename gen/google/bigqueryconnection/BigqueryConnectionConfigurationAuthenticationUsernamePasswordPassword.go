package bigqueryconnection


type BigqueryConnectionConfigurationAuthenticationUsernamePasswordPassword struct {
	// The plaintext password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#plaintext BigqueryConnection#plaintext}
	Plaintext *string `field:"required" json:"plaintext" yaml:"plaintext"`
}

