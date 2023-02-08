package captcha

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CaptchaConfig struct {
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
	// Name of the CAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha#name Captcha#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Secret key issued from the CAPTCHA vendor to perform server-side validation for a CAPTCHA token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha#secret_key Captcha#secret_key}
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// Site key issued from the CAPTCHA vendor to render a CAPTCHA on a page.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha#site_key Captcha#site_key}
	SiteKey *string `field:"required" json:"siteKey" yaml:"siteKey"`
	// Captcha type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha#type Captcha#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/captcha#id Captcha#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

