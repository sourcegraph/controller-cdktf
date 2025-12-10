package alertroute


type AlertRouteChannelConfig struct {
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#condition_groups AlertRoute#condition_groups}
	ConditionGroups interface{} `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#ms_teams_targets AlertRoute#ms_teams_targets}.
	MsTeamsTargets *AlertRouteChannelConfigMsTeamsTargets `field:"optional" json:"msTeamsTargets" yaml:"msTeamsTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#slack_targets AlertRoute#slack_targets}.
	SlackTargets *AlertRouteChannelConfigSlackTargets `field:"optional" json:"slackTargets" yaml:"slackTargets"`
}

