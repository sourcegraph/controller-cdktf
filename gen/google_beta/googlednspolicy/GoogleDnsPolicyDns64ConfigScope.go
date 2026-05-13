package googlednspolicy


type GoogleDnsPolicyDns64ConfigScope struct {
	// Controls whether DNS64 is enabled globally at the network level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dns_policy#all_queries GoogleDnsPolicy#all_queries}
	AllQueries interface{} `field:"optional" json:"allQueries" yaml:"allQueries"`
}

