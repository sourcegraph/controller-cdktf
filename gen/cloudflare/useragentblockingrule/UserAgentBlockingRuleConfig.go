package useragentblockingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserAgentBlockingRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#configuration UserAgentBlockingRule#configuration}
	Configuration *UserAgentBlockingRuleConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// An informative summary of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#description UserAgentBlockingRule#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The action to apply to a matched request. Available values: `block`, `challenge`, `js_challenge`, `managed_challenge`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#mode UserAgentBlockingRule#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// When true, indicates that the rule is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#paused UserAgentBlockingRule#paused}
	Paused interface{} `field:"required" json:"paused" yaml:"paused"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#zone_id UserAgentBlockingRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/user_agent_blocking_rule#id UserAgentBlockingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

