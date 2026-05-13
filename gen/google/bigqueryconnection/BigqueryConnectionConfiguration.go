package bigqueryconnection


type BigqueryConnectionConfiguration struct {
	// asset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#asset BigqueryConnection#asset}
	Asset *BigqueryConnectionConfigurationAsset `field:"required" json:"asset" yaml:"asset"`
	// The ID of the connector.
	//
	// Possible values include 'google-alloydb', 'google-cloudsql-mysql',
	// 'google-cloudsql-postgres', and other connector IDs supported by the BigQuery Connector framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#connector_id BigqueryConnection#connector_id}
	ConnectorId *string `field:"required" json:"connectorId" yaml:"connectorId"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#authentication BigqueryConnection#authentication}
	Authentication *BigqueryConnectionConfigurationAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#endpoint BigqueryConnection#endpoint}
	Endpoint *BigqueryConnectionConfigurationEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#network BigqueryConnection#network}
	Network *BigqueryConnectionConfigurationNetwork `field:"optional" json:"network" yaml:"network"`
}

