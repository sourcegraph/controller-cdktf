package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyPrimaryBackupPrimary struct {
	// The Internet IP addresses to be health checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dns_record_set#external_endpoints GoogleDnsRecordSet#external_endpoints}
	ExternalEndpoints *[]*string `field:"optional" json:"externalEndpoints" yaml:"externalEndpoints"`
	// internal_load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dns_record_set#internal_load_balancers GoogleDnsRecordSet#internal_load_balancers}
	InternalLoadBalancers interface{} `field:"optional" json:"internalLoadBalancers" yaml:"internalLoadBalancers"`
}

