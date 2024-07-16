package googleclouddomainsregistration


type GoogleClouddomainsRegistrationDnsSettingsCustomDns struct {
	// Required.
	//
	// A list of name servers that store the DNS zone for this domain. Each name server is a domain
	// name, with Unicode domain names expressed in Punycode format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#name_servers GoogleClouddomainsRegistration#name_servers}
	NameServers *[]*string `field:"required" json:"nameServers" yaml:"nameServers"`
	// ds_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#ds_records GoogleClouddomainsRegistration#ds_records}
	DsRecords interface{} `field:"optional" json:"dsRecords" yaml:"dsRecords"`
}

