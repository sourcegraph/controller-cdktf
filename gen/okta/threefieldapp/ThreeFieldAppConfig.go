package threefieldapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ThreeFieldAppConfig struct {
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
	// Login button field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#button_selector ThreeFieldApp#button_selector}
	ButtonSelector *string `field:"required" json:"buttonSelector" yaml:"buttonSelector"`
	// Extra field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#extra_field_selector ThreeFieldApp#extra_field_selector}
	ExtraFieldSelector *string `field:"required" json:"extraFieldSelector" yaml:"extraFieldSelector"`
	// Value for extra form field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#extra_field_value ThreeFieldApp#extra_field_value}
	ExtraFieldValue *string `field:"required" json:"extraFieldValue" yaml:"extraFieldValue"`
	// Pretty name of app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#label ThreeFieldApp#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Login password field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#password_selector ThreeFieldApp#password_selector}
	PasswordSelector *string `field:"required" json:"passwordSelector" yaml:"passwordSelector"`
	// Login URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#url ThreeFieldApp#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Login username field CSS selector.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#username_selector ThreeFieldApp#username_selector}
	UsernameSelector *string `field:"required" json:"usernameSelector" yaml:"usernameSelector"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#accessibility_error_redirect_url ThreeFieldApp#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#accessibility_login_redirect_url ThreeFieldApp#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#accessibility_self_service ThreeFieldApp#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#admin_note ThreeFieldApp#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#app_links_json ThreeFieldApp#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#auto_submit_toolbar ThreeFieldApp#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Application credentials scheme.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#credentials_scheme ThreeFieldApp#credentials_scheme}
	CredentialsScheme *string `field:"optional" json:"credentialsScheme" yaml:"credentialsScheme"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#enduser_note ThreeFieldApp#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Groups associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#groups ThreeFieldApp#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#hide_ios ThreeFieldApp#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#hide_web ThreeFieldApp#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#id ThreeFieldApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Local path to logo of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#logo ThreeFieldApp#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Allow user to reveal password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#reveal_password ThreeFieldApp#reveal_password}
	RevealPassword interface{} `field:"optional" json:"revealPassword" yaml:"revealPassword"`
	// Shared password, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#shared_password ThreeFieldApp#shared_password}
	SharedPassword *string `field:"optional" json:"sharedPassword" yaml:"sharedPassword"`
	// Shared username, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#shared_username ThreeFieldApp#shared_username}
	SharedUsername *string `field:"optional" json:"sharedUsername" yaml:"sharedUsername"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#skip_groups ThreeFieldApp#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#skip_users ThreeFieldApp#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// Status of application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#status ThreeFieldApp#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#timeouts ThreeFieldApp#timeouts}
	Timeouts *ThreeFieldAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A regex that further restricts URL to the specified regex.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#url_regex ThreeFieldApp#url_regex}
	UrlRegex *string `field:"optional" json:"urlRegex" yaml:"urlRegex"`
	// Username template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#user_name_template ThreeFieldApp#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#user_name_template_push_status ThreeFieldApp#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#user_name_template_suffix ThreeFieldApp#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#user_name_template_type ThreeFieldApp#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#users ThreeFieldApp#users}
	Users interface{} `field:"optional" json:"users" yaml:"users"`
}

