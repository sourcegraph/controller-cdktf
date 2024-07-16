package googlerecaptchaenterprisekey


type GoogleRecaptchaEnterpriseKeyAndroidSettings struct {
	// If set to true, it means allowed_package_names will not be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#allow_all_package_names GoogleRecaptchaEnterpriseKey#allow_all_package_names}
	AllowAllPackageNames interface{} `field:"optional" json:"allowAllPackageNames" yaml:"allowAllPackageNames"`
	// Android package names of apps allowed to use the key. Example: 'com.companyname.appname'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_recaptcha_enterprise_key#allowed_package_names GoogleRecaptchaEnterpriseKey#allowed_package_names}
	AllowedPackageNames *[]*string `field:"optional" json:"allowedPackageNames" yaml:"allowedPackageNames"`
}

