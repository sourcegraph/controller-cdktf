package googlegkebackuprestoreplaniammember


type GoogleGkeBackupRestorePlanIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan_iam_member#expression GoogleGkeBackupRestorePlanIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan_iam_member#title GoogleGkeBackupRestorePlanIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan_iam_member#description GoogleGkeBackupRestorePlanIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

