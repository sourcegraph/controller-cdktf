package backupdrrestoreworkload

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupDrRestoreWorkloadConfig struct {
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
	// Required. The ID of the backup to restore from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#backup_id BackupDrRestoreWorkload#backup_id}
	BackupId *string `field:"required" json:"backupId" yaml:"backupId"`
	// Required. The ID of the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#backup_vault_id BackupDrRestoreWorkload#backup_vault_id}
	BackupVaultId *string `field:"required" json:"backupVaultId" yaml:"backupVaultId"`
	// Required. The ID of the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#data_source_id BackupDrRestoreWorkload#data_source_id}
	DataSourceId *string `field:"required" json:"dataSourceId" yaml:"dataSourceId"`
	// Required. The location for the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#location BackupDrRestoreWorkload#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Optional. A field mask used to clear server-side default values during restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#clear_overrides_field_mask BackupDrRestoreWorkload#clear_overrides_field_mask}
	ClearOverridesFieldMask *string `field:"optional" json:"clearOverridesFieldMask" yaml:"clearOverridesFieldMask"`
	// compute_instance_restore_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#compute_instance_restore_properties BackupDrRestoreWorkload#compute_instance_restore_properties}
	ComputeInstanceRestoreProperties *BackupDrRestoreWorkloadComputeInstanceRestoreProperties `field:"optional" json:"computeInstanceRestoreProperties" yaml:"computeInstanceRestoreProperties"`
	// compute_instance_target_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#compute_instance_target_environment BackupDrRestoreWorkload#compute_instance_target_environment}
	ComputeInstanceTargetEnvironment *BackupDrRestoreWorkloadComputeInstanceTargetEnvironment `field:"optional" json:"computeInstanceTargetEnvironment" yaml:"computeInstanceTargetEnvironment"`
	// Optional.
	//
	// If true (default), running terraform destroy will delete the live resource in GCP.
	// If false, only the restore record is removed from the state, leaving the resource active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#delete_restored_instance BackupDrRestoreWorkload#delete_restored_instance}
	DeleteRestoredInstance interface{} `field:"optional" json:"deleteRestoredInstance" yaml:"deleteRestoredInstance"`
	// disk_restore_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#disk_restore_properties BackupDrRestoreWorkload#disk_restore_properties}
	DiskRestoreProperties *BackupDrRestoreWorkloadDiskRestoreProperties `field:"optional" json:"diskRestoreProperties" yaml:"diskRestoreProperties"`
	// disk_target_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#disk_target_environment BackupDrRestoreWorkload#disk_target_environment}
	DiskTargetEnvironment *BackupDrRestoreWorkloadDiskTargetEnvironment `field:"optional" json:"diskTargetEnvironment" yaml:"diskTargetEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#id BackupDrRestoreWorkload#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource name of the backup instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#name BackupDrRestoreWorkload#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// region_disk_target_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#region_disk_target_environment BackupDrRestoreWorkload#region_disk_target_environment}
	RegionDiskTargetEnvironment *BackupDrRestoreWorkloadRegionDiskTargetEnvironment `field:"optional" json:"regionDiskTargetEnvironment" yaml:"regionDiskTargetEnvironment"`
	// Optional.
	//
	// An optional request ID to identify requests. Specify a unique request ID
	// so that if you must retry your request, the server will know to ignore
	// the request if it has already been completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#request_id BackupDrRestoreWorkload#request_id}
	RequestId *string `field:"optional" json:"requestId" yaml:"requestId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#timeouts BackupDrRestoreWorkload#timeouts}
	Timeouts *BackupDrRestoreWorkloadTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

