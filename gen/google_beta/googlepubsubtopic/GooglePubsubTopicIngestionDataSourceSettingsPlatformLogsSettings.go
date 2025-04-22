package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings struct {
	// The minimum severity level of Platform Logs that will be written.
	//
	// If unspecified,
	// no Platform Logs will be written. Default value: "SEVERITY_UNSPECIFIED" Possible values: ["SEVERITY_UNSPECIFIED", "DISABLED", "DEBUG", "INFO", "WARNING", "ERROR"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#severity GooglePubsubTopic#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
}

