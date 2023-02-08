package googledataprocmetastorefederationiambinding


type GoogleDataprocMetastoreFederationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation_iam_binding#expression GoogleDataprocMetastoreFederationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation_iam_binding#title GoogleDataprocMetastoreFederationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_federation_iam_binding#description GoogleDataprocMetastoreFederationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

