package customizedsigninpage


type CustomizedSigninPageWidgetCustomizations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#widget_generation CustomizedSigninPage#widget_generation}.
	WidgetGeneration *string `field:"required" json:"widgetGeneration" yaml:"widgetGeneration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#authenticator_page_custom_link_label CustomizedSigninPage#authenticator_page_custom_link_label}.
	AuthenticatorPageCustomLinkLabel *string `field:"optional" json:"authenticatorPageCustomLinkLabel" yaml:"authenticatorPageCustomLinkLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#authenticator_page_custom_link_url CustomizedSigninPage#authenticator_page_custom_link_url}.
	AuthenticatorPageCustomLinkUrl *string `field:"optional" json:"authenticatorPageCustomLinkUrl" yaml:"authenticatorPageCustomLinkUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#classic_recovery_flow_email_or_username_label CustomizedSigninPage#classic_recovery_flow_email_or_username_label}.
	ClassicRecoveryFlowEmailOrUsernameLabel *string `field:"optional" json:"classicRecoveryFlowEmailOrUsernameLabel" yaml:"classicRecoveryFlowEmailOrUsernameLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#custom_link_1_label CustomizedSigninPage#custom_link_1_label}.
	CustomLink1Label *string `field:"optional" json:"customLink1Label" yaml:"customLink1Label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#custom_link_1_url CustomizedSigninPage#custom_link_1_url}.
	CustomLink1Url *string `field:"optional" json:"customLink1Url" yaml:"customLink1Url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#custom_link_2_label CustomizedSigninPage#custom_link_2_label}.
	CustomLink2Label *string `field:"optional" json:"customLink2Label" yaml:"customLink2Label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#custom_link_2_url CustomizedSigninPage#custom_link_2_url}.
	CustomLink2Url *string `field:"optional" json:"customLink2Url" yaml:"customLink2Url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#forgot_password_label CustomizedSigninPage#forgot_password_label}.
	ForgotPasswordLabel *string `field:"optional" json:"forgotPasswordLabel" yaml:"forgotPasswordLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#forgot_password_url CustomizedSigninPage#forgot_password_url}.
	ForgotPasswordUrl *string `field:"optional" json:"forgotPasswordUrl" yaml:"forgotPasswordUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#help_label CustomizedSigninPage#help_label}.
	HelpLabel *string `field:"optional" json:"helpLabel" yaml:"helpLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#help_url CustomizedSigninPage#help_url}.
	HelpUrl *string `field:"optional" json:"helpUrl" yaml:"helpUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#password_info_tip CustomizedSigninPage#password_info_tip}.
	PasswordInfoTip *string `field:"optional" json:"passwordInfoTip" yaml:"passwordInfoTip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#password_label CustomizedSigninPage#password_label}.
	PasswordLabel *string `field:"optional" json:"passwordLabel" yaml:"passwordLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#show_password_visibility_toggle CustomizedSigninPage#show_password_visibility_toggle}.
	ShowPasswordVisibilityToggle interface{} `field:"optional" json:"showPasswordVisibilityToggle" yaml:"showPasswordVisibilityToggle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#show_user_identifier CustomizedSigninPage#show_user_identifier}.
	ShowUserIdentifier interface{} `field:"optional" json:"showUserIdentifier" yaml:"showUserIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#sign_in_label CustomizedSigninPage#sign_in_label}.
	SignInLabel *string `field:"optional" json:"signInLabel" yaml:"signInLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#unlock_account_label CustomizedSigninPage#unlock_account_label}.
	UnlockAccountLabel *string `field:"optional" json:"unlockAccountLabel" yaml:"unlockAccountLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#unlock_account_url CustomizedSigninPage#unlock_account_url}.
	UnlockAccountUrl *string `field:"optional" json:"unlockAccountUrl" yaml:"unlockAccountUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#username_info_tip CustomizedSigninPage#username_info_tip}.
	UsernameInfoTip *string `field:"optional" json:"usernameInfoTip" yaml:"usernameInfoTip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/customized_signin_page#username_label CustomizedSigninPage#username_label}.
	UsernameLabel *string `field:"optional" json:"usernameLabel" yaml:"usernameLabel"`
}

