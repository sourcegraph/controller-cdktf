package googlecloudtasksqueue


type GoogleCloudTasksQueueHttpTargetHeaderOverridesHeader struct {
	// The Key of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_tasks_queue#key GoogleCloudTasksQueue#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Value of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_tasks_queue#value GoogleCloudTasksQueue#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

