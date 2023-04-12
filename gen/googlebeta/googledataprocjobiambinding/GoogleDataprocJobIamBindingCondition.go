package googledataprocjobiambinding


type GoogleDataprocJobIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job_iam_binding#expression GoogleDataprocJobIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job_iam_binding#title GoogleDataprocJobIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job_iam_binding#description GoogleDataprocJobIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

