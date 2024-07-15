package googledatastreamconnectionprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDatastreamConnectionProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The connection profile identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#connection_profile_id GoogleDatastreamConnectionProfile#connection_profile_id}
	ConnectionProfileId *string `field:"required" json:"connectionProfileId" yaml:"connectionProfileId"`
	// Display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#display_name GoogleDatastreamConnectionProfile#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the location this connection profile is located in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#location GoogleDatastreamConnectionProfile#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// bigquery_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#bigquery_profile GoogleDatastreamConnectionProfile#bigquery_profile}
	BigqueryProfile *GoogleDatastreamConnectionProfileBigqueryProfile `field:"optional" json:"bigqueryProfile" yaml:"bigqueryProfile"`
	// Create the connection profile without validating it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#create_without_validation GoogleDatastreamConnectionProfile#create_without_validation}
	CreateWithoutValidation interface{} `field:"optional" json:"createWithoutValidation" yaml:"createWithoutValidation"`
	// forward_ssh_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#forward_ssh_connectivity GoogleDatastreamConnectionProfile#forward_ssh_connectivity}
	ForwardSshConnectivity *GoogleDatastreamConnectionProfileForwardSshConnectivity `field:"optional" json:"forwardSshConnectivity" yaml:"forwardSshConnectivity"`
	// gcs_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#gcs_profile GoogleDatastreamConnectionProfile#gcs_profile}
	GcsProfile *GoogleDatastreamConnectionProfileGcsProfile `field:"optional" json:"gcsProfile" yaml:"gcsProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#id GoogleDatastreamConnectionProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#labels GoogleDatastreamConnectionProfile#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// mysql_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#mysql_profile GoogleDatastreamConnectionProfile#mysql_profile}
	MysqlProfile *GoogleDatastreamConnectionProfileMysqlProfile `field:"optional" json:"mysqlProfile" yaml:"mysqlProfile"`
	// oracle_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#oracle_profile GoogleDatastreamConnectionProfile#oracle_profile}
	OracleProfile *GoogleDatastreamConnectionProfileOracleProfile `field:"optional" json:"oracleProfile" yaml:"oracleProfile"`
	// postgresql_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#postgresql_profile GoogleDatastreamConnectionProfile#postgresql_profile}
	PostgresqlProfile *GoogleDatastreamConnectionProfilePostgresqlProfile `field:"optional" json:"postgresqlProfile" yaml:"postgresqlProfile"`
	// private_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#private_connectivity GoogleDatastreamConnectionProfile#private_connectivity}
	PrivateConnectivity *GoogleDatastreamConnectionProfilePrivateConnectivity `field:"optional" json:"privateConnectivity" yaml:"privateConnectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#project GoogleDatastreamConnectionProfile#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// sql_server_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#sql_server_profile GoogleDatastreamConnectionProfile#sql_server_profile}
	SqlServerProfile *GoogleDatastreamConnectionProfileSqlServerProfile `field:"optional" json:"sqlServerProfile" yaml:"sqlServerProfile"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#timeouts GoogleDatastreamConnectionProfile#timeouts}
	Timeouts *GoogleDatastreamConnectionProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

