package googledataprocjob


type GoogleDataprocJobSparksqlConfigLoggingConfig struct {
	// Optional.
	//
	// The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job#driver_log_levels GoogleDataprocJob#driver_log_levels}
	DriverLogLevels *map[string]*string `field:"required" json:"driverLogLevels" yaml:"driverLogLevels"`
}

