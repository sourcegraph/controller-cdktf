package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaceLabelsResourceLabels struct {
	// The key of the kubernetes label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_backup_backup_plan#key GoogleGkeBackupBackupPlan#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the Label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_backup_backup_plan#value GoogleGkeBackupBackupPlan#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

