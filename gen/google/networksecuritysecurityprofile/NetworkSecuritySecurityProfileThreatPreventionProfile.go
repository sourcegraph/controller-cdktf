package networksecuritysecurityprofile

type NetworkSecuritySecurityProfileThreatPreventionProfile struct {
	// antivirus_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_security_profile#antivirus_overrides NetworkSecuritySecurityProfile#antivirus_overrides}
	AntivirusOverrides any `field:"optional" json:"antivirusOverrides" yaml:"antivirusOverrides"`
	// severity_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_security_profile#severity_overrides NetworkSecuritySecurityProfile#severity_overrides}
	SeverityOverrides any `field:"optional" json:"severityOverrides" yaml:"severityOverrides"`
	// threat_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_security_profile#threat_overrides NetworkSecuritySecurityProfile#threat_overrides}
	ThreatOverrides any `field:"optional" json:"threatOverrides" yaml:"threatOverrides"`
}
