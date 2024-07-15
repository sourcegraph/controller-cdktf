package googledatabasemigrationserviceconnectionprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDatabaseMigrationServiceConnectionProfileConfig struct {
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
	// The ID of the connection profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#connection_profile_id GoogleDatabaseMigrationServiceConnectionProfile#connection_profile_id}
	ConnectionProfileId *string `field:"required" json:"connectionProfileId" yaml:"connectionProfileId"`
	// alloydb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#alloydb GoogleDatabaseMigrationServiceConnectionProfile#alloydb}
	Alloydb *GoogleDatabaseMigrationServiceConnectionProfileAlloydb `field:"optional" json:"alloydb" yaml:"alloydb"`
	// cloudsql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#cloudsql GoogleDatabaseMigrationServiceConnectionProfile#cloudsql}
	Cloudsql *GoogleDatabaseMigrationServiceConnectionProfileCloudsql `field:"optional" json:"cloudsql" yaml:"cloudsql"`
	// The connection profile display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#display_name GoogleDatabaseMigrationServiceConnectionProfile#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#id GoogleDatabaseMigrationServiceConnectionProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#labels GoogleDatabaseMigrationServiceConnectionProfile#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where the connection profile should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#location GoogleDatabaseMigrationServiceConnectionProfile#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#mysql GoogleDatabaseMigrationServiceConnectionProfile#mysql}
	Mysql *GoogleDatabaseMigrationServiceConnectionProfileMysql `field:"optional" json:"mysql" yaml:"mysql"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#oracle GoogleDatabaseMigrationServiceConnectionProfile#oracle}
	Oracle *GoogleDatabaseMigrationServiceConnectionProfileOracle `field:"optional" json:"oracle" yaml:"oracle"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#postgresql GoogleDatabaseMigrationServiceConnectionProfile#postgresql}
	Postgresql *GoogleDatabaseMigrationServiceConnectionProfilePostgresql `field:"optional" json:"postgresql" yaml:"postgresql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#project GoogleDatabaseMigrationServiceConnectionProfile#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile#timeouts GoogleDatabaseMigrationServiceConnectionProfile#timeouts}
	Timeouts *GoogleDatabaseMigrationServiceConnectionProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

