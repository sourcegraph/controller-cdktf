package googledatabasemigrationserviceprivateconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDatabaseMigrationServicePrivateConnectionConfig struct {
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
	// The name of the location this private connection is located in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#location GoogleDatabaseMigrationServicePrivateConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The private connectivity identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#private_connection_id GoogleDatabaseMigrationServicePrivateConnection#private_connection_id}
	PrivateConnectionId *string `field:"required" json:"privateConnectionId" yaml:"privateConnectionId"`
	// vpc_peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#vpc_peering_config GoogleDatabaseMigrationServicePrivateConnection#vpc_peering_config}
	VpcPeeringConfig *GoogleDatabaseMigrationServicePrivateConnectionVpcPeeringConfig `field:"required" json:"vpcPeeringConfig" yaml:"vpcPeeringConfig"`
	// Display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#display_name GoogleDatabaseMigrationServicePrivateConnection#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#id GoogleDatabaseMigrationServicePrivateConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#labels GoogleDatabaseMigrationServicePrivateConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#project GoogleDatabaseMigrationServicePrivateConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_private_connection#timeouts GoogleDatabaseMigrationServicePrivateConnection#timeouts}
	Timeouts *GoogleDatabaseMigrationServicePrivateConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

