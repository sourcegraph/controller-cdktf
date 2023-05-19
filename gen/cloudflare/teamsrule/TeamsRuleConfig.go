package teamsrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamsRuleConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#account_id TeamsRule#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The action executed by matched teams rule.
	//
	// Available values: `allow`, `block`, `safesearch`, `ytrestricted`, `on`, `off`, `scan`, `noscan`, `isolate`, `noisolate`, `override`, `l4_override`, `egress`, `audit_ssh`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#action TeamsRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The description of the teams rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#description TeamsRule#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the teams rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#name TeamsRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The evaluation precedence of the teams rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#precedence TeamsRule#precedence}
	Precedence *float64 `field:"required" json:"precedence" yaml:"precedence"`
	// The wirefilter expression to be used for device_posture check matching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#device_posture TeamsRule#device_posture}
	DevicePosture *string `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Indicator of rule enablement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#enabled TeamsRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The protocol or layer to evaluate the traffic and identity expressions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#filters TeamsRule#filters}
	Filters *[]*string `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#id TeamsRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The wirefilter expression to be used for identity matching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#identity TeamsRule#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// rule_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#rule_settings TeamsRule#rule_settings}
	RuleSettings *TeamsRuleRuleSettings `field:"optional" json:"ruleSettings" yaml:"ruleSettings"`
	// The wirefilter expression to be used for traffic matching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#traffic TeamsRule#traffic}
	Traffic *string `field:"optional" json:"traffic" yaml:"traffic"`
}

