package googlefirestorebackupschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirestoreBackupScheduleConfig struct {
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
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// You can set this to a value up to 14 weeks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#retention GoogleFirestoreBackupSchedule#retention}
	Retention *string `field:"required" json:"retention" yaml:"retention"`
	// daily_recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#daily_recurrence GoogleFirestoreBackupSchedule#daily_recurrence}
	DailyRecurrence *GoogleFirestoreBackupScheduleDailyRecurrence `field:"optional" json:"dailyRecurrence" yaml:"dailyRecurrence"`
	// The Firestore database id. Defaults to '"(default)"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#database GoogleFirestoreBackupSchedule#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#id GoogleFirestoreBackupSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#project GoogleFirestoreBackupSchedule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#timeouts GoogleFirestoreBackupSchedule#timeouts}
	Timeouts *GoogleFirestoreBackupScheduleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// weekly_recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_backup_schedule#weekly_recurrence GoogleFirestoreBackupSchedule#weekly_recurrence}
	WeeklyRecurrence *GoogleFirestoreBackupScheduleWeeklyRecurrence `field:"optional" json:"weeklyRecurrence" yaml:"weeklyRecurrence"`
}

