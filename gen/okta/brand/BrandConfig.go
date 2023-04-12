package brand

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BrandConfig struct {
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
	// Consent for updating the custom privacy policy URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/brand#agree_to_custom_privacy_policy Brand#agree_to_custom_privacy_policy}
	AgreeToCustomPrivacyPolicy interface{} `field:"optional" json:"agreeToCustomPrivacyPolicy" yaml:"agreeToCustomPrivacyPolicy"`
	// Brand ID - Note: Okta API for brands only reads and updates therefore the okta_brand resource needs to act as a quasi data source.
	//
	// Do this by setting brand_id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/brand#brand_id Brand#brand_id}
	BrandId *string `field:"optional" json:"brandId" yaml:"brandId"`
	// Custom privacy policy URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/brand#custom_privacy_policy_url Brand#custom_privacy_policy_url}
	CustomPrivacyPolicyUrl *string `field:"optional" json:"customPrivacyPolicyUrl" yaml:"customPrivacyPolicyUrl"`
	// Removes "Powered by Okta" from the Okta-hosted sign-in page and "© 2021 Okta, Inc." from the Okta End-User Dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/brand#remove_powered_by_okta Brand#remove_powered_by_okta}
	RemovePoweredByOkta interface{} `field:"optional" json:"removePoweredByOkta" yaml:"removePoweredByOkta"`
}

