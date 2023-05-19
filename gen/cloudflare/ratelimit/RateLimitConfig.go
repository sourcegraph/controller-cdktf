package ratelimit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RateLimitConfig struct {
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
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#action RateLimit#action}
	Action *RateLimitAction `field:"required" json:"action" yaml:"action"`
	// The time in seconds to count matching traffic.
	//
	// If the count exceeds threshold within this period the action will be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#period RateLimit#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The threshold that triggers the rate limit mitigations, combine with period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#threshold RateLimit#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#zone_id RateLimit#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#bypass_url_patterns RateLimit#bypass_url_patterns}.
	BypassUrlPatterns *[]*string `field:"optional" json:"bypassUrlPatterns" yaml:"bypassUrlPatterns"`
	// correlate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#correlate RateLimit#correlate}
	Correlate *RateLimitCorrelate `field:"optional" json:"correlate" yaml:"correlate"`
	// A note that you can use to describe the reason for a rate limit.
	//
	// This value is sanitized and all tags are removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#description RateLimit#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether this ratelimit is currently disabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#disabled RateLimit#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#id RateLimit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/rate_limit#match RateLimit#match}
	Match *RateLimitMatch `field:"optional" json:"match" yaml:"match"`
}

