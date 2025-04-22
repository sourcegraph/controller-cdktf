package googledataprocmetastoretableiambinding


type GoogleDataprocMetastoreTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_metastore_table_iam_binding#expression GoogleDataprocMetastoreTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_metastore_table_iam_binding#title GoogleDataprocMetastoreTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_metastore_table_iam_binding#description GoogleDataprocMetastoreTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

