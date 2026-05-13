package gkebackupbackupplan


type GkeBackupBackupPlanBackupConfigSelectedNamespaceLabelsResourceLabels struct {
	// The key of the kubernetes label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gke_backup_backup_plan#key GkeBackupBackupPlan#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the Label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gke_backup_backup_plan#value GkeBackupBackupPlan#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

