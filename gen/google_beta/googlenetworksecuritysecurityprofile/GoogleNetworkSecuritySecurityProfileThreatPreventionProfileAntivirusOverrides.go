package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileThreatPreventionProfileAntivirusOverrides struct {
	// Threat action override. For some threat types, only a subset of actions applies. Possible values: ["ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_security_profile#action GoogleNetworkSecuritySecurityProfile#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Required protocol to match. Possible values: ["SMTP", "SMB", "POP3", "IMAP", "HTTP2", "HTTP", "FTP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_security_profile#protocol GoogleNetworkSecuritySecurityProfile#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

