package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupConfig struct {
	// If True, include all namespaced resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#all_namespaces GoogleGkeBackupBackupPlan#all_namespaces}
	AllNamespaces interface{} `field:"optional" json:"allNamespaces" yaml:"allNamespaces"`
	// encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#encryption_key GoogleGkeBackupBackupPlan#encryption_key}
	EncryptionKey *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// This flag specifies whether Kubernetes Secret resources should be included when they fall into the scope of Backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#include_secrets GoogleGkeBackupBackupPlan#include_secrets}
	IncludeSecrets interface{} `field:"optional" json:"includeSecrets" yaml:"includeSecrets"`
	// This flag specifies whether volume data should be backed up when PVCs are included in the scope of a Backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#include_volume_data GoogleGkeBackupBackupPlan#include_volume_data}
	IncludeVolumeData interface{} `field:"optional" json:"includeVolumeData" yaml:"includeVolumeData"`
	// This flag specifies whether Backups will not fail when Backup for GKE detects Kubernetes configuration that is non-standard or requires additional setup to restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#permissive_mode GoogleGkeBackupBackupPlan#permissive_mode}
	PermissiveMode interface{} `field:"optional" json:"permissiveMode" yaml:"permissiveMode"`
	// selected_applications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#selected_applications GoogleGkeBackupBackupPlan#selected_applications}
	SelectedApplications *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications `field:"optional" json:"selectedApplications" yaml:"selectedApplications"`
	// selected_namespaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan#selected_namespaces GoogleGkeBackupBackupPlan#selected_namespaces}
	SelectedNamespaces *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces `field:"optional" json:"selectedNamespaces" yaml:"selectedNamespaces"`
}

