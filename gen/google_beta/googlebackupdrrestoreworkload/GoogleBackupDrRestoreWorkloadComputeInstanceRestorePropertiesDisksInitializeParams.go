package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksInitializeParams struct {
	// Optional. Specifies the disk name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#disk_name GoogleBackupDrRestoreWorkload#disk_name}
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// Optional. URL of the zone where the disk should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#replica_zones GoogleBackupDrRestoreWorkload#replica_zones}
	ReplicaZones *[]*string `field:"optional" json:"replicaZones" yaml:"replicaZones"`
}

