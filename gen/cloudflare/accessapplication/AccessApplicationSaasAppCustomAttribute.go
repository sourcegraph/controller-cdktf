package accessapplication


type AccessApplicationSaasAppCustomAttribute struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_application#source AccessApplication#source}
	Source *AccessApplicationSaasAppCustomAttributeSource `field:"required" json:"source" yaml:"source"`
	// A friendly name for the attribute as provided to the SaaS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_application#friendly_name AccessApplication#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// The name of the attribute as provided to the SaaS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_application#name AccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A globally unique name for an identity or service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_application#name_format AccessApplication#name_format}
	NameFormat *string `field:"optional" json:"nameFormat" yaml:"nameFormat"`
	// True if the attribute must be always present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_application#required AccessApplication#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

