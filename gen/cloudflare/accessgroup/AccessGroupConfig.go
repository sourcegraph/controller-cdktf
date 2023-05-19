package accessgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessGroupConfig struct {
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
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#include AccessGroup#include}
	Include interface{} `field:"required" json:"include" yaml:"include"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#name AccessGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account identifier to target for the resource.
	//
	// Conflicts with `zone_id`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#account_id AccessGroup#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#exclude AccessGroup#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#id AccessGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// require block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#require AccessGroup#require}
	Require interface{} `field:"optional" json:"require" yaml:"require"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#zone_id AccessGroup#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

