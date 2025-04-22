package googlebackupdrbackupplan


type GoogleBackupDrBackupPlanBackupRules struct {
	// Configures the duration for which backup data will be kept.
	//
	// The value should be greater than or equal to minimum enforced retention of the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#backup_retention_days GoogleBackupDrBackupPlan#backup_retention_days}
	BackupRetentionDays *float64 `field:"required" json:"backupRetentionDays" yaml:"backupRetentionDays"`
	// The unique ID of this 'BackupRule'. The 'rule_id' is unique per 'BackupPlan'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#rule_id GoogleBackupDrBackupPlan#rule_id}
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// standard_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_plan#standard_schedule GoogleBackupDrBackupPlan#standard_schedule}
	StandardSchedule *GoogleBackupDrBackupPlanBackupRulesStandardSchedule `field:"required" json:"standardSchedule" yaml:"standardSchedule"`
}

