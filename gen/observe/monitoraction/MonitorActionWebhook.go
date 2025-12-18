package monitoraction


type MonitorActionWebhook struct {
	// Template string used to fill body of request sent to URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#body_template MonitorAction#body_template}
	BodyTemplate *string `field:"required" json:"bodyTemplate" yaml:"bodyTemplate"`
	// Template string used to generate the request URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#url_template MonitorAction#url_template}
	UrlTemplate *string `field:"required" json:"urlTemplate" yaml:"urlTemplate"`
	// Extra headers to send along with the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#headers MonitorAction#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// HTTP method to use for the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#method MonitorAction#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
}

