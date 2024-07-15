package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyAdvancedOptionsConfig struct {
	// json_custom_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#json_custom_config GoogleComputeSecurityPolicy#json_custom_config}
	JsonCustomConfig *GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig `field:"optional" json:"jsonCustomConfig" yaml:"jsonCustomConfig"`
	// JSON body parsing. Supported values include: "DISABLED", "STANDARD".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#json_parsing GoogleComputeSecurityPolicy#json_parsing}
	JsonParsing *string `field:"optional" json:"jsonParsing" yaml:"jsonParsing"`
	// Logging level. Supported values include: "NORMAL", "VERBOSE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#log_level GoogleComputeSecurityPolicy#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// An optional list of case-insensitive request header names to use for resolving the callers client IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#user_ip_request_headers GoogleComputeSecurityPolicy#user_ip_request_headers}
	UserIpRequestHeaders *[]*string `field:"optional" json:"userIpRequestHeaders" yaml:"userIpRequestHeaders"`
}

