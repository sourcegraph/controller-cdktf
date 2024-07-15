package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNames struct {
	// The name of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#name GoogleGkeBackupBackupPlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#namespace GoogleGkeBackupBackupPlan#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

