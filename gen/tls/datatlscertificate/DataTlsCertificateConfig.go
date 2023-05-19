package datatlscertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTlsCertificateConfig struct {
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
	// The content of the certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/data-sources/certificate#content DataTlsCertificate#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// URL of the endpoint to get the certificates from.
	//
	// Accepted schemes are: `https`, `tls`. For scheme `https://` it will use the HTTP protocol and apply the `proxy` configuration of the provider, if set. For scheme `tls://` it will instead use a secure TCP socket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/data-sources/certificate#url DataTlsCertificate#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Whether to verify the certificate chain while parsing it or not (default: `true`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/data-sources/certificate#verify_chain DataTlsCertificate#verify_chain}
	VerifyChain interface{} `field:"optional" json:"verifyChain" yaml:"verifyChain"`
}

