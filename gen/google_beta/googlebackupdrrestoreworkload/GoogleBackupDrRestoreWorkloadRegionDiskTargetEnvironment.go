package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironment struct {
	// Required. Target project for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#project GoogleBackupDrRestoreWorkload#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Required. Target region for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#region GoogleBackupDrRestoreWorkload#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Required. Target URLs of the replica zones for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#replica_zones GoogleBackupDrRestoreWorkload#replica_zones}
	ReplicaZones *[]*string `field:"required" json:"replicaZones" yaml:"replicaZones"`
}

