package zonelockdown

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneLockdownConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#configurations ZoneLockdown#configurations}
	Configurations interface{} `field:"required" json:"configurations" yaml:"configurations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#urls ZoneLockdown#urls}.
	Urls *[]*string `field:"required" json:"urls" yaml:"urls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#zone_id ZoneLockdown#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#description ZoneLockdown#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#id ZoneLockdown#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#paused ZoneLockdown#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/zone_lockdown#priority ZoneLockdown#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

