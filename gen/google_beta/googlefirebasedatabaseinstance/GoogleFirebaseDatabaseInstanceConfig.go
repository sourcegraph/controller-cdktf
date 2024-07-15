package googlefirebasedatabaseinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseDatabaseInstanceConfig struct {
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
	// The globally unique identifier of the Firebase Realtime Database instance. Instance IDs cannot be reused after deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_database_instance#instance_id GoogleFirebaseDatabaseInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// A reference to the region where the Firebase Realtime database resides. Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_database_instance#region GoogleFirebaseDatabaseInstance#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The intended database state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_database_instance#desired_state GoogleFirebaseDatabaseInstance#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_database_instance#id GoogleFirebaseDatabaseInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_database_instance#project GoogleFirebaseDatabaseInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_database_instance#timeouts GoogleFirebaseDatabaseInstance#timeouts}
	Timeouts *GoogleFirebaseDatabaseInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The database type.
	//
	// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan.
	// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value: "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_database_instance#type GoogleFirebaseDatabaseInstance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

