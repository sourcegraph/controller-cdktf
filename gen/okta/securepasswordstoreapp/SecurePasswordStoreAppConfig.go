package securepasswordstoreapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurePasswordStoreAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#label SecurePasswordStoreApp#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Login password field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#password_field SecurePasswordStoreApp#password_field}
	PasswordField *string `field:"required" json:"passwordField" yaml:"passwordField"`
	// Login URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#url SecurePasswordStoreApp#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Login username field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#username_field SecurePasswordStoreApp#username_field}
	UsernameField *string `field:"required" json:"usernameField" yaml:"usernameField"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#accessibility_error_redirect_url SecurePasswordStoreApp#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#accessibility_login_redirect_url SecurePasswordStoreApp#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#accessibility_self_service SecurePasswordStoreApp#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#admin_note SecurePasswordStoreApp#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#app_links_json SecurePasswordStoreApp#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#auto_submit_toolbar SecurePasswordStoreApp#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Application credentials scheme.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#credentials_scheme SecurePasswordStoreApp#credentials_scheme}
	CredentialsScheme *string `field:"optional" json:"credentialsScheme" yaml:"credentialsScheme"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#enduser_note SecurePasswordStoreApp#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Groups associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#groups SecurePasswordStoreApp#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#hide_ios SecurePasswordStoreApp#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#hide_web SecurePasswordStoreApp#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#id SecurePasswordStoreApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Local path to logo of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#logo SecurePasswordStoreApp#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Name of optional param in the login form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#optional_field1 SecurePasswordStoreApp#optional_field1}
	OptionalField1 *string `field:"optional" json:"optionalField1" yaml:"optionalField1"`
	// Name of optional value in login form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#optional_field1_value SecurePasswordStoreApp#optional_field1_value}
	OptionalField1Value *string `field:"optional" json:"optionalField1Value" yaml:"optionalField1Value"`
	// Name of optional param in the login form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#optional_field2 SecurePasswordStoreApp#optional_field2}
	OptionalField2 *string `field:"optional" json:"optionalField2" yaml:"optionalField2"`
	// Name of optional value in login form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#optional_field2_value SecurePasswordStoreApp#optional_field2_value}
	OptionalField2Value *string `field:"optional" json:"optionalField2Value" yaml:"optionalField2Value"`
	// Name of optional param in the login form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#optional_field3 SecurePasswordStoreApp#optional_field3}
	OptionalField3 *string `field:"optional" json:"optionalField3" yaml:"optionalField3"`
	// Name of optional value in login form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#optional_field3_value SecurePasswordStoreApp#optional_field3_value}
	OptionalField3Value *string `field:"optional" json:"optionalField3Value" yaml:"optionalField3Value"`
	// Allow user to reveal password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#reveal_password SecurePasswordStoreApp#reveal_password}
	RevealPassword interface{} `field:"optional" json:"revealPassword" yaml:"revealPassword"`
	// Shared password, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#shared_password SecurePasswordStoreApp#shared_password}
	SharedPassword *string `field:"optional" json:"sharedPassword" yaml:"sharedPassword"`
	// Shared username, required for certain schemes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#shared_username SecurePasswordStoreApp#shared_username}
	SharedUsername *string `field:"optional" json:"sharedUsername" yaml:"sharedUsername"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#skip_groups SecurePasswordStoreApp#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#skip_users SecurePasswordStoreApp#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// Status of application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#status SecurePasswordStoreApp#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#timeouts SecurePasswordStoreApp#timeouts}
	Timeouts *SecurePasswordStoreAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Username template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#user_name_template SecurePasswordStoreApp#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#user_name_template_push_status SecurePasswordStoreApp#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#user_name_template_suffix SecurePasswordStoreApp#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#user_name_template_type SecurePasswordStoreApp#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/secure_password_store_app#users SecurePasswordStoreApp#users}
	Users interface{} `field:"optional" json:"users" yaml:"users"`
}

