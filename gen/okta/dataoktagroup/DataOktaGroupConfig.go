package dataoktagroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaGroupConfig struct {
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
	// Force delay of the group read by N seconds.
	//
	// Useful when eventual consistency of group information needs to be allowed for; for instance, when group rules are known to have been applied.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/group#delay_read_seconds DataOktaGroup#delay_read_seconds}
	DelayReadSeconds *string `field:"optional" json:"delayReadSeconds" yaml:"delayReadSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/group#id DataOktaGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Fetch group users, having default off cuts down on API calls.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/group#include_users DataOktaGroup#include_users}
	IncludeUsers interface{} `field:"optional" json:"includeUsers" yaml:"includeUsers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/group#name DataOktaGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type of the group.
	//
	// When specified in the terraform resource, will act as a filter when searching for the group
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/group#type DataOktaGroup#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

