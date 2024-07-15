package googlenetappvolume


type GoogleNetappVolumeSnapshotPolicyHourlySchedule struct {
	// The maximum number of snapshots to keep for the hourly schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#snapshots_to_keep GoogleNetappVolume#snapshots_to_keep}
	SnapshotsToKeep *float64 `field:"required" json:"snapshotsToKeep" yaml:"snapshotsToKeep"`
	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#minute GoogleNetappVolume#minute}
	Minute *float64 `field:"optional" json:"minute" yaml:"minute"`
}

