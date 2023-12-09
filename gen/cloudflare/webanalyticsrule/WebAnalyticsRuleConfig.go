package webanalyticsrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebAnalyticsRuleConfig struct {
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
	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#account_id WebAnalyticsRule#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The host to apply the rule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#host WebAnalyticsRule#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Whether the rule includes or excludes the matched traffic from being measured in Web Analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#inclusive WebAnalyticsRule#inclusive}
	Inclusive interface{} `field:"required" json:"inclusive" yaml:"inclusive"`
	// Whether the rule is paused or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#is_paused WebAnalyticsRule#is_paused}
	IsPaused interface{} `field:"required" json:"isPaused" yaml:"isPaused"`
	// A list of paths to apply the rule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#paths WebAnalyticsRule#paths}
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// The Web Analytics ruleset id. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#ruleset_id WebAnalyticsRule#ruleset_id}
	RulesetId *string `field:"required" json:"rulesetId" yaml:"rulesetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#id WebAnalyticsRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/web_analytics_rule#timeouts WebAnalyticsRule#timeouts}
	Timeouts *WebAnalyticsRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

