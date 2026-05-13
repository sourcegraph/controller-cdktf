package googlebigqueryconnection


type GoogleBigqueryConnectionConfigurationNetwork struct {
	// private_service_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#private_service_connect GoogleBigqueryConnection#private_service_connect}
	PrivateServiceConnect *GoogleBigqueryConnectionConfigurationNetworkPrivateServiceConnect `field:"optional" json:"privateServiceConnect" yaml:"privateServiceConnect"`
}

