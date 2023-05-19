package googlefirestoredatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirestoreDatabaseConfig struct {
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
	// The location of the database. Available databases are listed at https://cloud.google.com/firestore/docs/locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#location_id GoogleFirestoreDatabase#location_id}
	LocationId *string `field:"required" json:"locationId" yaml:"locationId"`
	// The ID to use for the database, which will become the final component of the database's resource name.
	//
	// This value should be 4-63
	// characters. Valid characters are /[a-z][0-9]-/ with first character
	// a letter and the last a letter or a number. Must not be
	// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	// "(default)" database id is also valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#name GoogleFirestoreDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the database. See https://cloud.google.com/datastore/docs/firestore-or-datastore for information about how to choose. Possible values: ["FIRESTORE_NATIVE", "DATASTORE_MODE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#type GoogleFirestoreDatabase#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The App Engine integration mode to use for this database. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#app_engine_integration_mode GoogleFirestoreDatabase#app_engine_integration_mode}
	AppEngineIntegrationMode *string `field:"optional" json:"appEngineIntegrationMode" yaml:"appEngineIntegrationMode"`
	// The concurrency control mode to use for this database. Possible values: ["OPTIMISTIC", "PESSIMISTIC", "OPTIMISTIC_WITH_ENTITY_GROUPS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#concurrency_mode GoogleFirestoreDatabase#concurrency_mode}
	ConcurrencyMode *string `field:"optional" json:"concurrencyMode" yaml:"concurrencyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#id GoogleFirestoreDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#project GoogleFirestoreDatabase#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_database#timeouts GoogleFirestoreDatabase#timeouts}
	Timeouts *GoogleFirestoreDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

