package oracledatabaseexascaledbstoragevault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseExascaleDbStorageVaultConfig struct {
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
	// The display name for the ExascaleDbStorageVault.
	//
	// The name does not have to
	// be unique within your project. The name must be 1-255 characters long and
	// can only contain alphanumeric characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#display_name OracleDatabaseExascaleDbStorageVault#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID of the ExascaleDbStorageVault to create.
	//
	// This value is
	// restricted to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of
	// 63 characters in length. The value must start with a letter and end with a
	// letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#exascale_db_storage_vault_id OracleDatabaseExascaleDbStorageVault#exascale_db_storage_vault_id}
	ExascaleDbStorageVaultId *string `field:"required" json:"exascaleDbStorageVaultId" yaml:"exascaleDbStorageVaultId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#location OracleDatabaseExascaleDbStorageVault#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#properties OracleDatabaseExascaleDbStorageVault#properties}
	Properties *OracleDatabaseExascaleDbStorageVaultProperties `field:"required" json:"properties" yaml:"properties"`
	// Whether or not to allow Terraform to destroy the instance.
	//
	// Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#deletion_protection OracleDatabaseExascaleDbStorageVault#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The GCP Oracle zone where Oracle ExascaleDbStorageVault is hosted.
	//
	// Example: us-east4-b-r2.
	// If not specified, the system will pick a zone based on availability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#gcp_oracle_zone OracleDatabaseExascaleDbStorageVault#gcp_oracle_zone}
	GcpOracleZone *string `field:"optional" json:"gcpOracleZone" yaml:"gcpOracleZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#id OracleDatabaseExascaleDbStorageVault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels or tags associated with the ExascaleDbStorageVault.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#labels OracleDatabaseExascaleDbStorageVault#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#project OracleDatabaseExascaleDbStorageVault#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#timeouts OracleDatabaseExascaleDbStorageVault#timeouts}
	Timeouts *OracleDatabaseExascaleDbStorageVaultTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

