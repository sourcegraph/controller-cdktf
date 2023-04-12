package googledataprocclusteriambinding


type GoogleDataprocClusterIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster_iam_binding#expression GoogleDataprocClusterIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster_iam_binding#title GoogleDataprocClusterIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster_iam_binding#description GoogleDataprocClusterIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

