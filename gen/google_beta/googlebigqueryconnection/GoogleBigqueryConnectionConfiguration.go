package googlebigqueryconnection


type GoogleBigqueryConnectionConfiguration struct {
	// asset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#asset GoogleBigqueryConnection#asset}
	Asset *GoogleBigqueryConnectionConfigurationAsset `field:"required" json:"asset" yaml:"asset"`
	// The ID of the connector.
	//
	// Possible values include 'google-alloydb', 'google-cloudsql-mysql',
	// 'google-cloudsql-postgres', and other connector IDs supported by the BigQuery Connector framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#connector_id GoogleBigqueryConnection#connector_id}
	ConnectorId *string `field:"required" json:"connectorId" yaml:"connectorId"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#authentication GoogleBigqueryConnection#authentication}
	Authentication *GoogleBigqueryConnectionConfigurationAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#endpoint GoogleBigqueryConnection#endpoint}
	Endpoint *GoogleBigqueryConnectionConfigurationEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#network GoogleBigqueryConnection#network}
	Network *GoogleBigqueryConnectionConfigurationNetwork `field:"optional" json:"network" yaml:"network"`
}

