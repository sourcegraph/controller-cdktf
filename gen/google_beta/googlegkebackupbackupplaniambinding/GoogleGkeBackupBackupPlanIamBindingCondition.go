package googlegkebackupbackupplaniambinding


type GoogleGkeBackupBackupPlanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan_iam_binding#expression GoogleGkeBackupBackupPlanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan_iam_binding#title GoogleGkeBackupBackupPlanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan_iam_binding#description GoogleGkeBackupBackupPlanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

