package brand

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BrandConfig struct {
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
	// Name of the brand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#name Brand#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Is a required input flag with when changing custom_privacy_url, shouldn't be considered as a readable property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#agree_to_custom_privacy_policy Brand#agree_to_custom_privacy_policy}
	AgreeToCustomPrivacyPolicy interface{} `field:"optional" json:"agreeToCustomPrivacyPolicy" yaml:"agreeToCustomPrivacyPolicy"`
	// Brand ID - Note: Okta API for brands only reads and updates therefore the okta_brand resource needs to act as a quasi data source.
	//
	// Do this by setting brand_id. `DEPRECATED`: Okta has fully support brand creation, this attribute is a no op and will be removed
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#brand_id Brand#brand_id}
	BrandId *string `field:"optional" json:"brandId" yaml:"brandId"`
	// Custom privacy policy URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#custom_privacy_policy_url Brand#custom_privacy_policy_url}
	CustomPrivacyPolicyUrl *string `field:"optional" json:"customPrivacyPolicyUrl" yaml:"customPrivacyPolicyUrl"`
	// Default app app instance id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#default_app_app_instance_id Brand#default_app_app_instance_id}
	DefaultAppAppInstanceId *string `field:"optional" json:"defaultAppAppInstanceId" yaml:"defaultAppAppInstanceId"`
	// Default app app link name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#default_app_app_link_name Brand#default_app_app_link_name}
	DefaultAppAppLinkName *string `field:"optional" json:"defaultAppAppLinkName" yaml:"defaultAppAppLinkName"`
	// Default app classic application uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#default_app_classic_application_uri Brand#default_app_classic_application_uri}
	DefaultAppClassicApplicationUri *string `field:"optional" json:"defaultAppClassicApplicationUri" yaml:"defaultAppClassicApplicationUri"`
	// The language specified as an IETF BCP 47 language tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#locale Brand#locale}
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Removes "Powered by Okta" from the Okta-hosted sign-in page and "© 2021 Okta, Inc." from the Okta End-User Dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand#remove_powered_by_okta Brand#remove_powered_by_okta}
	RemovePoweredByOkta interface{} `field:"optional" json:"removePoweredByOkta" yaml:"removePoweredByOkta"`
}

