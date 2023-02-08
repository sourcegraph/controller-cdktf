package signonpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SignonPolicyConfig struct {
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
	// Policy Name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy#name SignonPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policy Description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy#description SignonPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of Group IDs to Include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy#groups_included SignonPolicy#groups_included}
	GroupsIncluded *[]*string `field:"optional" json:"groupsIncluded" yaml:"groupsIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy#id SignonPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Policy Priority, this attribute can be set to a valid priority.
	//
	// To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last (lowest) if not there.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy#priority SignonPolicy#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Policy Status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/signon_policy#status SignonPolicy#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

