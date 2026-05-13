package googlebigquerydatapolicyv2datapolicyiambinding


type GoogleBigqueryDatapolicyv2DataPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_binding#expression GoogleBigqueryDatapolicyv2DataPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_binding#title GoogleBigqueryDatapolicyv2DataPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_binding#description GoogleBigqueryDatapolicyv2DataPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

