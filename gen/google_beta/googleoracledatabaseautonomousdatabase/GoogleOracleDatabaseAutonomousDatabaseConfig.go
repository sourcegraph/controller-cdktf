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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#autonomous_database_id GoogleOracleDatabaseAutonomousDatabase#autonomous_database_id}
	AutonomousDatabaseId *string `field:"required" json:"autonomousDatabaseId" yaml:"autonomousDatabaseId"`
	// The subnet CIDR range for the Autonmous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#cidr GoogleOracleDatabaseAutonomousDatabase#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The name of the Autonomous Database.
	//
	// The database name must be unique in
	// the project. The name must begin with a letter and can
	// contain a maximum of 30 alphanumeric characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#database GoogleOracleDatabaseAutonomousDatabase#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Resource ID segment making up resource 'name'. See documentation for resource type 'oracledatabase.googleapis.com/AutonomousDatabaseBackup'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#location GoogleOracleDatabaseAutonomousDatabase#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the VPC network used by the Autonomous Database. Format: projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#network GoogleOracleDatabaseAutonomousDatabase#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#properties GoogleOracleDatabaseAutonomousDatabase#properties}
	Properties *GoogleOracleDatabaseAutonomousDatabaseProperties `field:"required" json:"properties" yaml:"properties"`
	// The password for the default ADMIN user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#admin_password GoogleOracleDatabaseAutonomousDatabase#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Whether or not to allow Terraform to destroy the instance.
	//
	// Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#deletion_protection GoogleOracleDatabaseAutonomousDatabase#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The display name for the Autonomous Database. The name does not have to be unique within your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#display_name GoogleOracleDatabaseAutonomousDatabase#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#id GoogleOracleDatabaseAutonomousDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels or tags associated with the Autonomous Database.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#labels GoogleOracleDatabaseAutonomousDatabase#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#project GoogleOracleDatabaseAutonomousDatabase#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_oracle_database_autonomous_database#timeouts GoogleOracleDatabaseAutonomousDatabase#timeouts}
	Timeouts *GoogleOracleDatabaseAutonomousDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

