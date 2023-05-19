package accessidentityprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessIdentityProviderConfig struct {
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
	// Friendly name of the Access Identity Provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#name AccessIdentityProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The provider type to use.
	//
	// Available values: `centrify`, `facebook`, `google-apps`, `oidc`, `github`, `google`, `saml`, `linkedin`, `azureAD`, `okta`, `onetimepin`, `onelogin`, `yandex`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#type AccessIdentityProvider#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The account identifier to target for the resource.
	//
	// Conflicts with `zone_id`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#account_id AccessIdentityProvider#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#config AccessIdentityProvider#config}
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#id AccessIdentityProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// scim_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#scim_config AccessIdentityProvider#scim_config}
	ScimConfig interface{} `field:"optional" json:"scimConfig" yaml:"scimConfig"`
	// The zone identifier to target for the resource.
	//
	// Conflicts with `account_id`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#zone_id AccessIdentityProvider#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

