package googlegkebackuprestoreplan


type GoogleGkeBackupRestorePlanRestoreConfigRestoreOrderGroupKindDependencies struct {
	// requiring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#requiring GoogleGkeBackupRestorePlan#requiring}
	Requiring *GoogleGkeBackupRestorePlanRestoreConfigRestoreOrderGroupKindDependenciesRequiring `field:"required" json:"requiring" yaml:"requiring"`
	// satisfying block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#satisfying GoogleGkeBackupRestorePlan#satisfying}
	Satisfying *GoogleGkeBackupRestorePlanRestoreConfigRestoreOrderGroupKindDependenciesSatisfying `field:"required" json:"satisfying" yaml:"satisfying"`
}

