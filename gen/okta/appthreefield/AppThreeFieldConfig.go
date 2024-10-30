package appthreefield

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppThreeFieldConfig struct {
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
	// Login button field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#button_selector AppThreeField#button_selector}
	ButtonSelector *string `field:"required" json:"buttonSelector" yaml:"buttonSelector"`
	// Extra field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#extra_field_selector AppThreeField#extra_field_selector}
	ExtraFieldSelector *string `field:"required" json:"extraFieldSelector" yaml:"extraFieldSelector"`
	// Value for extra form field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#extra_field_value AppThreeField#extra_field_value}
	ExtraFieldValue *string `field:"required" json:"extraFieldValue" yaml:"extraFieldValue"`
	// The Application's display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#label AppThreeField#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Login password field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#password_selector AppThreeField#password_selector}
	PasswordSelector *string `field:"required" json:"passwordSelector" yaml:"passwordSelector"`
	// Login URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#url AppThreeField#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Login username field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#username_selector AppThreeField#username_selector}
	UsernameSelector *string `field:"required" json:"usernameSelector" yaml:"usernameSelector"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#accessibility_error_redirect_url AppThreeField#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#accessibility_login_redirect_url AppThreeField#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service. Default is `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#accessibility_self_service AppThreeField#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#admin_note AppThreeField#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app. The value for each application link should be boolean.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#app_links_json AppThreeField#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#auto_submit_toolbar AppThreeField#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Application credentials scheme. One of: `EDIT_USERNAME_AND_PASSWORD`, `ADMIN_SETS_CREDENTIALS`, `EDIT_PASSWORD_ONLY`, `EXTERNAL_PASSWORD_SYNC`, or `SHARED_USERNAME_AND_PASSWORD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#credentials_scheme AppThreeField#credentials_scheme}
	CredentialsScheme *string `field:"optional" json:"credentialsScheme" yaml:"credentialsScheme"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#enduser_note AppThreeField#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#hide_ios AppThreeField#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#hide_web AppThreeField#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#id AppThreeField#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Local file path to the logo.
	//
	// The file must be in PNG, JPG, or GIF format, and less than 1 MB in size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#logo AppThreeField#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Allow user to reveal password. It can not be set to `true` if `credentials_scheme` is `ADMIN_SETS_CREDENTIALS`, `SHARED_USERNAME_AND_PASSWORD` or `EXTERNAL_PASSWORD_SYNC`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#reveal_password AppThreeField#reveal_password}
	RevealPassword interface{} `field:"optional" json:"revealPassword" yaml:"revealPassword"`
	// Shared password, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#shared_password AppThreeField#shared_password}
	SharedPassword *string `field:"optional" json:"sharedPassword" yaml:"sharedPassword"`
	// Shared username, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#shared_username AppThreeField#shared_username}
	SharedUsername *string `field:"optional" json:"sharedUsername" yaml:"sharedUsername"`
	// Status of application. By default, it is `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#status AppThreeField#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#timeouts AppThreeField#timeouts}
	Timeouts *AppThreeFieldTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A regex that further restricts URL to the specified regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#url_regex AppThreeField#url_regex}
	UrlRegex *string `field:"optional" json:"urlRegex" yaml:"urlRegex"`
	// Username template. Default: `${source.login}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#user_name_template AppThreeField#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update. Valid values: `PUSH` and `DONT_PUSH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#user_name_template_push_status AppThreeField#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#user_name_template_suffix AppThreeField#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type. Default: `BUILT_IN`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_three_field#user_name_template_type AppThreeField#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
}

