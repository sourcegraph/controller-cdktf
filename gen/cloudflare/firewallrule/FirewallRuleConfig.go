package firewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallRuleConfig struct {
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
	// The action to apply to a matched request. Available values: `block`, `challenge`, `allow`, `js_challenge`, `managed_challenge`, `log`, `bypass`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#action FirewallRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The identifier of the Filter to use for determining if the Firewall Rule should be triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#filter_id FirewallRule#filter_id}
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#zone_id FirewallRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// A description of the rule to help identify it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#description FirewallRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#id FirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether this filter based firewall rule is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#paused FirewallRule#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// The priority of the rule to allow control of processing order.
	//
	// A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#priority FirewallRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// List of products to bypass for a request when the bypass action is used.
	//
	// Available values: `zoneLockdown`, `uaBlock`, `bic`, `hot`, `securityLevel`, `rateLimit`, `waf`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/firewall_rule#products FirewallRule#products}
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
}

