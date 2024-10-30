package customizedsigninpage


type CustomizedSigninPageContentSecurityPolicySetting struct {
	// enforced or report_only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#mode CustomizedSigninPage#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#report_uri CustomizedSigninPage#report_uri}.
	ReportUri *string `field:"optional" json:"reportUri" yaml:"reportUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#src_list CustomizedSigninPage#src_list}.
	SrcList *[]*string `field:"optional" json:"srcList" yaml:"srcList"`
}

