package googlenetworkservicesservicelbpolicies


type GoogleNetworkServicesServiceLbPoliciesFailoverConfig struct {
	// Optional.
	//
	// The percentage threshold that a load balancer will begin to send traffic to failover backends. If the percentage of endpoints in a MIG/NEG is smaller than this value, traffic would be sent to failover backends if possible. This field should be set to a value between 1 and 99. The default value is 50 for Global external HTTP(S) load balancer (classic) and Proxyless service mesh, and 70 for others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#failover_health_threshold GoogleNetworkServicesServiceLbPolicies#failover_health_threshold}
	FailoverHealthThreshold *float64 `field:"required" json:"failoverHealthThreshold" yaml:"failoverHealthThreshold"`
}

