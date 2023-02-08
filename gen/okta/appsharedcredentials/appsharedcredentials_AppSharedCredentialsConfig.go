package appsharedcredentials

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSharedCredentialsConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#label AppSharedCredentials#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#accessibility_error_redirect_url AppSharedCredentials#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#accessibility_login_redirect_url AppSharedCredentials#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#accessibility_self_service AppSharedCredentials#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#admin_note AppSharedCredentials#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#app_links_json AppSharedCredentials#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#auto_submit_toolbar AppSharedCredentials#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Login button field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#button_field AppSharedCredentials#button_field}
	ButtonField *string `field:"optional" json:"buttonField" yaml:"buttonField"`
	// CSS selector for the checkbox.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#checkbox AppSharedCredentials#checkbox}
	Checkbox *string `field:"optional" json:"checkbox" yaml:"checkbox"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#enduser_note AppSharedCredentials#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Groups associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#groups AppSharedCredentials#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#hide_ios AppSharedCredentials#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#hide_web AppSharedCredentials#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#id AppSharedCredentials#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Local path to logo of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#logo AppSharedCredentials#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Login password field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#password_field AppSharedCredentials#password_field}
	PasswordField *string `field:"optional" json:"passwordField" yaml:"passwordField"`
	// Preconfigured app name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#preconfigured_app AppSharedCredentials#preconfigured_app}
	PreconfiguredApp *string `field:"optional" json:"preconfiguredApp" yaml:"preconfiguredApp"`
	// Secondary URL of the sign-in page for this app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#redirect_url AppSharedCredentials#redirect_url}
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// Shared password, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#shared_password AppSharedCredentials#shared_password}
	SharedPassword *string `field:"optional" json:"sharedPassword" yaml:"sharedPassword"`
	// Shared username, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#shared_username AppSharedCredentials#shared_username}
	SharedUsername *string `field:"optional" json:"sharedUsername" yaml:"sharedUsername"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#skip_groups AppSharedCredentials#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#skip_users AppSharedCredentials#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// Status of application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#status AppSharedCredentials#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#timeouts AppSharedCredentials#timeouts}
	Timeouts *AppSharedCredentialsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Login URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#url AppSharedCredentials#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// A regex that further restricts URL to the specified regex.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#url_regex AppSharedCredentials#url_regex}
	UrlRegex *string `field:"optional" json:"urlRegex" yaml:"urlRegex"`
	// Login username field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#username_field AppSharedCredentials#username_field}
	UsernameField *string `field:"optional" json:"usernameField" yaml:"usernameField"`
	// Username template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#user_name_template AppSharedCredentials#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#user_name_template_push_status AppSharedCredentials#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#user_name_template_suffix AppSharedCredentials#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#user_name_template_type AppSharedCredentials#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_shared_credentials#users AppSharedCredentials#users}
	Users interface{} `field:"optional" json:"users" yaml:"users"`
}

