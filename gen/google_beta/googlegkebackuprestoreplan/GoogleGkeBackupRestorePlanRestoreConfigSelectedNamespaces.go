package googlegkebackuprestoreplan


type GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespaces struct {
	// A list of Kubernetes Namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#namespaces GoogleGkeBackupRestorePlan#namespaces}
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

