package dnsresponsepolicyrule

type DnsResponsePolicyRuleLocalData struct {
	// local_datas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dns_response_policy_rule#local_datas DnsResponsePolicyRule#local_datas}
	LocalDatas any `field:"required" json:"localDatas" yaml:"localDatas"`
}
