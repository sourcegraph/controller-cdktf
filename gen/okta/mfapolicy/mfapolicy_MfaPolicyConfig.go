package mfapolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MfaPolicyConfig struct {
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
	// Policy Name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#name MfaPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policy Description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#description MfaPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#duo MfaPolicy#duo}.
	Duo *map[string]*string `field:"optional" json:"duo" yaml:"duo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#external_idp MfaPolicy#external_idp}.
	ExternalIdp *map[string]*string `field:"optional" json:"externalIdp" yaml:"externalIdp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#fido_u2f MfaPolicy#fido_u2f}.
	FidoU2F *map[string]*string `field:"optional" json:"fidoU2F" yaml:"fidoU2F"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#fido_webauthn MfaPolicy#fido_webauthn}.
	FidoWebauthn *map[string]*string `field:"optional" json:"fidoWebauthn" yaml:"fidoWebauthn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#google_otp MfaPolicy#google_otp}.
	GoogleOtp *map[string]*string `field:"optional" json:"googleOtp" yaml:"googleOtp"`
	// List of Group IDs to Include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#groups_included MfaPolicy#groups_included}
	GroupsIncluded *[]*string `field:"optional" json:"groupsIncluded" yaml:"groupsIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#hotp MfaPolicy#hotp}.
	Hotp *map[string]*string `field:"optional" json:"hotp" yaml:"hotp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#id MfaPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Is the policy using Okta Identity Engine (OIE) with authenticators instead of factors?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#is_oie MfaPolicy#is_oie}
	IsOie interface{} `field:"optional" json:"isOie" yaml:"isOie"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_call MfaPolicy#okta_call}.
	OktaCall *map[string]*string `field:"optional" json:"oktaCall" yaml:"oktaCall"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_email MfaPolicy#okta_email}.
	OktaEmail *map[string]*string `field:"optional" json:"oktaEmail" yaml:"oktaEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_otp MfaPolicy#okta_otp}.
	OktaOtp *map[string]*string `field:"optional" json:"oktaOtp" yaml:"oktaOtp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_password MfaPolicy#okta_password}.
	OktaPassword *map[string]*string `field:"optional" json:"oktaPassword" yaml:"oktaPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_push MfaPolicy#okta_push}.
	OktaPush *map[string]*string `field:"optional" json:"oktaPush" yaml:"oktaPush"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_question MfaPolicy#okta_question}.
	OktaQuestion *map[string]*string `field:"optional" json:"oktaQuestion" yaml:"oktaQuestion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_sms MfaPolicy#okta_sms}.
	OktaSms *map[string]*string `field:"optional" json:"oktaSms" yaml:"oktaSms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#okta_verify MfaPolicy#okta_verify}.
	OktaVerify *map[string]*string `field:"optional" json:"oktaVerify" yaml:"oktaVerify"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#onprem_mfa MfaPolicy#onprem_mfa}.
	OnpremMfa *map[string]*string `field:"optional" json:"onpremMfa" yaml:"onpremMfa"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#phone_number MfaPolicy#phone_number}.
	PhoneNumber *map[string]*string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Policy Priority, this attribute can be set to a valid priority.
	//
	// To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last (lowest) if not there.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#priority MfaPolicy#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#rsa_token MfaPolicy#rsa_token}.
	RsaToken *map[string]*string `field:"optional" json:"rsaToken" yaml:"rsaToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#security_question MfaPolicy#security_question}.
	SecurityQuestion *map[string]*string `field:"optional" json:"securityQuestion" yaml:"securityQuestion"`
	// Policy Status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#status MfaPolicy#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#symantec_vip MfaPolicy#symantec_vip}.
	SymantecVip *map[string]*string `field:"optional" json:"symantecVip" yaml:"symantecVip"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#webauthn MfaPolicy#webauthn}.
	Webauthn *map[string]*string `field:"optional" json:"webauthn" yaml:"webauthn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/mfa_policy#yubikey_token MfaPolicy#yubikey_token}.
	YubikeyToken *map[string]*string `field:"optional" json:"yubikeyToken" yaml:"yubikeyToken"`
}

