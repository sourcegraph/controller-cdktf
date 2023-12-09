package accesscustompage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessCustomPageConfig struct {
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
	// Friendly name of the Access Custom Page configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_custom_page#name AccessCustomPage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of Access custom page to create. Available values: `identity_denied`, `forbidden`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_custom_page#type AccessCustomPage#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The account identifier to target for the resource.
	//
	// Conflicts with `zone_id`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_custom_page#account_id AccessCustomPage#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Number of apps to display on the custom page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_custom_page#app_count AccessCustomPage#app_count}
	AppCount *float64 `field:"optional" json:"appCount" yaml:"appCount"`
	// Custom HTML to display on the custom page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_custom_page#custom_html AccessCustomPage#custom_html}
	CustomHtml *string `field:"optional" json:"customHtml" yaml:"customHtml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_custom_page#id AccessCustomPage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The zone identifier to target for the resource.
	//
	// Conflicts with `account_id`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_custom_page#zone_id AccessCustomPage#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

