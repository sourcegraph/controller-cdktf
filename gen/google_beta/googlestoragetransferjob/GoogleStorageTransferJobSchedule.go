package googlestoragetransferjob


type GoogleStorageTransferJobSchedule struct {
	// schedule_start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#schedule_start_date GoogleStorageTransferJob#schedule_start_date}
	ScheduleStartDate *GoogleStorageTransferJobScheduleScheduleStartDate `field:"required" json:"scheduleStartDate" yaml:"scheduleStartDate"`
	// Interval between the start of each scheduled transfer.
	//
	// If unspecified, the default value is 24 hours. This value may not be less than 1 hour. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#repeat_interval GoogleStorageTransferJob#repeat_interval}
	RepeatInterval *string `field:"optional" json:"repeatInterval" yaml:"repeatInterval"`
	// schedule_end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#schedule_end_date GoogleStorageTransferJob#schedule_end_date}
	ScheduleEndDate *GoogleStorageTransferJobScheduleScheduleEndDate `field:"optional" json:"scheduleEndDate" yaml:"scheduleEndDate"`
	// start_time_of_day block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#start_time_of_day GoogleStorageTransferJob#start_time_of_day}
	StartTimeOfDay *GoogleStorageTransferJobScheduleStartTimeOfDay `field:"optional" json:"startTimeOfDay" yaml:"startTimeOfDay"`
}

