package googlegkebackuprestoreplan


type GoogleGkeBackupRestorePlanRestoreConfigTransformationRules struct {
	// field_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#field_actions GoogleGkeBackupRestorePlan#field_actions}
	FieldActions interface{} `field:"required" json:"fieldActions" yaml:"fieldActions"`
	// The description is a user specified string description of the transformation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#description GoogleGkeBackupRestorePlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// resource_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#resource_filter GoogleGkeBackupRestorePlan#resource_filter}
	ResourceFilter *GoogleGkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter `field:"optional" json:"resourceFilter" yaml:"resourceFilter"`
}

