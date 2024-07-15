package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigHttpCheckServiceAgentAuthentication struct {
	// The type of authentication to use. Possible values: ["SERVICE_AGENT_AUTHENTICATION_TYPE_UNSPECIFIED", "OIDC_TOKEN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#type GoogleMonitoringUptimeCheckConfig#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

