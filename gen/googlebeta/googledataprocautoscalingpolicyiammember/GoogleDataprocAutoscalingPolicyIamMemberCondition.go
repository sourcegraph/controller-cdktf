package googledataprocautoscalingpolicyiammember


type GoogleDataprocAutoscalingPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_autoscaling_policy_iam_member#expression GoogleDataprocAutoscalingPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_autoscaling_policy_iam_member#title GoogleDataprocAutoscalingPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_autoscaling_policy_iam_member#description GoogleDataprocAutoscalingPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

