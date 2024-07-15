package googlecloudschedulerjob


type GoogleCloudSchedulerJobPubsubTarget struct {
	// The full resource name for the Cloud Pub/Sub topic to which messages will be published when a job is delivered.
	//
	// ~>**NOTE:**
	// The topic name must be in the same format as required by PubSub's
	// PublishRequest.name, e.g. 'projects/my-project/topics/my-topic'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#topic_name GoogleCloudSchedulerJob#topic_name}
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#attributes GoogleCloudSchedulerJob#attributes}
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_scheduler_job#data GoogleCloudSchedulerJob#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
}

