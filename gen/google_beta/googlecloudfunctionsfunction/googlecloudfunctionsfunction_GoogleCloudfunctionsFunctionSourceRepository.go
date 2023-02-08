package googlecloudfunctionsfunction


type GoogleCloudfunctionsFunctionSourceRepository struct {
	// The URL pointing to the hosted repository where the function is defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudfunctions_function#url GoogleCloudfunctionsFunction#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

