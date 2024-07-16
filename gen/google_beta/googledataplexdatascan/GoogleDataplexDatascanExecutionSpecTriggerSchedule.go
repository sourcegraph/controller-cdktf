package googledataplexdatascan


type GoogleDataplexDatascanExecutionSpecTriggerSchedule struct {
	// Cron schedule for running scans periodically. This field is required for Schedule scans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#cron GoogleDataplexDatascan#cron}
	Cron *string `field:"required" json:"cron" yaml:"cron"`
}

