package domain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DomainConfig struct {
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
	// Custom Domain name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain#name Domain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional.
	//
	// Certificate source type that indicates whether the certificate is provided by the user or Okta. Accepted values: MANUAL, OKTA_MANAGED. Warning: Use of OKTA_MANAGED requires a feature flag to be enabled. Default value = MANUAL
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain#certificate_source_type Domain#certificate_source_type}
	CertificateSourceType *string `field:"optional" json:"certificateSourceType" yaml:"certificateSourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain#id Domain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether the domain should be verified during creation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/domain#verify Domain#verify}
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
}

