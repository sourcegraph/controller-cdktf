package accountmember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountMemberConfig struct {
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
	// Account ID to create the account member in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/account_member#account_id AccountMember#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The email address of the user who you wish to manage.
	//
	// Following creation, this field becomes read only via the API and cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/account_member#email_address AccountMember#email_address}
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// List of account role IDs that you want to assign to a member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/account_member#role_ids AccountMember#role_ids}
	RoleIds *[]*string `field:"required" json:"roleIds" yaml:"roleIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/account_member#id AccountMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A member's status in the account. Available values: `accepted`, `pending`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/account_member#status AccountMember#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

