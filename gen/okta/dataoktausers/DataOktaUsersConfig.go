package dataoktausers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaUsersConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Search operator used when joining multiple search clauses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#compound_search_operator DataOktaUsers#compound_search_operator}
	CompoundSearchOperator *string `field:"optional" json:"compoundSearchOperator" yaml:"compoundSearchOperator"`
	// Force delay of the users read by N seconds.
	//
	// Useful when eventual consistency of users information needs to be allowed for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#delay_read_seconds DataOktaUsers#delay_read_seconds}
	DelayReadSeconds *string `field:"optional" json:"delayReadSeconds" yaml:"delayReadSeconds"`
	// Find users based on group membership using the id of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#group_id DataOktaUsers#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#id DataOktaUsers#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Fetch group memberships for each user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#include_groups DataOktaUsers#include_groups}
	IncludeGroups any `field:"optional" json:"includeGroups" yaml:"includeGroups"`
	// Fetch user roles for each user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#include_roles DataOktaUsers#include_roles}
	IncludeRoles any `field:"optional" json:"includeRoles" yaml:"includeRoles"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/users#search DataOktaUsers#search}
	Search any `field:"optional" json:"search" yaml:"search"`
}
