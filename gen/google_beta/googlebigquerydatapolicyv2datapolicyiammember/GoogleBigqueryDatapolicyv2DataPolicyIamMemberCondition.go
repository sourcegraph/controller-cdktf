package googlebigquerydatapolicyv2datapolicyiammember


type GoogleBigqueryDatapolicyv2DataPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#expression GoogleBigqueryDatapolicyv2DataPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#title GoogleBigqueryDatapolicyv2DataPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#description GoogleBigqueryDatapolicyv2DataPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

