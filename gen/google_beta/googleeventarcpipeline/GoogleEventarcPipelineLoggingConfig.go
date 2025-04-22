package googleeventarcpipeline


type GoogleEventarcPipelineLoggingConfig struct {
	// The minimum severity of logs that will be sent to Stackdriver/Platform Telemetry.
	//
	// Logs at severitiy ≥ this value will be sent, unless it is NONE. Possible values: ["NONE", "DEBUG", "INFO", "NOTICE", "WARNING", "ERROR", "CRITICAL", "ALERT", "EMERGENCY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#log_severity GoogleEventarcPipeline#log_severity}
	LogSeverity *string `field:"optional" json:"logSeverity" yaml:"logSeverity"`
}

