package googledataprocautoscalingpolicyiambinding


type GoogleDataprocAutoscalingPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy_iam_binding#expression GoogleDataprocAutoscalingPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy_iam_binding#title GoogleDataprocAutoscalingPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy_iam_binding#description GoogleDataprocAutoscalingPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

