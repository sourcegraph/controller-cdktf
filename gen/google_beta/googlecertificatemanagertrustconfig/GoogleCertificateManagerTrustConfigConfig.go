package googlecertificatemanagertrustconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCertificateManagerTrustConfigConfig struct {
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
	// The trust config location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#location GoogleCertificateManagerTrustConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A user-defined name of the trust config. Trust config names must be unique globally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#name GoogleCertificateManagerTrustConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// allowlisted_certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#allowlisted_certificates GoogleCertificateManagerTrustConfig#allowlisted_certificates}
	AllowlistedCertificates interface{} `field:"optional" json:"allowlistedCertificates" yaml:"allowlistedCertificates"`
	// One or more paragraphs of text description of a trust config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#description GoogleCertificateManagerTrustConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#id GoogleCertificateManagerTrustConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the trust config.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#labels GoogleCertificateManagerTrustConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#project GoogleCertificateManagerTrustConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#timeouts GoogleCertificateManagerTrustConfig#timeouts}
	Timeouts *GoogleCertificateManagerTrustConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// trust_stores block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#trust_stores GoogleCertificateManagerTrustConfig#trust_stores}
	TrustStores interface{} `field:"optional" json:"trustStores" yaml:"trustStores"`
}

