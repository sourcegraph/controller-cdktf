package googlebackupdrbackupplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBackupDrBackupPlanConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#backup_plan_id GoogleBackupDrBackupPlan#backup_plan_id}
	BackupPlanId *string `field:"required" json:"backupPlanId" yaml:"backupPlanId"`
	// Backup vault where the backups gets stored using this Backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#backup_vault GoogleBackupDrBackupPlan#backup_vault}
	BackupVault *string `field:"required" json:"backupVault" yaml:"backupVault"`
	// The location for the backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#location GoogleBackupDrBackupPlan#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource type to which the 'BackupPlan' will be applied. Examples include, "compute.googleapis.com/Instance", "compute.googleapis.com/Disk", "sqladmin.googleapis.com/Instance", "alloydb.googleapis.com/Cluster", "file.googleapis.com/Instance" and "storage.googleapis.com/Bucket".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#resource_type GoogleBackupDrBackupPlan#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// backup_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#backup_rules GoogleBackupDrBackupPlan#backup_rules}
	BackupRules interface{} `field:"optional" json:"backupRules" yaml:"backupRules"`
	// The description allows for additional details about 'BackupPlan' and its use cases to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#description GoogleBackupDrBackupPlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// disk_backup_plan_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#disk_backup_plan_properties GoogleBackupDrBackupPlan#disk_backup_plan_properties}
	DiskBackupPlanProperties *GoogleBackupDrBackupPlanDiskBackupPlanProperties `field:"optional" json:"diskBackupPlanProperties" yaml:"diskBackupPlanProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#id GoogleBackupDrBackupPlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This is only applicable for CloudSql resource.
	//
	// Days for which logs will be stored. This value should be greater than or equal to minimum enforced log retention duration of the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#log_retention_days GoogleBackupDrBackupPlan#log_retention_days}
	LogRetentionDays *float64 `field:"optional" json:"logRetentionDays" yaml:"logRetentionDays"`
	// The maximum number of days for which an on-demand backup taken with custom retention can be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#max_custom_on_demand_retention_days GoogleBackupDrBackupPlan#max_custom_on_demand_retention_days}
	MaxCustomOnDemandRetentionDays *float64 `field:"optional" json:"maxCustomOnDemandRetentionDays" yaml:"maxCustomOnDemandRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#project GoogleBackupDrBackupPlan#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_plan#timeouts GoogleBackupDrBackupPlan#timeouts}
	Timeouts *GoogleBackupDrBackupPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

