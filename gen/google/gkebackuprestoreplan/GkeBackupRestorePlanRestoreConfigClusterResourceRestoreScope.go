package gkebackuprestoreplan

type GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope struct {
	// If True, all valid cluster-scoped resources will be restored. Mutually exclusive to any other field in 'clusterResourceRestoreScope'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_backup_restore_plan#all_group_kinds GkeBackupRestorePlan#all_group_kinds}
	AllGroupKinds any `field:"optional" json:"allGroupKinds" yaml:"allGroupKinds"`
	// excluded_group_kinds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_backup_restore_plan#excluded_group_kinds GkeBackupRestorePlan#excluded_group_kinds}
	ExcludedGroupKinds any `field:"optional" json:"excludedGroupKinds" yaml:"excludedGroupKinds"`
	// If True, no cluster-scoped resources will be restored. Mutually exclusive to any other field in 'clusterResourceRestoreScope'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_backup_restore_plan#no_group_kinds GkeBackupRestorePlan#no_group_kinds}
	NoGroupKinds any `field:"optional" json:"noGroupKinds" yaml:"noGroupKinds"`
	// selected_group_kinds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gke_backup_restore_plan#selected_group_kinds GkeBackupRestorePlan#selected_group_kinds}
	SelectedGroupKinds any `field:"optional" json:"selectedGroupKinds" yaml:"selectedGroupKinds"`
}
