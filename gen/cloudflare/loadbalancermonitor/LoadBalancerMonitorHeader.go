package loadbalancermonitor


type LoadBalancerMonitorHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#header LoadBalancerMonitor#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_monitor#values LoadBalancerMonitor#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

