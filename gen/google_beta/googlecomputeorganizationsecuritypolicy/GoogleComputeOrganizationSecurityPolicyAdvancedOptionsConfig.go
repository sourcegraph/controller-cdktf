package googlecomputeorganizationsecuritypolicy


type GoogleComputeOrganizationSecurityPolicyAdvancedOptionsConfig struct {
	// json_custom_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_organization_security_policy#json_custom_config GoogleComputeOrganizationSecurityPolicy#json_custom_config}
	JsonCustomConfig *GoogleComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig `field:"optional" json:"jsonCustomConfig" yaml:"jsonCustomConfig"`
	// JSON body parsing. Supported values include: "DISABLED", "STANDARD", "STANDARD_WITH_GRAPHQL". Possible values: ["DISABLED", "STANDARD", "STANDARD_WITH_GRAPHQL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_organization_security_policy#json_parsing GoogleComputeOrganizationSecurityPolicy#json_parsing}
	JsonParsing *string `field:"optional" json:"jsonParsing" yaml:"jsonParsing"`
	// Logging level. Supported values include: "NORMAL", "VERBOSE". Possible values: ["NORMAL", "VERBOSE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_organization_security_policy#log_level GoogleComputeOrganizationSecurityPolicy#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The maximum request size chosen by the customer with Waf enabled.
	//
	// Values supported are "8KB", "16KB", "32KB", "48KB" and "64KB".
	// Values are case insensitive. Possible values: ["8KB", "16KB", "32KB", "48KB", "64KB"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_organization_security_policy#request_body_inspection_size GoogleComputeOrganizationSecurityPolicy#request_body_inspection_size}
	RequestBodyInspectionSize *string `field:"optional" json:"requestBodyInspectionSize" yaml:"requestBodyInspectionSize"`
	// An optional list of case-insensitive request header names to use for resolving the client source IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_organization_security_policy#user_ip_request_headers GoogleComputeOrganizationSecurityPolicy#user_ip_request_headers}
	UserIpRequestHeaders *[]*string `field:"optional" json:"userIpRequestHeaders" yaml:"userIpRequestHeaders"`
}

