package accessrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessRuleConfig struct {
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
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_rule#configuration AccessRule#configuration}
	Configuration *AccessRuleConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The action to apply to a matched request. Available values: `block`, `challenge`, `whitelist`, `js_challenge`, `managed_challenge`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_rule#mode AccessRule#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The account identifier to target for the resource.
	//
	// Must provide only one of `account_id`, `zone_id`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_rule#account_id AccessRule#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_rule#id AccessRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A personal note about the rule. Typically used as a reminder or explanation for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_rule#notes AccessRule#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The zone identifier to target for the resource.
	//
	// Must provide only one of `account_id`, `zone_id`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_rule#zone_id AccessRule#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

