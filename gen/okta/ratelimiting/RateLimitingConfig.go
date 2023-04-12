package ratelimiting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RateLimitingConfig struct {
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
	// Called during authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/rate_limiting#authorize RateLimiting#authorize}
	Authorize *string `field:"required" json:"authorize" yaml:"authorize"`
	// Called when accessing the Okta hosted login page.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/rate_limiting#login RateLimiting#login}
	Login *string `field:"required" json:"login" yaml:"login"`
	// Enables rate limit warning, violation, notification emails and banners when this org meets rate limits.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/rate_limiting#communications_enabled RateLimiting#communications_enabled}
	CommunicationsEnabled interface{} `field:"optional" json:"communicationsEnabled" yaml:"communicationsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/rate_limiting#id RateLimiting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

