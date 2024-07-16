package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides struct {
	// Threat action. Possible values: ["ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_security_profile#action GoogleNetworkSecuritySecurityProfile#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Vendor-specific ID of a threat to override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_security_profile#threat_id GoogleNetworkSecuritySecurityProfile#threat_id}
	ThreatId *string `field:"required" json:"threatId" yaml:"threatId"`
}

