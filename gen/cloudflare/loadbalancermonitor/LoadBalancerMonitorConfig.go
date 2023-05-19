package loadbalancermonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerMonitorConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#account_id LoadBalancerMonitor#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Do not validate the certificate when monitor use HTTPS.  Only valid if `type` is "http" or "https".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#allow_insecure LoadBalancerMonitor#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// Free text description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#description LoadBalancerMonitor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A case-insensitive sub-string to look for in the response body.
	//
	// If this string is not found, the origin will be marked as unhealthy. Only valid if `type` is "http" or "https".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#expected_body LoadBalancerMonitor#expected_body}
	ExpectedBody *string `field:"optional" json:"expectedBody" yaml:"expectedBody"`
	// The expected HTTP response code or code range of the health check.
	//
	// Eg `2xx`. Only valid and required if `type` is "http" or "https".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#expected_codes LoadBalancerMonitor#expected_codes}
	ExpectedCodes *string `field:"optional" json:"expectedCodes" yaml:"expectedCodes"`
	// Follow redirects if returned by the origin. Only valid if `type` is "http" or "https".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#follow_redirects LoadBalancerMonitor#follow_redirects}
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#header LoadBalancerMonitor#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#id LoadBalancerMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The interval between each health check.
	//
	// Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations. Defaults to `60`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#interval LoadBalancerMonitor#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The method to use for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#method LoadBalancerMonitor#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The endpoint path to health check against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#path LoadBalancerMonitor#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port number to use for the healthcheck, required when creating a TCP monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#port LoadBalancerMonitor#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Assign this monitor to emulate the specified zone while probing. Only valid if `type` is "http" or "https".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#probe_zone LoadBalancerMonitor#probe_zone}
	ProbeZone *string `field:"optional" json:"probeZone" yaml:"probeZone"`
	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
	//
	// Retries are attempted immediately. Defaults to `2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#retries LoadBalancerMonitor#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The timeout (in seconds) before marking the health check as failed. Defaults to `5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#timeout LoadBalancerMonitor#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The protocol to use for the healthcheck. Available values: `http`, `https`, `tcp`, `udp_icmp`, `icmp_ping`, `smtp`. Defaults to `http`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#type LoadBalancerMonitor#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

