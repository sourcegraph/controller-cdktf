package oracledatabasedbsystem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseDbSystemConfig struct {
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
	// The ID of the DbSystem to create.
	//
	// This value is
	// restricted to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of
	// 63 characters in length. The value must start with a letter and end with a
	// letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_system_id OracleDatabaseDbSystem#db_system_id}
	DbSystemId *string `field:"required" json:"dbSystemId" yaml:"dbSystemId"`
	// The display name for the System db. The name does not have to be unique within your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#display_name OracleDatabaseDbSystem#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#location OracleDatabaseDbSystem#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the OdbSubnet associated with the DbSystem for IP allocation. Format: projects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#odb_subnet OracleDatabaseDbSystem#odb_subnet}
	OdbSubnet *string `field:"required" json:"odbSubnet" yaml:"odbSubnet"`
	// Whether or not to allow Terraform to destroy the instance.
	//
	// Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#deletion_protection OracleDatabaseDbSystem#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The GCP Oracle zone where Oracle DbSystem is hosted.
	//
	// Example: us-east4-b-r2.
	// If not specified, the system will pick a zone based on availability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#gcp_oracle_zone OracleDatabaseDbSystem#gcp_oracle_zone}
	GcpOracleZone *string `field:"optional" json:"gcpOracleZone" yaml:"gcpOracleZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#id OracleDatabaseDbSystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels or tags associated with the DbSystem.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#labels OracleDatabaseDbSystem#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the OdbNetwork associated with the DbSystem.
	//
	// Format: projects/{project}/locations/{location}/odbNetworks/{odb_network}
	// It is optional but if specified, this should match the parent ODBNetwork of
	// the OdbSubnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#odb_network OracleDatabaseDbSystem#odb_network}
	OdbNetwork *string `field:"optional" json:"odbNetwork" yaml:"odbNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#project OracleDatabaseDbSystem#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#properties OracleDatabaseDbSystem#properties}
	Properties *OracleDatabaseDbSystemProperties `field:"optional" json:"properties" yaml:"properties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#timeouts OracleDatabaseDbSystem#timeouts}
	Timeouts *OracleDatabaseDbSystemTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

