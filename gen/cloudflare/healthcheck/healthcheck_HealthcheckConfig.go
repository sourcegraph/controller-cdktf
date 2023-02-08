package healthcheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcheckConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#address Healthcheck#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#name Healthcheck#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#type Healthcheck#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#zone_id Healthcheck#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#allow_insecure Healthcheck#allow_insecure}.
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#check_regions Healthcheck#check_regions}.
	CheckRegions *[]*string `field:"optional" json:"checkRegions" yaml:"checkRegions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#consecutive_fails Healthcheck#consecutive_fails}.
	ConsecutiveFails *float64 `field:"optional" json:"consecutiveFails" yaml:"consecutiveFails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#consecutive_successes Healthcheck#consecutive_successes}.
	ConsecutiveSuccesses *float64 `field:"optional" json:"consecutiveSuccesses" yaml:"consecutiveSuccesses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#description Healthcheck#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#expected_body Healthcheck#expected_body}.
	ExpectedBody *string `field:"optional" json:"expectedBody" yaml:"expectedBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#expected_codes Healthcheck#expected_codes}.
	ExpectedCodes *[]*string `field:"optional" json:"expectedCodes" yaml:"expectedCodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#follow_redirects Healthcheck#follow_redirects}.
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#header Healthcheck#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#id Healthcheck#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#interval Healthcheck#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#method Healthcheck#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#notification_email_addresses Healthcheck#notification_email_addresses}.
	NotificationEmailAddresses *[]*string `field:"optional" json:"notificationEmailAddresses" yaml:"notificationEmailAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#notification_suspended Healthcheck#notification_suspended}.
	NotificationSuspended interface{} `field:"optional" json:"notificationSuspended" yaml:"notificationSuspended"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#path Healthcheck#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#port Healthcheck#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#retries Healthcheck#retries}.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#suspended Healthcheck#suspended}.
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#timeout Healthcheck#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#timeouts Healthcheck#timeouts}
	Timeouts *HealthcheckTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

