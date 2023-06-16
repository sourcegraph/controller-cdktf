package googlegameservicesrealm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGameServicesRealmConfig struct {
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
	// GCP region of the Realm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#realm_id GoogleGameServicesRealm#realm_id}
	RealmId *string `field:"required" json:"realmId" yaml:"realmId"`
	// Required.
	//
	// Time zone where all realm-specific policies are evaluated. The value of
	// this field must be from the IANA time zone database:
	// https://www.iana.org/time-zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#time_zone GoogleGameServicesRealm#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Human readable description of the realm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#description GoogleGameServicesRealm#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#id GoogleGameServicesRealm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels associated with this realm. Each label is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#labels GoogleGameServicesRealm#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Location of the Realm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#location GoogleGameServicesRealm#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#project GoogleGameServicesRealm#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_game_services_realm#timeouts GoogleGameServicesRealm#timeouts}
	Timeouts *GoogleGameServicesRealmTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

