package domaincertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DomainCertificateConfig struct {
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
	// Certificate content.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain_certificate#certificate DomainCertificate#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Certificate chain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain_certificate#certificate_chain DomainCertificate#certificate_chain}
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// Domain's ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain_certificate#domain_id DomainCertificate#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Certificate private key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain_certificate#private_key DomainCertificate#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain_certificate#id DomainCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Certificate type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain_certificate#type DomainCertificate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

