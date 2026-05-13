package managedkafkaconnector


type ManagedKafkaConnectorTaskRestartPolicy struct {
	// The maximum amount of time to wait before retrying a failed task.
	//
	// This sets an upper bound for the backoff delay.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/managed_kafka_connector#maximum_backoff ManagedKafkaConnector#maximum_backoff}
	MaximumBackoff *string `field:"optional" json:"maximumBackoff" yaml:"maximumBackoff"`
	// The minimum amount of time to wait before retrying a failed task.
	//
	// This sets a lower bound for the backoff delay.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/managed_kafka_connector#minimum_backoff ManagedKafkaConnector#minimum_backoff}
	MinimumBackoff *string `field:"optional" json:"minimumBackoff" yaml:"minimumBackoff"`
}

