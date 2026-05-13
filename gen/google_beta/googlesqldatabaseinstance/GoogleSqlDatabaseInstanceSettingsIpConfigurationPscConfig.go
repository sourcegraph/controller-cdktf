package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfig struct {
	// List of consumer projects that are allow-listed for PSC connections to this instance.
	//
	// This instance can be connected to with PSC from any network in these projects. Each consumer project in this list may be represented by a project number (numeric) or by a project id (alphanumeric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#allowed_consumer_projects GoogleSqlDatabaseInstance#allowed_consumer_projects}
	AllowedConsumerProjects *[]*string `field:"optional" json:"allowedConsumerProjects" yaml:"allowedConsumerProjects"`
	// Name of network attachment resource used to authorize a producer service to connect a PSC interface to the consumer's VPC.
	//
	// For example: "projects/myProject/regions/myRegion/networkAttachments/myNetworkAttachment". This is required to enable outbound connection on a PSC instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#network_attachment_uri GoogleSqlDatabaseInstance#network_attachment_uri}
	NetworkAttachmentUri *string `field:"optional" json:"networkAttachmentUri" yaml:"networkAttachmentUri"`
	// psc_auto_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#psc_auto_connections GoogleSqlDatabaseInstance#psc_auto_connections}
	PscAutoConnections interface{} `field:"optional" json:"pscAutoConnections" yaml:"pscAutoConnections"`
	// Whether PSC connectivity is enabled for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#psc_enabled GoogleSqlDatabaseInstance#psc_enabled}
	PscEnabled interface{} `field:"optional" json:"pscEnabled" yaml:"pscEnabled"`
}

