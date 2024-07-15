package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileThreatPreventionProfile struct {
	// severity_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_security_profile#severity_overrides GoogleNetworkSecuritySecurityProfile#severity_overrides}
	SeverityOverrides interface{} `field:"optional" json:"severityOverrides" yaml:"severityOverrides"`
	// threat_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_security_profile#threat_overrides GoogleNetworkSecuritySecurityProfile#threat_overrides}
	ThreatOverrides interface{} `field:"optional" json:"threatOverrides" yaml:"threatOverrides"`
}

