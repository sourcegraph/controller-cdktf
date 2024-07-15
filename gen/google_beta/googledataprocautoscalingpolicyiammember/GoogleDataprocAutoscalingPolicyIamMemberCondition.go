package googledataprocautoscalingpolicyiammember


type GoogleDataprocAutoscalingPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy_iam_member#expression GoogleDataprocAutoscalingPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy_iam_member#title GoogleDataprocAutoscalingPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_autoscaling_policy_iam_member#description GoogleDataprocAutoscalingPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

