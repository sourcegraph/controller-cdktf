package grouprule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRuleConfig struct {
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
	// The expression value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#expression_value GroupRule#expression_value}
	ExpressionValue *string `field:"required" json:"expressionValue" yaml:"expressionValue"`
	// The list of group ids to assign the users to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#group_assignments GroupRule#group_assignments}
	GroupAssignments *[]*string `field:"required" json:"groupAssignments" yaml:"groupAssignments"`
	// The name of the Group Rule (min character 1; max characters 50).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#name GroupRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The expression type to use to invoke the rule. The default is `urn:okta:expression:1.0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#expression_type GroupRule#expression_type}
	ExpressionType *string `field:"optional" json:"expressionType" yaml:"expressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#id GroupRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Remove users added by this rule from the assigned group after deleting this resource. Default is `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#remove_assigned_users GroupRule#remove_assigned_users}
	RemoveAssignedUsers interface{} `field:"optional" json:"removeAssignedUsers" yaml:"removeAssignedUsers"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#status GroupRule#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The list of user IDs that would be excluded when rules are processed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_rule#users_excluded GroupRule#users_excluded}
	UsersExcluded *[]*string `field:"optional" json:"usersExcluded" yaml:"usersExcluded"`
}

