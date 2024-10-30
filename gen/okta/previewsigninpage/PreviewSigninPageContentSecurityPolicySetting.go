package previewsigninpage


type PreviewSigninPageContentSecurityPolicySetting struct {
	// enforced or report_only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/preview_signin_page#mode PreviewSigninPage#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/preview_signin_page#report_uri PreviewSigninPage#report_uri}.
	ReportUri *string `field:"optional" json:"reportUri" yaml:"reportUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/preview_signin_page#src_list PreviewSigninPage#src_list}.
	SrcList *[]*string `field:"optional" json:"srcList" yaml:"srcList"`
}

