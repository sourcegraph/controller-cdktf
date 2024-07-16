package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigContentMatchers struct {
	// String or regex content to match (max 1024 bytes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#content GoogleMonitoringUptimeCheckConfig#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// json_path_matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#json_path_matcher GoogleMonitoringUptimeCheckConfig#json_path_matcher}
	JsonPathMatcher *GoogleMonitoringUptimeCheckConfigContentMatchersJsonPathMatcher `field:"optional" json:"jsonPathMatcher" yaml:"jsonPathMatcher"`
	// The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
	//
	// Default value: "CONTAINS_STRING" Possible values: ["CONTAINS_STRING", "NOT_CONTAINS_STRING", "MATCHES_REGEX", "NOT_MATCHES_REGEX", "MATCHES_JSON_PATH", "NOT_MATCHES_JSON_PATH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#matcher GoogleMonitoringUptimeCheckConfig#matcher}
	Matcher *string `field:"optional" json:"matcher" yaml:"matcher"`
}

