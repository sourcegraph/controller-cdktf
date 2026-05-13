package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeature struct {
	// Possible values: ["FEATURE_TYPE_UNSPECIFIED", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "MULTI_IP_SUBNET", "UEFI_COMPATIBLE", "SECURE_BOOT", "GVNIC", "SEV_CAPABLE", "BARE_METAL_LINUX_COMPATIBLE", "SUSPEND_RESUME_COMPATIBLE", "SEV_LIVE_MIGRATABLE", "SEV_SNP_CAPABLE", "TDX_CAPABLE", "IDPF", "SEV_LIVE_MIGRATABLE_V2"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#type GoogleBackupDrRestoreWorkload#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

