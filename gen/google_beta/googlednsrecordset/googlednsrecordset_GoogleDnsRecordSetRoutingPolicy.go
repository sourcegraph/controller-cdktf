package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicy struct {
	// geo block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#geo GoogleDnsRecordSet#geo}
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// wrr block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#wrr GoogleDnsRecordSet#wrr}
	Wrr interface{} `field:"optional" json:"wrr" yaml:"wrr"`
}

