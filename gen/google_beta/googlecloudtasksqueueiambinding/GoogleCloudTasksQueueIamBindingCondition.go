package googlecloudtasksqueueiambinding


type GoogleCloudTasksQueueIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue_iam_binding#expression GoogleCloudTasksQueueIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue_iam_binding#title GoogleCloudTasksQueueIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue_iam_binding#description GoogleCloudTasksQueueIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

