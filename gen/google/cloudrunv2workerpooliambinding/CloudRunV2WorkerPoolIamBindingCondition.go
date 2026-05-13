package cloudrunv2workerpooliambinding


type CloudRunV2WorkerPoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool_iam_binding#expression CloudRunV2WorkerPoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool_iam_binding#title CloudRunV2WorkerPoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_run_v2_worker_pool_iam_binding#description CloudRunV2WorkerPoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

