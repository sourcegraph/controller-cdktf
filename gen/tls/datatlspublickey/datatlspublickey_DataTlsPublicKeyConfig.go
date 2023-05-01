package datatlspublickey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTlsPublicKeyConfig struct {
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
	// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is _mutually exclusive_ with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/d/public_key#private_key_openssh DataTlsPublicKey#private_key_openssh}
	PrivateKeyOpenssh *string `field:"optional" json:"privateKeyOpenssh" yaml:"privateKeyOpenssh"`
	// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is _mutually exclusive_ with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/d/public_key#private_key_pem DataTlsPublicKey#private_key_pem}
	PrivateKeyPem *string `field:"optional" json:"privateKeyPem" yaml:"privateKeyPem"`
}

