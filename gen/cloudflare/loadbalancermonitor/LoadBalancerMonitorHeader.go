package loadbalancermonitor


type LoadBalancerMonitorHeader struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#header LoadBalancerMonitor#header}
	Header *string `field:"required" json:"header" yaml:"header"`
	// A list of values for the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_monitor#values LoadBalancerMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

