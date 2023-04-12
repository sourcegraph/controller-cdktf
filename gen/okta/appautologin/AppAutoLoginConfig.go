package appautologin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppAutoLoginConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#label AppAutoLogin#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#accessibility_error_redirect_url AppAutoLogin#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#accessibility_login_redirect_url AppAutoLogin#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#accessibility_self_service AppAutoLogin#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#admin_note AppAutoLogin#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#app_links_json AppAutoLogin#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Application settings in JSON format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#app_settings_json AppAutoLogin#app_settings_json}
	AppSettingsJson *string `field:"optional" json:"appSettingsJson" yaml:"appSettingsJson"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#auto_submit_toolbar AppAutoLogin#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Application credentials scheme.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#credentials_scheme AppAutoLogin#credentials_scheme}
	CredentialsScheme *string `field:"optional" json:"credentialsScheme" yaml:"credentialsScheme"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#enduser_note AppAutoLogin#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Groups associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#groups AppAutoLogin#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#hide_ios AppAutoLogin#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#hide_web AppAutoLogin#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#id AppAutoLogin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Local path to logo of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#logo AppAutoLogin#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Preconfigured app name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#preconfigured_app AppAutoLogin#preconfigured_app}
	PreconfiguredApp *string `field:"optional" json:"preconfiguredApp" yaml:"preconfiguredApp"`
	// Allow user to reveal password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#reveal_password AppAutoLogin#reveal_password}
	RevealPassword interface{} `field:"optional" json:"revealPassword" yaml:"revealPassword"`
	// Shared password, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#shared_password AppAutoLogin#shared_password}
	SharedPassword *string `field:"optional" json:"sharedPassword" yaml:"sharedPassword"`
	// Shared username, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#shared_username AppAutoLogin#shared_username}
	SharedUsername *string `field:"optional" json:"sharedUsername" yaml:"sharedUsername"`
	// Post login redirect URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#sign_on_redirect_url AppAutoLogin#sign_on_redirect_url}
	SignOnRedirectUrl *string `field:"optional" json:"signOnRedirectUrl" yaml:"signOnRedirectUrl"`
	// Login URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#sign_on_url AppAutoLogin#sign_on_url}
	SignOnUrl *string `field:"optional" json:"signOnUrl" yaml:"signOnUrl"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#skip_groups AppAutoLogin#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#skip_users AppAutoLogin#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// Status of application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#status AppAutoLogin#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#timeouts AppAutoLogin#timeouts}
	Timeouts *AppAutoLoginTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Username template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#user_name_template AppAutoLogin#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#user_name_template_push_status AppAutoLogin#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#user_name_template_suffix AppAutoLogin#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#user_name_template_type AppAutoLogin#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_auto_login#users AppAutoLogin#users}
	Users interface{} `field:"optional" json:"users" yaml:"users"`
}

