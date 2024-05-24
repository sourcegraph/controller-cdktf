package zonehold

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneHoldConfig struct {
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
	// Enablement status of the zone hold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/zone_hold#hold ZoneHold#hold}
	Hold interface{} `field:"required" json:"hold" yaml:"hold"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/zone_hold#zone_id ZoneHold#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The RFC3339 compatible timestamp when to automatically re-enable the zone hold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/zone_hold#hold_after ZoneHold#hold_after}
	HoldAfter *string `field:"optional" json:"holdAfter" yaml:"holdAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/zone_hold#id ZoneHold#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to extend to block any subdomain of the given zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/zone_hold#include_subdomains ZoneHold#include_subdomains}
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
}

