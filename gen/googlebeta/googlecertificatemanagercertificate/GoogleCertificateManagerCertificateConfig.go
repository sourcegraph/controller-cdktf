package googlecertificatemanagercertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCertificateManagerCertificateConfig struct {
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
	// A user-defined name of the certificate.
	//
	// Certificate names must be unique
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#name GoogleCertificateManagerCertificate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#description GoogleCertificateManagerCertificate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#id GoogleCertificateManagerCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the Certificate resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#labels GoogleCertificateManagerCertificate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// managed block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#managed GoogleCertificateManagerCertificate#managed}
	Managed *GoogleCertificateManagerCertificateManaged `field:"optional" json:"managed" yaml:"managed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#project GoogleCertificateManagerCertificate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The scope of the certificate.
	//
	// DEFAULT: Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	//
	// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// Currently allowed only for managed certificates.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#scope GoogleCertificateManagerCertificate#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// self_managed block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#self_managed GoogleCertificateManagerCertificate#self_managed}
	SelfManaged *GoogleCertificateManagerCertificateSelfManaged `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#timeouts GoogleCertificateManagerCertificate#timeouts}
	Timeouts *GoogleCertificateManagerCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

