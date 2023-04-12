package googlednspolicy


type GoogleDnsPolicyAlternativeNameServerConfig struct {
	// target_name_servers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_policy#target_name_servers GoogleDnsPolicy#target_name_servers}
	TargetNameServers interface{} `field:"required" json:"targetNameServers" yaml:"targetNameServers"`
}

