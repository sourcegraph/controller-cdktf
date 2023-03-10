package gameliftgameservergroup


type GameliftGameServerGroupAutoScalingPolicy struct {
	// target_tracking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_game_server_group#target_tracking_configuration GameliftGameServerGroup#target_tracking_configuration}
	TargetTrackingConfiguration *GameliftGameServerGroupAutoScalingPolicyTargetTrackingConfiguration `field:"required" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_game_server_group#estimated_instance_warmup GameliftGameServerGroup#estimated_instance_warmup}.
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
}

