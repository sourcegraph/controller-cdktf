package googledataprocjob


type GoogleDataprocJobScheduling struct {
	// Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#max_failures_per_hour GoogleDataprocJob#max_failures_per_hour}
	MaxFailuresPerHour *float64 `field:"required" json:"maxFailuresPerHour" yaml:"maxFailuresPerHour"`
	// Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#max_failures_total GoogleDataprocJob#max_failures_total}
	MaxFailuresTotal *float64 `field:"required" json:"maxFailuresTotal" yaml:"maxFailuresTotal"`
}

