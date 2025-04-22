package googlememorystoreinstance


type GoogleMemorystoreInstanceAutomatedBackupConfig struct {
	// fixed_frequency_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#fixed_frequency_schedule GoogleMemorystoreInstance#fixed_frequency_schedule}
	FixedFrequencySchedule *GoogleMemorystoreInstanceAutomatedBackupConfigFixedFrequencySchedule `field:"required" json:"fixedFrequencySchedule" yaml:"fixedFrequencySchedule"`
	// How long to keep automated backups before the backups are deleted.
	//
	// The value should be between 1 day and 365 days. If not specified, the default value is 35 days.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". The default_value is "3024000s"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#retention GoogleMemorystoreInstance#retention}
	Retention *string `field:"required" json:"retention" yaml:"retention"`
}

