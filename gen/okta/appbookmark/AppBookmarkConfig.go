package appbookmark

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppBookmarkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Pretty name of app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#label AppBookmark#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#url AppBookmark#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#accessibility_error_redirect_url AppBookmark#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#accessibility_login_redirect_url AppBookmark#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#accessibility_self_service AppBookmark#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#admin_note AppBookmark#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#app_links_json AppBookmark#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Id of this apps authentication policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#authentication_policy AppBookmark#authentication_policy}
	AuthenticationPolicy *string `field:"optional" json:"authenticationPolicy" yaml:"authenticationPolicy"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#auto_submit_toolbar AppBookmark#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#enduser_note AppBookmark#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Groups associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#groups AppBookmark#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#hide_ios AppBookmark#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#hide_web AppBookmark#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#id AppBookmark#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Local path to logo of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#logo AppBookmark#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#request_integration AppBookmark#request_integration}.
	RequestIntegration interface{} `field:"optional" json:"requestIntegration" yaml:"requestIntegration"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#skip_groups AppBookmark#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#skip_users AppBookmark#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// Status of application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#status AppBookmark#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#timeouts AppBookmark#timeouts}
	Timeouts *AppBookmarkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#users AppBookmark#users}
	Users interface{} `field:"optional" json:"users" yaml:"users"`
}

