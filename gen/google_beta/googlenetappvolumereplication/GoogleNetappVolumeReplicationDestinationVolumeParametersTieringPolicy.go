package googlenetappvolumereplication


type GoogleNetappVolumeReplicationDestinationVolumeParametersTieringPolicy struct {
	// Optional.
	//
	// Time in days to mark the volume's data block as cold and make it eligible for tiering, can be range from 2-183.
	// Default is 31.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_netapp_volume_replication#cooling_threshold_days GoogleNetappVolumeReplication#cooling_threshold_days}
	CoolingThresholdDays *float64 `field:"optional" json:"coolingThresholdDays" yaml:"coolingThresholdDays"`
	// Optional.
	//
	// Flag indicating if the volume has tiering policy enable/pause. Default is PAUSED. Default value: "PAUSED" Possible values: ["ENABLED", "PAUSED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_netapp_volume_replication#tier_action GoogleNetappVolumeReplication#tier_action}
	TierAction *string `field:"optional" json:"tierAction" yaml:"tierAction"`
}

