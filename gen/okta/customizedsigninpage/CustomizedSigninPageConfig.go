package customizedsigninpage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomizedSigninPageConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// brand id of the preview signin page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#brand_id CustomizedSigninPage#brand_id}
	BrandId *string `field:"required" json:"brandId" yaml:"brandId"`
	// page content of the preview signin page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#page_content CustomizedSigninPage#page_content}
	PageContent *string `field:"required" json:"pageContent" yaml:"pageContent"`
	// widget version specified as a Semver.
	//
	// The following are currently supported
	// 			*, ^1, ^2, ^3, ^4, ^5, ^6, ^7, 1.6, 1.7, 1.8, 1.9, 1.10, 1.11, 1.12, 1.13, 2.1, 2.2, 2.3, 2.4,
	// 			2.5, 2.6, 2.7, 2.8, 2.9, 2.10, 2.11, 2.12, 2.13, 2.14, 2.15, 2.16, 2.17, 2.18, 2.19, 2.20, 2.21,
	// 			3.0, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 4.0, 4.1, 4.2, 4.3, 4.4, 4.5, 5.0, 5.1, 5.2, 5.3,
	// 			5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 5.10, 5.11, 5.12, 5.13, 5.14, 5.15, 5.16, 6.0, 6.1, 6.2, 6.3, 6.4, 6.5,
	// 			6.6, 6.7, 6.8, 6.9, 7.0, 7.1, 7.2, 7.3, 7.4, 7.5, 7.6, 7.7, 7.8, 7.9, 7.10, 7.11, 7.12, 7.13.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#widget_version CustomizedSigninPage#widget_version}
	WidgetVersion *string `field:"required" json:"widgetVersion" yaml:"widgetVersion"`
	// content_security_policy_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#content_security_policy_setting CustomizedSigninPage#content_security_policy_setting}
	ContentSecurityPolicySetting *CustomizedSigninPageContentSecurityPolicySetting `field:"optional" json:"contentSecurityPolicySetting" yaml:"contentSecurityPolicySetting"`
	// widget_customizations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#widget_customizations CustomizedSigninPage#widget_customizations}
	WidgetCustomizations *CustomizedSigninPageWidgetCustomizations `field:"optional" json:"widgetCustomizations" yaml:"widgetCustomizations"`
}

