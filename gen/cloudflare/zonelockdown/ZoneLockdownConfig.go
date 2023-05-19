package zonelockdown

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneLockdownConfig struct {
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
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#configurations ZoneLockdown#configurations}
	Configurations interface{} `field:"required" json:"configurations" yaml:"configurations"`
	// A list of simple wildcard patterns to match requests against. The order of the urls is unimportant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#urls ZoneLockdown#urls}
	Urls *[]*string `field:"required" json:"urls" yaml:"urls"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#zone_id ZoneLockdown#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// A description about the lockdown entry. Typically used as a reminder or explanation for the lockdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#description ZoneLockdown#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#id ZoneLockdown#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Boolean of whether this zone lockdown is currently paused. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#paused ZoneLockdown#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone_lockdown#priority ZoneLockdown#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

