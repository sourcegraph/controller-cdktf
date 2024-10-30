package authenticator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticatorConfig struct {
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
	// A human-readable string that identifies the authenticator.
	//
	// Some authenticators are available by feature flag on the organization. Possible values inclue: `duo`, `external_idp`, `google_otp`, `okta_email`, `okta_password`, `okta_verify`, `onprem_mfa`, `phone_number`, `rsa_token`, `security_question`, `webauthn`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#key Authenticator#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Display name of the Authenticator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#name Authenticator#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#id Authenticator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name does not trigger change detection (legacy behavior).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#legacy_ignore_name Authenticator#legacy_ignore_name}
	LegacyIgnoreName interface{} `field:"optional" json:"legacyIgnoreName" yaml:"legacyIgnoreName"`
	// The RADIUS server port (for example 1812).
	//
	// This is defined when the On-Prem RADIUS server is configured. Used only for authenticators with type `security_key`.  Conflicts with `provider_json` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_auth_port Authenticator#provider_auth_port}
	ProviderAuthPort *float64 `field:"optional" json:"providerAuthPort" yaml:"providerAuthPort"`
	// (DUO specific) - The Duo Security API hostname. Conflicts with `provider_json` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_host Authenticator#provider_host}
	ProviderHost *string `field:"optional" json:"providerHost" yaml:"providerHost"`
	// Server host name or IP address.
	//
	// Default is `localhost`. Used only for authenticators with type `security_key`. Conflicts with `provider_json` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_hostname Authenticator#provider_hostname}
	ProviderHostname *string `field:"optional" json:"providerHostname" yaml:"providerHostname"`
	// (DUO specific) - The Duo Security integration key.  Conflicts with `provider_json` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_integration_key Authenticator#provider_integration_key}
	ProviderIntegrationKey *string `field:"optional" json:"providerIntegrationKey" yaml:"providerIntegrationKey"`
	// Provider JSON allows for expressive providervalues.
	//
	// This argument conflicts with the other 'provider_xxx' arguments. The [CreateProvider](https://developer.okta.com/docs/reference/api/authenticators-admin/#request) illustrates detailed provider values for a Duo authenticator. [Provider values](https://developer.okta.com/docs/reference/api/authenticators-admin/#authenticators-administration-api-object)are listed in Okta API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_json Authenticator#provider_json}
	ProviderJson *string `field:"optional" json:"providerJson" yaml:"providerJson"`
	// (DUO specific) - The Duo Security secret key.  Conflicts with `provider_json` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_secret_key Authenticator#provider_secret_key}
	ProviderSecretKey *string `field:"optional" json:"providerSecretKey" yaml:"providerSecretKey"`
	// An authentication key that must be defined when the RADIUS server is configured, and must be the same on both the RADIUS client and server.
	//
	// Used only for authenticators with type `security_key`. Conflicts with `provider_json` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_shared_secret Authenticator#provider_shared_secret}
	ProviderSharedSecret *string `field:"optional" json:"providerSharedSecret" yaml:"providerSharedSecret"`
	// Username template expected by the provider. Used only for authenticators with type `security_key`.  Conflicts with `provider_json` argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#provider_user_name_template Authenticator#provider_user_name_template}
	ProviderUserNameTemplate *string `field:"optional" json:"providerUserNameTemplate" yaml:"providerUserNameTemplate"`
	// Settings for the authenticator.
	//
	// The settings JSON contains values based on Authenticator key. It is not used for authenticators with type `security_key`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#settings Authenticator#settings}
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// Authenticator status: `ACTIVE` or `INACTIVE`. Default: `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator#status Authenticator#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

