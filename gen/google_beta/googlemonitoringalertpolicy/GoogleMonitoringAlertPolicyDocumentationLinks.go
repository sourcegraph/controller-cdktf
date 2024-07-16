package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyDocumentationLinks struct {
	// A short display name for the link.
	//
	// The display name must not be empty or exceed 63 characters. Example: "playbook".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#display_name GoogleMonitoringAlertPolicy#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The url of a webpage.
	//
	// A url can be templatized by using variables in the path or the query parameters. The total length of a URL should not exceed 2083 characters before and after variable expansion. Example: "https://my_domain.com/playbook?name=${resource.name}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#url GoogleMonitoringAlertPolicy#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

