package googlegkebackuprestoreplaniambinding


type GoogleGkeBackupRestorePlanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan_iam_binding#expression GoogleGkeBackupRestorePlanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan_iam_binding#title GoogleGkeBackupRestorePlanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan_iam_binding#description GoogleGkeBackupRestorePlanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

