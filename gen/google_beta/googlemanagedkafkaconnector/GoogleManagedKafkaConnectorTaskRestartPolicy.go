package googlemanagedkafkaconnector


type GoogleManagedKafkaConnectorTaskRestartPolicy struct {
	// The maximum amount of time to wait before retrying a failed task.
	//
	// This sets an upper bound for the backoff delay.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#maximum_backoff GoogleManagedKafkaConnector#maximum_backoff}
	MaximumBackoff *string `field:"optional" json:"maximumBackoff" yaml:"maximumBackoff"`
	// The minimum amount of time to wait before retrying a failed task.
	//
	// This sets a lower bound for the backoff delay.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#minimum_backoff GoogleManagedKafkaConnector#minimum_backoff}
	MinimumBackoff *string `field:"optional" json:"minimumBackoff" yaml:"minimumBackoff"`
}

