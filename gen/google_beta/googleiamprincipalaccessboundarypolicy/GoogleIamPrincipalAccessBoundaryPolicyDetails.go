package googleiamprincipalaccessboundarypolicy


type GoogleIamPrincipalAccessBoundaryPolicyDetails struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iam_principal_access_boundary_policy#rules GoogleIamPrincipalAccessBoundaryPolicy#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The version number that indicates which Google Cloud services are included in the enforcement (e.g. \"latest\", \"1\", ...). If empty, the PAB policy version will be set to the current latest version, and this version won't get updated when new versions are released.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iam_principal_access_boundary_policy#enforcement_version GoogleIamPrincipalAccessBoundaryPolicy#enforcement_version}
	EnforcementVersion *string `field:"optional" json:"enforcementVersion" yaml:"enforcementVersion"`
}

