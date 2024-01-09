package datacloudflareaccessapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAccessApplicationConfig struct {
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
	// The account identifier to target for the resource. Must provide only one of `zone_id`, `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/access_application#account_id DataCloudflareAccessApplication#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The primary hostname and path that Access will secure. Must provide only one of `name`, `domain`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/access_application#domain DataCloudflareAccessApplication#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/access_application#id DataCloudflareAccessApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Friendly name of the Access Application. Must provide only one of `name`, `domain`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/access_application#name DataCloudflareAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The zone identifier to target for the resource. Must provide only one of `zone_id`, `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/access_application#zone_id DataCloudflareAccessApplication#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

