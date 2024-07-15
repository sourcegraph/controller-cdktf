package googlecloudfunctionsfunction


type GoogleCloudfunctionsFunctionSourceRepository struct {
	// The URL pointing to the hosted repository where the function is defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions_function#url GoogleCloudfunctionsFunction#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

