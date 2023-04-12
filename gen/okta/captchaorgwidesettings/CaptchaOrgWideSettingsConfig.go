package captchaorgwidesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CaptchaOrgWideSettingsConfig struct {
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
	// ID of the CAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha_org_wide_settings#captcha_id CaptchaOrgWideSettings#captcha_id}
	CaptchaId *string `field:"optional" json:"captchaId" yaml:"captchaId"`
	// Set of pages that have CAPTCHA enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha_org_wide_settings#enabled_for CaptchaOrgWideSettings#enabled_for}
	EnabledFor *[]*string `field:"optional" json:"enabledFor" yaml:"enabledFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha_org_wide_settings#id CaptchaOrgWideSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

