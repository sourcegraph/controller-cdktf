package monitorv2

type MonitorV2Actions struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#action MonitorV2#action}
	Action *MonitorV2ActionsAction `field:"optional" json:"action" yaml:"action"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#conditions MonitorV2#conditions}
	Conditions *MonitorV2ActionsConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// The alarm level(s) at which this monitor should trigger this shared action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#levels MonitorV2#levels}
	Levels *[]*string `field:"optional" json:"levels" yaml:"levels"`
	// The OID of this shared action. This should be used for existing shared actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#oid MonitorV2#oid}
	Oid *string `field:"optional" json:"oid" yaml:"oid"`
	// If true, notifications will be sent if the monitor stops triggering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#send_end_notifications MonitorV2#send_end_notifications}
	SendEndNotifications any `field:"optional" json:"sendEndNotifications" yaml:"sendEndNotifications"`
	// Determines how frequently you will be reminded of an ongoing alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#send_reminders_interval MonitorV2#send_reminders_interval}
	SendRemindersInterval *string `field:"optional" json:"sendRemindersInterval" yaml:"sendRemindersInterval"`
}
