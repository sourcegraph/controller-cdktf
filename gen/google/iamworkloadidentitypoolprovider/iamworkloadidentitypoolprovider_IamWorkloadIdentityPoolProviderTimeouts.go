package iamworkloadidentitypoolprovider


type IamWorkloadIdentityPoolProviderTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_workload_identity_pool_provider#create IamWorkloadIdentityPoolProvider#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_workload_identity_pool_provider#delete IamWorkloadIdentityPoolProvider#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_workload_identity_pool_provider#update IamWorkloadIdentityPoolProvider#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
