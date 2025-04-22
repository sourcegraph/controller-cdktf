package googlemanagedkafkaconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleManagedKafkaConnectorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The connect cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#connect_cluster GoogleManagedKafkaConnector#connect_cluster}
	ConnectCluster *string `field:"required" json:"connectCluster" yaml:"connectCluster"`
	// The ID to use for the connector, which will become the final component of the connector's name.
	//
	// This value is structured like: 'my-connector-id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#connector_id GoogleManagedKafkaConnector#connector_id}
	ConnectorId *string `field:"required" json:"connectorId" yaml:"connectorId"`
	// ID of the location of the Kafka Connect resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#location GoogleManagedKafkaConnector#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Connector config as keys/values. The keys of the map are connector property names, for example: 'connector.class', 'tasks.max', 'key.converter'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#configs GoogleManagedKafkaConnector#configs}
	Configs *map[string]*string `field:"optional" json:"configs" yaml:"configs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#id GoogleManagedKafkaConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#project GoogleManagedKafkaConnector#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// task_restart_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#task_restart_policy GoogleManagedKafkaConnector#task_restart_policy}
	TaskRestartPolicy *GoogleManagedKafkaConnectorTaskRestartPolicy `field:"optional" json:"taskRestartPolicy" yaml:"taskRestartPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connector#timeouts GoogleManagedKafkaConnector#timeouts}
	Timeouts *GoogleManagedKafkaConnectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

