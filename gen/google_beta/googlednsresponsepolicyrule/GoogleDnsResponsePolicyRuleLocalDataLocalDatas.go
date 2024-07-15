package googlednsresponsepolicyrule


type GoogleDnsResponsePolicyRuleLocalDataLocalDatas struct {
	// For example, www.example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_response_policy_rule#name GoogleDnsResponsePolicyRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One of valid DNS resource types.
	//
	// Possible values: ["A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "HTTPS", "IPSECVPNKEY", "MX", "NAPTR", "NS", "PTR", "SOA", "SPF", "SRV", "SSHFP", "SVCB", "TLSA", "TXT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_response_policy_rule#type GoogleDnsResponsePolicyRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_response_policy_rule#rrdatas GoogleDnsResponsePolicyRule#rrdatas}
	Rrdatas *[]*string `field:"optional" json:"rrdatas" yaml:"rrdatas"`
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_response_policy_rule#ttl GoogleDnsResponsePolicyRule#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

