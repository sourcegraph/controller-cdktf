package schedulerschedule


type SchedulerScheduleFlexibleTimeWindow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#mode SchedulerSchedule#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#maximum_window_in_minutes SchedulerSchedule#maximum_window_in_minutes}.
	MaximumWindowInMinutes *float64 `field:"optional" json:"maximumWindowInMinutes" yaml:"maximumWindowInMinutes"`
}
