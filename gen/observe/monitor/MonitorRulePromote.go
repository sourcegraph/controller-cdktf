package monitor


type MonitorRulePromote struct {
	// Key used to deduplicate and and uniquely identify each event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#primary_key Monitor#primary_key}
	PrimaryKey *[]*string `field:"required" json:"primaryKey" yaml:"primaryKey"`
	// Dataset field used to set the description for this monitor's events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#description_field Monitor#description_field}
	DescriptionField *string `field:"optional" json:"descriptionField" yaml:"descriptionField"`
	// The dataset field used to group notifications for this monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#kind_field Monitor#kind_field}
	KindField *string `field:"optional" json:"kindField" yaml:"kindField"`
}

