package googlecomputemanagedsslcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeManagedSslCertificateConfig struct {
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
	// The unique identifier for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#certificate_id GoogleComputeManagedSslCertificate#certificate_id}
	CertificateId *float64 `field:"optional" json:"certificateId" yaml:"certificateId"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#description GoogleComputeManagedSslCertificate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#id GoogleComputeManagedSslCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#managed GoogleComputeManagedSslCertificate#managed}
	Managed *GoogleComputeManagedSslCertificateManaged `field:"optional" json:"managed" yaml:"managed"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// These are in the same namespace as the managed SSL certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#name GoogleComputeManagedSslCertificate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#project GoogleComputeManagedSslCertificate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#timeouts GoogleComputeManagedSslCertificate#timeouts}
	Timeouts *GoogleComputeManagedSslCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Enum field whose value is always 'MANAGED' - used to signal to the API which type this is.
	//
	// Default value: "MANAGED" Possible values: ["MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_managed_ssl_certificate#type GoogleComputeManagedSslCertificate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

