package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyWrr struct {
	// The ratio of traffic routed to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#weight GoogleDnsRecordSet#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// health_checked_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#health_checked_targets GoogleDnsRecordSet#health_checked_targets}
	HealthCheckedTargets *GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargets `field:"optional" json:"healthCheckedTargets" yaml:"healthCheckedTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#rrdatas GoogleDnsRecordSet#rrdatas}.
	Rrdatas *[]*string `field:"optional" json:"rrdatas" yaml:"rrdatas"`
}

