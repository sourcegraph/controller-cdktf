package googlespannerbackupschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSpannerBackupScheduleConfig struct {
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
	// The database to create the backup schedule on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#database GoogleSpannerBackupSchedule#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The instance to create the database on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#instance GoogleSpannerBackupSchedule#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: '3.5s'. You can set this to a value up to 366 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#retention_duration GoogleSpannerBackupSchedule#retention_duration}
	RetentionDuration *string `field:"required" json:"retentionDuration" yaml:"retentionDuration"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#encryption_config GoogleSpannerBackupSchedule#encryption_config}
	EncryptionConfig *GoogleSpannerBackupScheduleEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// full_backup_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#full_backup_spec GoogleSpannerBackupSchedule#full_backup_spec}
	FullBackupSpec *GoogleSpannerBackupScheduleFullBackupSpec `field:"optional" json:"fullBackupSpec" yaml:"fullBackupSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#id GoogleSpannerBackupSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// incremental_backup_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#incremental_backup_spec GoogleSpannerBackupSchedule#incremental_backup_spec}
	IncrementalBackupSpec *GoogleSpannerBackupScheduleIncrementalBackupSpec `field:"optional" json:"incrementalBackupSpec" yaml:"incrementalBackupSpec"`
	// A unique identifier for the backup schedule, which cannot be changed after the backup schedule is created.
	//
	// Values are of the form [a-z][-a-z0-9]*[a-z0-9].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#name GoogleSpannerBackupSchedule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#project GoogleSpannerBackupSchedule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#spec GoogleSpannerBackupSchedule#spec}
	Spec *GoogleSpannerBackupScheduleSpec `field:"optional" json:"spec" yaml:"spec"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#timeouts GoogleSpannerBackupSchedule#timeouts}
	Timeouts *GoogleSpannerBackupScheduleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

