package healthcheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcheckConfig struct {
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
	// The hostname or IP address of the origin server to run health checks on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#address Healthcheck#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens, and underscores are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#name Healthcheck#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The protocol to use for the health check. Available values: `TCP`, `HTTP`, `HTTPS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#type Healthcheck#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#zone_id Healthcheck#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Do not validate the certificate when the health check uses HTTPS. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#allow_insecure Healthcheck#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// A list of regions from which to run health checks.
	//
	// If not set, Cloudflare will pick a default region. Available values: `WNAM`, `ENAM`, `WEU`, `EEU`, `NSAM`, `SSAM`, `OC`, `ME`, `NAF`, `SAF`, `IN`, `SEAS`, `NEAS`, `ALL_REGIONS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#check_regions Healthcheck#check_regions}
	CheckRegions *[]*string `field:"optional" json:"checkRegions" yaml:"checkRegions"`
	// The number of consecutive fails required from a health check before changing the health to unhealthy. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#consecutive_fails Healthcheck#consecutive_fails}
	ConsecutiveFails *float64 `field:"optional" json:"consecutiveFails" yaml:"consecutiveFails"`
	// The number of consecutive successes required from a health check before changing the health to healthy. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#consecutive_successes Healthcheck#consecutive_successes}
	ConsecutiveSuccesses *float64 `field:"optional" json:"consecutiveSuccesses" yaml:"consecutiveSuccesses"`
	// A human-readable description of the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#description Healthcheck#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A case-insensitive sub-string to look for in the response body.
	//
	// If this string is not found the origin will be marked as unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#expected_body Healthcheck#expected_body}
	ExpectedBody *string `field:"optional" json:"expectedBody" yaml:"expectedBody"`
	// The expected HTTP response codes (e.g. '200') or code ranges (e.g. '2xx' for all codes starting with 2) of the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#expected_codes Healthcheck#expected_codes}
	ExpectedCodes *[]*string `field:"optional" json:"expectedCodes" yaml:"expectedCodes"`
	// Follow redirects if the origin returns a 3xx status code. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#follow_redirects Healthcheck#follow_redirects}
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#header Healthcheck#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#id Healthcheck#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The interval between each health check.
	//
	// Shorter intervals may give quicker notifications if the origin status changes, but will increase the load on the origin as we check from multiple locations. Defaults to `60`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#interval Healthcheck#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The HTTP method to use for the health check. Available values: `connection_established`, `GET`, `HEAD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#method Healthcheck#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The endpoint path to health check against. Defaults to `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#path Healthcheck#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port number to connect to for the health check. Defaults to `80`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#port Healthcheck#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
	//
	// Retries are attempted immediately. Defaults to `2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#retries Healthcheck#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// If suspended, no health checks are sent to the origin. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#suspended Healthcheck#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// The timeout (in seconds) before marking the health check as failed. Defaults to `5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#timeout Healthcheck#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/healthcheck#timeouts Healthcheck#timeouts}
	Timeouts *HealthcheckTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

