package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyWrr struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#rrdatas GoogleDnsRecordSet#rrdatas}.
	Rrdatas *[]*string `field:"required" json:"rrdatas" yaml:"rrdatas"`
	// The ratio of traffic routed to the target.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#weight GoogleDnsRecordSet#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

