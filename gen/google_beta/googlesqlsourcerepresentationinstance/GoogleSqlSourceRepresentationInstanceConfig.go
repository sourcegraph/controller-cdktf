package googlesqlsourcerepresentationinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSqlSourceRepresentationInstanceConfig struct {
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
	// The MySQL version running on your source database server.
	//
	// Possible values: ["MYSQL_5_6", "MYSQL_5_7", "MYSQL_8_0", "POSTGRES_9_6", "POSTGRES_10", "POSTGRES_11", "POSTGRES_12", "POSTGRES_13", "POSTGRES_14"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#database_version GoogleSqlSourceRepresentationInstance#database_version}
	DatabaseVersion *string `field:"required" json:"databaseVersion" yaml:"databaseVersion"`
	// The IPv4 address and port for the external server, or the the DNS address for the external server.
	//
	// If the external server is hosted on Cloud SQL, the port is 5432.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#host GoogleSqlSourceRepresentationInstance#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#name GoogleSqlSourceRepresentationInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CA certificate on the external server. Include only if SSL/TLS is used on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#ca_certificate GoogleSqlSourceRepresentationInstance#ca_certificate}
	CaCertificate *string `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// The client certificate on the external server.
	//
	// Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#client_certificate GoogleSqlSourceRepresentationInstance#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// The private key file for the client certificate on the external server.
	//
	// Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#client_key GoogleSqlSourceRepresentationInstance#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// A file in the bucket that contains the data from the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#dump_file_path GoogleSqlSourceRepresentationInstance#dump_file_path}
	DumpFilePath *string `field:"optional" json:"dumpFilePath" yaml:"dumpFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#id GoogleSqlSourceRepresentationInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The password for the replication user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#password GoogleSqlSourceRepresentationInstance#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The externally accessible port for the source database server. Defaults to 3306.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#port GoogleSqlSourceRepresentationInstance#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#project GoogleSqlSourceRepresentationInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Region in which the created instance should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#region GoogleSqlSourceRepresentationInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#timeouts GoogleSqlSourceRepresentationInstance#timeouts}
	Timeouts *GoogleSqlSourceRepresentationInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The replication user account on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_source_representation_instance#username GoogleSqlSourceRepresentationInstance#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

