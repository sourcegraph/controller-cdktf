package googlegkebackupbackupplaniammember


type GoogleGkeBackupBackupPlanIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan_iam_member#expression GoogleGkeBackupBackupPlanIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan_iam_member#title GoogleGkeBackupBackupPlanIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_backup_plan_iam_member#description GoogleGkeBackupBackupPlanIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

