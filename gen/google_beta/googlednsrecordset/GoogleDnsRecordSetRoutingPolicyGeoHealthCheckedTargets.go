package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyGeoHealthCheckedTargets struct {
	// internal_load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#internal_load_balancers GoogleDnsRecordSet#internal_load_balancers}
	InternalLoadBalancers interface{} `field:"required" json:"internalLoadBalancers" yaml:"internalLoadBalancers"`
}

