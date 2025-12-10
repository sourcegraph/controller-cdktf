package escalationpath


type EscalationPathPathIfElseThenPathNotifyChannelTargets struct {
	// Uniquely identifies an entity of this type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#id EscalationPath#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Controls what type of entity this target identifies, such as EscalationPolicy or User.
	//
	// Possible values are: `schedule`, `user`, `slack_channel`, `msteams_channel`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#type EscalationPath#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The urgency of this escalation path target. Possible values are: `high`, `low`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#urgency EscalationPath#urgency}
	Urgency *string `field:"required" json:"urgency" yaml:"urgency"`
	// Only set for schedule targets, and either currently_on_call, all_users or all_users_for_rota and specifies which users to fetch from the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#schedule_mode EscalationPath#schedule_mode}
	ScheduleMode *string `field:"optional" json:"scheduleMode" yaml:"scheduleMode"`
}

