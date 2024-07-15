package googlebigquerydatapolicydatapolicyiammember


type GoogleBigqueryDatapolicyDataPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_datapolicy_data_policy_iam_member#expression GoogleBigqueryDatapolicyDataPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_datapolicy_data_policy_iam_member#title GoogleBigqueryDatapolicyDataPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_datapolicy_data_policy_iam_member#description GoogleBigqueryDatapolicyDataPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

