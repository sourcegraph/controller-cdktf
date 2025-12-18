package monitor


type MonitorNotificationSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#importance Monitor#importance}.
	Importance *string `field:"optional" json:"importance" yaml:"importance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#merge Monitor#merge}.
	Merge *string `field:"optional" json:"merge" yaml:"merge"`
	// Enables a final update when a monitor notification is closed (no longer triggered).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#notify_on_close Monitor#notify_on_close}
	NotifyOnClose interface{} `field:"optional" json:"notifyOnClose" yaml:"notifyOnClose"`
	// How often to send reminders when a monitor notification is triggered. To disable reminder notifications, omit this attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#reminder_frequency Monitor#reminder_frequency}
	ReminderFrequency *string `field:"optional" json:"reminderFrequency" yaml:"reminderFrequency"`
}

