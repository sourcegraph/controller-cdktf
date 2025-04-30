package googleiamworkloadidentitypoolprovider


type GoogleIamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchors struct {
	// PEM certificate of the PKI used for validation. Must only contain one ca certificate(either root or intermediate cert).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool_provider#pem_certificate GoogleIamWorkloadIdentityPoolProvider#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
}

