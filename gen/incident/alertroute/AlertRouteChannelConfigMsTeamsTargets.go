package alertroute


type AlertRouteChannelConfigMsTeamsTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#binding AlertRoute#binding}.
	Binding *AlertRouteChannelConfigMsTeamsTargetsBinding `field:"required" json:"binding" yaml:"binding"`
	// The visibility of the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#channel_visibility AlertRoute#channel_visibility}
	ChannelVisibility *string `field:"required" json:"channelVisibility" yaml:"channelVisibility"`
}

