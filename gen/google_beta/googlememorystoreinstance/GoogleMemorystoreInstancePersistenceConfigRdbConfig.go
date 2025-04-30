package googlememorystoreinstance


type GoogleMemorystoreInstancePersistenceConfigRdbConfig struct {
	// Optional. Period between RDB snapshots.   Possible values:  ONE_HOUR SIX_HOURS TWELVE_HOURS TWENTY_FOUR_HOURS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#rdb_snapshot_period GoogleMemorystoreInstance#rdb_snapshot_period}
	RdbSnapshotPeriod *string `field:"optional" json:"rdbSnapshotPeriod" yaml:"rdbSnapshotPeriod"`
	// Optional.
	//
	// Time that the first snapshot was/will be attempted, and to which future
	// snapshots will be aligned. If not provided, the current time will be
	// used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#rdb_snapshot_start_time GoogleMemorystoreInstance#rdb_snapshot_start_time}
	RdbSnapshotStartTime *string `field:"optional" json:"rdbSnapshotStartTime" yaml:"rdbSnapshotStartTime"`
}

