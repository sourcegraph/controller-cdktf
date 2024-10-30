package dataoktauser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaUserConfig struct {
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
	// Search operator used when joining multiple search clauses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/user#compound_search_operator DataOktaUser#compound_search_operator}
	CompoundSearchOperator *string `field:"optional" json:"compoundSearchOperator" yaml:"compoundSearchOperator"`
	// Force delay of the user read by N seconds.
	//
	// Useful when eventual consistency of user information needs to be allowed for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/user#delay_read_seconds DataOktaUser#delay_read_seconds}
	DelayReadSeconds *string `field:"optional" json:"delayReadSeconds" yaml:"delayReadSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/user#id DataOktaUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/user#search DataOktaUser#search}
	Search interface{} `field:"optional" json:"search" yaml:"search"`
	// Do not populate user groups information (prevents additional API call).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/user#skip_groups DataOktaUser#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Do not populate user roles information (prevents additional API call).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/user#skip_roles DataOktaUser#skip_roles}
	SkipRoles interface{} `field:"optional" json:"skipRoles" yaml:"skipRoles"`
	// Retrieve a single user based on their id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/user#user_id DataOktaUser#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

