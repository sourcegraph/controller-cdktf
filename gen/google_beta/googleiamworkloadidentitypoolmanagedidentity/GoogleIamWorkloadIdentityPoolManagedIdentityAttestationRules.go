package googleiamworkloadidentitypoolmanagedidentity


type GoogleIamWorkloadIdentityPoolManagedIdentityAttestationRules struct {
	// A single workload operating on Google Cloud. For example: '//compute.googleapis.com/projects/123/uid/zones/us-central1-a/instances/12345678'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool_managed_identity#google_cloud_resource GoogleIamWorkloadIdentityPoolManagedIdentity#google_cloud_resource}
	GoogleCloudResource *string `field:"required" json:"googleCloudResource" yaml:"googleCloudResource"`
}

