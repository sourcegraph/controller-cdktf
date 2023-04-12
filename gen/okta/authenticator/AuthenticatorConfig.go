package authenticator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticatorConfig struct {
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
	// A human-readable string that identifies the Authenticator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#key Authenticator#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Display name of the Authenticator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#name Authenticator#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#id Authenticator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The RADIUS server port (for example 1812). This is defined when the On-Prem RADIUS server is configured.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_auth_port Authenticator#provider_auth_port}
	ProviderAuthPort *float64 `field:"optional" json:"providerAuthPort" yaml:"providerAuthPort"`
	// The Duo Security API hostname.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_host Authenticator#provider_host}
	ProviderHost *string `field:"optional" json:"providerHost" yaml:"providerHost"`
	// Server host name or IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_hostname Authenticator#provider_hostname}
	ProviderHostname *string `field:"optional" json:"providerHostname" yaml:"providerHostname"`
	// The Duo Security integration key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_integration_key Authenticator#provider_integration_key}
	ProviderIntegrationKey *string `field:"optional" json:"providerIntegrationKey" yaml:"providerIntegrationKey"`
	// Provider in JSON format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_json Authenticator#provider_json}
	ProviderJson *string `field:"optional" json:"providerJson" yaml:"providerJson"`
	// The Duo Security secret key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_secret_key Authenticator#provider_secret_key}
	ProviderSecretKey *string `field:"optional" json:"providerSecretKey" yaml:"providerSecretKey"`
	// An authentication key that must be defined when the RADIUS server is configured, and must be the same on both the RADIUS client and server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_shared_secret Authenticator#provider_shared_secret}
	ProviderSharedSecret *string `field:"optional" json:"providerSharedSecret" yaml:"providerSharedSecret"`
	// Format expected by the provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#provider_user_name_template Authenticator#provider_user_name_template}
	ProviderUserNameTemplate *string `field:"optional" json:"providerUserNameTemplate" yaml:"providerUserNameTemplate"`
	// Authenticator settings in JSON format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#settings Authenticator#settings}
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// Authenticator status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/authenticator#status Authenticator#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

