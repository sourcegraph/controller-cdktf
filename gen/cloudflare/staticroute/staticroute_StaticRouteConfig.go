package staticroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StaticRouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#nexthop StaticRoute#nexthop}.
	Nexthop *string `field:"required" json:"nexthop" yaml:"nexthop"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#prefix StaticRoute#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#priority StaticRoute#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#account_id StaticRoute#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#colo_names StaticRoute#colo_names}.
	ColoNames *[]*string `field:"optional" json:"coloNames" yaml:"coloNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#colo_regions StaticRoute#colo_regions}.
	ColoRegions *[]*string `field:"optional" json:"coloRegions" yaml:"coloRegions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#description StaticRoute#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#id StaticRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/static_route#weight StaticRoute#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

