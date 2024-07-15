package googlegkebackuprestoreplan


type GoogleGkeBackupRestorePlanRestoreConfigTransformationRulesFieldActions struct {
	// Specifies the operation to perform. Possible values: ["REMOVE", "MOVE", "COPY", "ADD", "TEST", "REPLACE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#op GoogleGkeBackupRestorePlan#op}
	Op *string `field:"required" json:"op" yaml:"op"`
	// A string containing a JSON Pointer value that references the location in the target document to move the value from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#from_path GoogleGkeBackupRestorePlan#from_path}
	FromPath *string `field:"optional" json:"fromPath" yaml:"fromPath"`
	// A string containing a JSON-Pointer value that references a location within the target document where the operation is performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#path GoogleGkeBackupRestorePlan#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// A string that specifies the desired value in string format to use for transformation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_backup_restore_plan#value GoogleGkeBackupRestorePlan#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

