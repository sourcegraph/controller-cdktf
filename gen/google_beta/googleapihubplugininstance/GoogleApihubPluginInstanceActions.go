package googleapihubplugininstance


type GoogleApihubPluginInstanceActions struct {
	// This should map to one of the action id specified in actions_config in the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#action_id GoogleApihubPluginInstance#action_id}
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// curation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#curation_config GoogleApihubPluginInstance#curation_config}
	CurationConfig *GoogleApihubPluginInstanceActionsCurationConfig `field:"optional" json:"curationConfig" yaml:"curationConfig"`
	// The schedule for this plugin instance action.
	//
	// This can only be set if the
	// plugin supports API_HUB_SCHEDULE_TRIGGER mode for this action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#schedule_cron_expression GoogleApihubPluginInstance#schedule_cron_expression}
	ScheduleCronExpression *string `field:"optional" json:"scheduleCronExpression" yaml:"scheduleCronExpression"`
	// The time zone for the schedule cron expression. If not provided, UTC will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin_instance#schedule_time_zone GoogleApihubPluginInstance#schedule_time_zone}
	ScheduleTimeZone *string `field:"optional" json:"scheduleTimeZone" yaml:"scheduleTimeZone"`
}

