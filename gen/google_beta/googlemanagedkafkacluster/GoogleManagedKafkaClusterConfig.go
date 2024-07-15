package googlemanagedkafkacluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleManagedKafkaClusterConfig struct {
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
	// capacity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#capacity_config GoogleManagedKafkaCluster#capacity_config}
	CapacityConfig *GoogleManagedKafkaClusterCapacityConfig `field:"required" json:"capacityConfig" yaml:"capacityConfig"`
	// The ID to use for the cluster, which will become the final component of the cluster's name.
	//
	// The ID must be 1-63 characters long, and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' to comply with RFC 1035. This value is structured like: 'my-cluster-id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#cluster_id GoogleManagedKafkaCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// gcp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#gcp_config GoogleManagedKafkaCluster#gcp_config}
	GcpConfig *GoogleManagedKafkaClusterGcpConfig `field:"required" json:"gcpConfig" yaml:"gcpConfig"`
	// ID of the location of the Apache Kafka for BigQuery resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#location GoogleManagedKafkaCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#id GoogleManagedKafkaCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of label KEY=VALUE pairs to add.
	//
	// Keys must start with a lowercase character and contain only hyphens (-), underscores ( ), lowercase characters, and numbers. Values must contain only hyphens (-), underscores ( ), lowercase characters, and numbers.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#labels GoogleManagedKafkaCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#project GoogleManagedKafkaCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rebalance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#rebalance_config GoogleManagedKafkaCluster#rebalance_config}
	RebalanceConfig *GoogleManagedKafkaClusterRebalanceConfig `field:"optional" json:"rebalanceConfig" yaml:"rebalanceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#timeouts GoogleManagedKafkaCluster#timeouts}
	Timeouts *GoogleManagedKafkaClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

