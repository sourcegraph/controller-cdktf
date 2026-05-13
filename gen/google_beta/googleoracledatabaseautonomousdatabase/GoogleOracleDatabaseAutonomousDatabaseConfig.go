package googleoracledatabaseautonomousdatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleOracleDatabaseAutonomousDatabaseConfig struct {
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
	// The ID of the Autonomous Database to create.
	//
	// This value is restricted
	// to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63
	// characters in length. The value must start with a letter and end with
	// a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#autonomous_database_id GoogleOracleDatabaseAutonomousDatabase#autonomous_database_id}
	AutonomousDatabaseId *string `field:"required" json:"autonomousDatabaseId" yaml:"autonomousDatabaseId"`
	// Resource ID segment making up resource 'name'. See documentation for resource type 'oracledatabase.googleapis.com/AutonomousDatabaseBackup'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#location GoogleOracleDatabaseAutonomousDatabase#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The password for the default ADMIN user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#admin_password GoogleOracleDatabaseAutonomousDatabase#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// The subnet CIDR range for the Autonmous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#cidr GoogleOracleDatabaseAutonomousDatabase#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The name of the Autonomous Database.
	//
	// The database name must be unique in
	// the project. The name must begin with a letter and can
	// contain a maximum of 30 alphanumeric characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#database GoogleOracleDatabaseAutonomousDatabase#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Whether or not to allow Terraform to destroy the instance.
	//
	// Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#deletion_protection GoogleOracleDatabaseAutonomousDatabase#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The display name for the Autonomous Database. The name does not have to be unique within your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#display_name GoogleOracleDatabaseAutonomousDatabase#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#id GoogleOracleDatabaseAutonomousDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels or tags associated with the Autonomous Database.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#labels GoogleOracleDatabaseAutonomousDatabase#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the VPC network used by the Autonomous Database. Format: projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#network GoogleOracleDatabaseAutonomousDatabase#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The name of the OdbNetwork associated with the Autonomous Database.
	//
	// Format:
	// projects/{project}/locations/{location}/odbNetworks/{odb_network}
	// It is optional but if specified, this should match the parent ODBNetwork of
	// the odb_subnet and backup_odb_subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#odb_network GoogleOracleDatabaseAutonomousDatabase#odb_network}
	OdbNetwork *string `field:"optional" json:"odbNetwork" yaml:"odbNetwork"`
	// The name of the OdbSubnet associated with the Autonomous Database for IP allocation. Format: projects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#odb_subnet GoogleOracleDatabaseAutonomousDatabase#odb_subnet}
	OdbSubnet *string `field:"optional" json:"odbSubnet" yaml:"odbSubnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#project GoogleOracleDatabaseAutonomousDatabase#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#properties GoogleOracleDatabaseAutonomousDatabase#properties}
	Properties *GoogleOracleDatabaseAutonomousDatabaseProperties `field:"optional" json:"properties" yaml:"properties"`
	// source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#source_config GoogleOracleDatabaseAutonomousDatabase#source_config}
	SourceConfig *GoogleOracleDatabaseAutonomousDatabaseSourceConfig `field:"optional" json:"sourceConfig" yaml:"sourceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_autonomous_database#timeouts GoogleOracleDatabaseAutonomousDatabase#timeouts}
	Timeouts *GoogleOracleDatabaseAutonomousDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

