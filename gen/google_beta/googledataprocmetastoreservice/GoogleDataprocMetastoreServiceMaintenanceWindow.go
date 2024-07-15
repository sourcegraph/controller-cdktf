package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceMaintenanceWindow struct {
	// The day of week, when the window starts. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#day_of_week GoogleDataprocMetastoreService#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The hour of day (0-23) when the window starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#hour_of_day GoogleDataprocMetastoreService#hour_of_day}
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
}

