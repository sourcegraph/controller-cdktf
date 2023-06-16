package datacloudflareaccessidentityprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAccessIdentityProviderConfig struct {
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
	// Access Identity Provider name to search for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/access_identity_provider#name DataCloudflareAccessIdentityProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account identifier to target for the resource. Must provide only one of `zone_id`, `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/access_identity_provider#account_id DataCloudflareAccessIdentityProvider#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/access_identity_provider#id DataCloudflareAccessIdentityProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The zone identifier to target for the resource. Must provide only one of `zone_id`, `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/data-sources/access_identity_provider#zone_id DataCloudflareAccessIdentityProvider#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

