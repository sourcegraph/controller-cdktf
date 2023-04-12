package loadbalancermonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#allow_insecure LoadBalancerMonitor#allow_insecure}.
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#description LoadBalancerMonitor#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#expected_body LoadBalancerMonitor#expected_body}.
	ExpectedBody *string `field:"optional" json:"expectedBody" yaml:"expectedBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#expected_codes LoadBalancerMonitor#expected_codes}.
	ExpectedCodes *string `field:"optional" json:"expectedCodes" yaml:"expectedCodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#follow_redirects LoadBalancerMonitor#follow_redirects}.
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#header LoadBalancerMonitor#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#id LoadBalancerMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#interval LoadBalancerMonitor#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#method LoadBalancerMonitor#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#path LoadBalancerMonitor#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#port LoadBalancerMonitor#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#probe_zone LoadBalancerMonitor#probe_zone}.
	ProbeZone *string `field:"optional" json:"probeZone" yaml:"probeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#retries LoadBalancerMonitor#retries}.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#timeout LoadBalancerMonitor#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#type LoadBalancerMonitor#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

