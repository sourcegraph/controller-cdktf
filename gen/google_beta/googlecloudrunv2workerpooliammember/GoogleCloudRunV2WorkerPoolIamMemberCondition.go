package googlecloudrunv2workerpooliammember


type GoogleCloudRunV2WorkerPoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool_iam_member#expression GoogleCloudRunV2WorkerPoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool_iam_member#title GoogleCloudRunV2WorkerPoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool_iam_member#description GoogleCloudRunV2WorkerPoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

