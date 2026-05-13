package bigquerydatapolicyv2datapolicyiambinding


type BigqueryDatapolicyv2DataPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_datapolicyv2_data_policy_iam_binding#expression BigqueryDatapolicyv2DataPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_datapolicyv2_data_policy_iam_binding#title BigqueryDatapolicyv2DataPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_datapolicyv2_data_policy_iam_binding#description BigqueryDatapolicyv2DataPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

