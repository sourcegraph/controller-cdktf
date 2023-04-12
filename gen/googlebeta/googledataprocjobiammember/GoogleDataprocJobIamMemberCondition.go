package googledataprocjobiammember


type GoogleDataprocJobIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job_iam_member#expression GoogleDataprocJobIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job_iam_member#title GoogleDataprocJobIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job_iam_member#description GoogleDataprocJobIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

