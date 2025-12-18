package report


type ReportSchedule struct {
	// Every how many days, weeks, or months (based on frequency) should the report run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#every Report#every}
	Every *float64 `field:"required" json:"every" yaml:"every"`
	// The frequency of the report. This can be "Hourly", "Daily", "Weekly", or "Monthly".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#frequency Report#frequency}
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// The time of day to run this report in {HH:MM} format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#time_of_day Report#time_of_day}
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// The day of the month to run this report.
	//
	// Can be either 1 or 15.
	// This is only used if the frequency is "Monthly".
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#day_of_the_month Report#day_of_the_month}
	DayOfTheMonth *float64 `field:"optional" json:"dayOfTheMonth" yaml:"dayOfTheMonth"`
	// The day of the week to run this report.
	//
	// Can be "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", or "Sunday".
	// This is only used if the frequency is "Weekly".
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#day_of_the_week Report#day_of_the_week}
	DayOfTheWeek *string `field:"optional" json:"dayOfTheWeek" yaml:"dayOfTheWeek"`
	// The delay in minutes before the report is generated after the scheduled time.
	//
	// This is useful if the data is not yet available at the scheduled time.
	// Can be between 0 and 120 minutes. The default value is 30 minutes.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#generation_delay_minutes Report#generation_delay_minutes}
	GenerationDelayMinutes *float64 `field:"optional" json:"generationDelayMinutes" yaml:"generationDelayMinutes"`
	// The IANA timezone to run this report in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#timezone Report#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

