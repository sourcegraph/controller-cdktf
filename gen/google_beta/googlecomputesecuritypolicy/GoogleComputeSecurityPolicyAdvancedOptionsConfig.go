package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyAdvancedOptionsConfig struct {
	// json_custom_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_security_policy#json_custom_config GoogleComputeSecurityPolicy#json_custom_config}
	JsonCustomConfig *GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig `field:"optional" json:"jsonCustomConfig" yaml:"jsonCustomConfig"`
	// JSON body parsing. Supported values include: "DISABLED", "STANDARD".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_security_policy#json_parsing GoogleComputeSecurityPolicy#json_parsing}
	JsonParsing *string `field:"optional" json:"jsonParsing" yaml:"jsonParsing"`
	// Logging level. Supported values include: "NORMAL", "VERBOSE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_security_policy#log_level GoogleComputeSecurityPolicy#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The maximum request size chosen by the customer with Waf enabled.
	//
	// Values supported are "8KB", "16KB, "32KB", "48KB" and "64KB". Values are case insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_security_policy#request_body_inspection_size GoogleComputeSecurityPolicy#request_body_inspection_size}
	RequestBodyInspectionSize *string `field:"optional" json:"requestBodyInspectionSize" yaml:"requestBodyInspectionSize"`
	// An optional list of case-insensitive request header names to use for resolving the callers client IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_security_policy#user_ip_request_headers GoogleComputeSecurityPolicy#user_ip_request_headers}
	UserIpRequestHeaders *[]*string `field:"optional" json:"userIpRequestHeaders" yaml:"userIpRequestHeaders"`
}

