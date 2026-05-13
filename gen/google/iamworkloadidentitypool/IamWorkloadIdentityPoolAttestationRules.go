package iamworkloadidentitypool


type IamWorkloadIdentityPoolAttestationRules struct {
	// A single workload operating on Google Cloud. For example: '//run.googleapis.com/projects/123/type/Service/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#google_cloud_resource IamWorkloadIdentityPool#google_cloud_resource}
	GoogleCloudResource *string `field:"required" json:"googleCloudResource" yaml:"googleCloudResource"`
}

