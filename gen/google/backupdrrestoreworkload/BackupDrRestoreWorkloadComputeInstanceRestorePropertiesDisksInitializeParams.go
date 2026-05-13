package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksInitializeParams struct {
	// Optional. Specifies the disk name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#disk_name BackupDrRestoreWorkload#disk_name}
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// Optional. URL of the zone where the disk should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#replica_zones BackupDrRestoreWorkload#replica_zones}
	ReplicaZones *[]*string `field:"optional" json:"replicaZones" yaml:"replicaZones"`
}

