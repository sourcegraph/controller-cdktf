package monitoraction

type MonitorActionEmail struct {
	// Template string used to fill body of the email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#body_template MonitorAction#body_template}
	BodyTemplate *string `field:"required" json:"bodyTemplate" yaml:"bodyTemplate"`
	// Template string used to build subject line.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#subject_template MonitorAction#subject_template}
	SubjectTemplate *string `field:"required" json:"subjectTemplate" yaml:"subjectTemplate"`
	// Email address(es) to send alert to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#target_addresses MonitorAction#target_addresses}
	TargetAddresses *[]*string `field:"required" json:"targetAddresses" yaml:"targetAddresses"`
	// send the email as html allowing rich formatting in the body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#is_html MonitorAction#is_html}
	IsHtml any `field:"optional" json:"isHtml" yaml:"isHtml"`
}
