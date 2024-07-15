package googlebigquerydatapolicydatapolicyiambinding


type GoogleBigqueryDatapolicyDataPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_datapolicy_data_policy_iam_binding#expression GoogleBigqueryDatapolicyDataPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_datapolicy_data_policy_iam_binding#title GoogleBigqueryDatapolicyDataPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_datapolicy_data_policy_iam_binding#description GoogleBigqueryDatapolicyDataPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

