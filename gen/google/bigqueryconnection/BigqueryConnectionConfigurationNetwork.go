package bigqueryconnection


type BigqueryConnectionConfigurationNetwork struct {
	// private_service_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#private_service_connect BigqueryConnection#private_service_connect}
	PrivateServiceConnect *BigqueryConnectionConfigurationNetworkPrivateServiceConnect `field:"optional" json:"privateServiceConnect" yaml:"privateServiceConnect"`
}

