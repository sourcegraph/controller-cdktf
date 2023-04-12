package record

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RecordConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#name Record#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#type Record#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#zone_id Record#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#allow_overwrite Record#allow_overwrite}.
	AllowOverwrite interface{} `field:"optional" json:"allowOverwrite" yaml:"allowOverwrite"`
	// data block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#data Record#data}
	Data *RecordData `field:"optional" json:"data" yaml:"data"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#id Record#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#priority Record#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#proxied Record#proxied}.
	Proxied interface{} `field:"optional" json:"proxied" yaml:"proxied"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#timeouts Record#timeouts}
	Timeouts *RecordTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#ttl Record#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#value Record#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

