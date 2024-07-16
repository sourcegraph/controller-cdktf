package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicy struct {
	// Specifies whether to enable fencing for geo queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#enable_geo_fencing GoogleDnsRecordSet#enable_geo_fencing}
	EnableGeoFencing interface{} `field:"optional" json:"enableGeoFencing" yaml:"enableGeoFencing"`
	// geo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#geo GoogleDnsRecordSet#geo}
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// primary_backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#primary_backup GoogleDnsRecordSet#primary_backup}
	PrimaryBackup *GoogleDnsRecordSetRoutingPolicyPrimaryBackup `field:"optional" json:"primaryBackup" yaml:"primaryBackup"`
	// wrr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#wrr GoogleDnsRecordSet#wrr}
	Wrr interface{} `field:"optional" json:"wrr" yaml:"wrr"`
}

