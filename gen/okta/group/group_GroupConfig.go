package group

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Group name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group#name Group#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// JSON formatted custom attributes for a group. It must be JSON due to various types Okta allows.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group#custom_profile_attributes Group#custom_profile_attributes}
	CustomProfileAttributes *string `field:"optional" json:"customProfileAttributes" yaml:"customProfileAttributes"`
	// Group description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group#description Group#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group#id Group#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group#skip_users Group#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// Users associated with the group. This can also be done per user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group#users Group#users}
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

