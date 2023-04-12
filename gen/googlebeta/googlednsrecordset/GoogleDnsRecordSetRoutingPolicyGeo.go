package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyGeo struct {
	// The location name defined in Google Cloud.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#location GoogleDnsRecordSet#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#rrdatas GoogleDnsRecordSet#rrdatas}.
	Rrdatas *[]*string `field:"required" json:"rrdatas" yaml:"rrdatas"`
}

