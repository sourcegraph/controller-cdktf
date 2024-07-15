package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides struct {
	// Threat action override. Possible values: ["ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_security_profile#action GoogleNetworkSecuritySecurityProfile#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Severity level to match. Possible values: ["CRITICAL", "HIGH", "INFORMATIONAL", "LOW", "MEDIUM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_security_profile#severity GoogleNetworkSecuritySecurityProfile#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
}

