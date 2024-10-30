package factortotp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FactorTotpConfig struct {
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
	// The TOTP name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/factor_totp#name FactorTotp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Clock drift interval.
	//
	// This setting allows you to build in tolerance for any drift between the token's current time and the server's current time. Valid values: `3`, `5`, `10`. Default is `3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/factor_totp#clock_drift_interval FactorTotp#clock_drift_interval}
	ClockDriftInterval *float64 `field:"optional" json:"clockDriftInterval" yaml:"clockDriftInterval"`
	// HMAC Algorithm. Valid values: `HMacSHA1`, `HMacSHA256`, `HMacSHA512`. Default is `HMacSHA512`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/factor_totp#hmac_algorithm FactorTotp#hmac_algorithm}
	HmacAlgorithm *string `field:"optional" json:"hmacAlgorithm" yaml:"hmacAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/factor_totp#id FactorTotp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Length of the password. Default is `6`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/factor_totp#otp_length FactorTotp#otp_length}
	OtpLength *float64 `field:"optional" json:"otpLength" yaml:"otpLength"`
	// Shared secret encoding. Valid values: `base32`, `base64`, `hexadecimal`. Default is `base32`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/factor_totp#shared_secret_encoding FactorTotp#shared_secret_encoding}
	SharedSecretEncoding *string `field:"optional" json:"sharedSecretEncoding" yaml:"sharedSecretEncoding"`
	// Time step in seconds. Valid values: `15`, `30`, `60`. Default is `15`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/factor_totp#time_step FactorTotp#time_step}
	TimeStep *float64 `field:"optional" json:"timeStep" yaml:"timeStep"`
}

