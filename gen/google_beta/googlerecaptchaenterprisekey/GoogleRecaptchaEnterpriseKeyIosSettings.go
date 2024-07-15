package googlerecaptchaenterprisekey


type GoogleRecaptchaEnterpriseKeyIosSettings struct {
	// If set to true, it means allowed_bundle_ids will not be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#allow_all_bundle_ids GoogleRecaptchaEnterpriseKey#allow_all_bundle_ids}
	AllowAllBundleIds interface{} `field:"optional" json:"allowAllBundleIds" yaml:"allowAllBundleIds"`
	// iOS bundle ids of apps allowed to use the key. Example: 'com.companyname.productname.appname'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#allowed_bundle_ids GoogleRecaptchaEnterpriseKey#allowed_bundle_ids}
	AllowedBundleIds *[]*string `field:"optional" json:"allowedBundleIds" yaml:"allowedBundleIds"`
}

