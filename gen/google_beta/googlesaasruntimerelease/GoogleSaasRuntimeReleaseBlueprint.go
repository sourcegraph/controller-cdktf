package googlesaasruntimerelease


type GoogleSaasRuntimeReleaseBlueprint struct {
	// URI to a blueprint used by the Unit (required unless unitKind or release is set).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#package GoogleSaasRuntimeRelease#package}
	Package *string `field:"optional" json:"package" yaml:"package"`
}

