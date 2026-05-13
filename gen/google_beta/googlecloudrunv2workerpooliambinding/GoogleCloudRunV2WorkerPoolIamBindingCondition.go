package googlecloudrunv2workerpooliambinding


type GoogleCloudRunV2WorkerPoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool_iam_binding#expression GoogleCloudRunV2WorkerPoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool_iam_binding#title GoogleCloudRunV2WorkerPoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool_iam_binding#description GoogleCloudRunV2WorkerPoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

