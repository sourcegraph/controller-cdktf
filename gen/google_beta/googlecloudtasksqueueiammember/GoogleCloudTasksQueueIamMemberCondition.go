package googlecloudtasksqueueiammember


type GoogleCloudTasksQueueIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue_iam_member#expression GoogleCloudTasksQueueIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue_iam_member#title GoogleCloudTasksQueueIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue_iam_member#description GoogleCloudTasksQueueIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

