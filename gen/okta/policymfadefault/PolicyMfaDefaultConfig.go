package policymfadefault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyMfaDefaultConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#duo PolicyMfaDefault#duo}.
	Duo *map[string]*string `field:"optional" json:"duo" yaml:"duo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#external_idp PolicyMfaDefault#external_idp}.
	ExternalIdp *map[string]*string `field:"optional" json:"externalIdp" yaml:"externalIdp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#fido_u2f PolicyMfaDefault#fido_u2f}.
	FidoU2F *map[string]*string `field:"optional" json:"fidoU2F" yaml:"fidoU2F"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#fido_webauthn PolicyMfaDefault#fido_webauthn}.
	FidoWebauthn *map[string]*string `field:"optional" json:"fidoWebauthn" yaml:"fidoWebauthn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#google_otp PolicyMfaDefault#google_otp}.
	GoogleOtp *map[string]*string `field:"optional" json:"googleOtp" yaml:"googleOtp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#hotp PolicyMfaDefault#hotp}.
	Hotp *map[string]*string `field:"optional" json:"hotp" yaml:"hotp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#id PolicyMfaDefault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Is the policy using Okta Identity Engine (OIE) with authenticators instead of factors?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#is_oie PolicyMfaDefault#is_oie}
	IsOie interface{} `field:"optional" json:"isOie" yaml:"isOie"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_call PolicyMfaDefault#okta_call}.
	OktaCall *map[string]*string `field:"optional" json:"oktaCall" yaml:"oktaCall"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_email PolicyMfaDefault#okta_email}.
	OktaEmail *map[string]*string `field:"optional" json:"oktaEmail" yaml:"oktaEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_otp PolicyMfaDefault#okta_otp}.
	OktaOtp *map[string]*string `field:"optional" json:"oktaOtp" yaml:"oktaOtp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_password PolicyMfaDefault#okta_password}.
	OktaPassword *map[string]*string `field:"optional" json:"oktaPassword" yaml:"oktaPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_push PolicyMfaDefault#okta_push}.
	OktaPush *map[string]*string `field:"optional" json:"oktaPush" yaml:"oktaPush"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_question PolicyMfaDefault#okta_question}.
	OktaQuestion *map[string]*string `field:"optional" json:"oktaQuestion" yaml:"oktaQuestion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_sms PolicyMfaDefault#okta_sms}.
	OktaSms *map[string]*string `field:"optional" json:"oktaSms" yaml:"oktaSms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#okta_verify PolicyMfaDefault#okta_verify}.
	OktaVerify *map[string]*string `field:"optional" json:"oktaVerify" yaml:"oktaVerify"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#onprem_mfa PolicyMfaDefault#onprem_mfa}.
	OnpremMfa *map[string]*string `field:"optional" json:"onpremMfa" yaml:"onpremMfa"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#phone_number PolicyMfaDefault#phone_number}.
	PhoneNumber *map[string]*string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#rsa_token PolicyMfaDefault#rsa_token}.
	RsaToken *map[string]*string `field:"optional" json:"rsaToken" yaml:"rsaToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#security_question PolicyMfaDefault#security_question}.
	SecurityQuestion *map[string]*string `field:"optional" json:"securityQuestion" yaml:"securityQuestion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#symantec_vip PolicyMfaDefault#symantec_vip}.
	SymantecVip *map[string]*string `field:"optional" json:"symantecVip" yaml:"symantecVip"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#webauthn PolicyMfaDefault#webauthn}.
	Webauthn *map[string]*string `field:"optional" json:"webauthn" yaml:"webauthn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_mfa_default#yubikey_token PolicyMfaDefault#yubikey_token}.
	YubikeyToken *map[string]*string `field:"optional" json:"yubikeyToken" yaml:"yubikeyToken"`
}

