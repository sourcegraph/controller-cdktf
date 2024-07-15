package googlemigrationcenterpreferenceset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleMigrationCenterPreferenceSetConfig struct {
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
	// Part of 'parent'. See documentation of 'projectsId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#location GoogleMigrationCenterPreferenceSet#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required.
	//
	// User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression '[a-z]([a-z0-9-]{0,61}[a-z0-9])?'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#preference_set_id GoogleMigrationCenterPreferenceSet#preference_set_id}
	PreferenceSetId *string `field:"required" json:"preferenceSetId" yaml:"preferenceSetId"`
	// A description of the preference set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#description GoogleMigrationCenterPreferenceSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name. Maximum length is 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#display_name GoogleMigrationCenterPreferenceSet#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#id GoogleMigrationCenterPreferenceSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#project GoogleMigrationCenterPreferenceSet#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#timeouts GoogleMigrationCenterPreferenceSet#timeouts}
	Timeouts *GoogleMigrationCenterPreferenceSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_machine_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_migration_center_preference_set#virtual_machine_preferences GoogleMigrationCenterPreferenceSet#virtual_machine_preferences}
	VirtualMachinePreferences *GoogleMigrationCenterPreferenceSetVirtualMachinePreferences `field:"optional" json:"virtualMachinePreferences" yaml:"virtualMachinePreferences"`
}

