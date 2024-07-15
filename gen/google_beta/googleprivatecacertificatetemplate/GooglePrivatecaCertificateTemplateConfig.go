package googleprivatecacertificatetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePrivatecaCertificateTemplateConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#location GooglePrivatecaCertificateTemplate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name for this CertificateTemplate in the format 'projects/* /locations/* /certificateTemplates/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#name GooglePrivatecaCertificateTemplate#name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional. A human-readable description of scenarios this template is intended for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#description GooglePrivatecaCertificateTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#id GooglePrivatecaCertificateTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#identity_constraints GooglePrivatecaCertificateTemplate#identity_constraints}
	IdentityConstraints *GooglePrivatecaCertificateTemplateIdentityConstraints `field:"optional" json:"identityConstraints" yaml:"identityConstraints"`
	// Optional. Labels with user-defined metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#labels GooglePrivatecaCertificateTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Optional.
	//
	// The maximum lifetime allowed for all issued certificates that use this template. If the issuing CaPool's IssuancePolicy specifies a maximum lifetime the minimum of the two durations will be the maximum lifetime for issued. Note that if the issuing CertificateAuthority expires before a Certificate's requested maximum_lifetime, the effective lifetime will be explicitly truncated to match it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#maximum_lifetime GooglePrivatecaCertificateTemplate#maximum_lifetime}
	MaximumLifetime *string `field:"optional" json:"maximumLifetime" yaml:"maximumLifetime"`
	// passthrough_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#passthrough_extensions GooglePrivatecaCertificateTemplate#passthrough_extensions}
	PassthroughExtensions *GooglePrivatecaCertificateTemplatePassthroughExtensions `field:"optional" json:"passthroughExtensions" yaml:"passthroughExtensions"`
	// predefined_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#predefined_values GooglePrivatecaCertificateTemplate#predefined_values}
	PredefinedValues *GooglePrivatecaCertificateTemplatePredefinedValues `field:"optional" json:"predefinedValues" yaml:"predefinedValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#project GooglePrivatecaCertificateTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_template#timeouts GooglePrivatecaCertificateTemplate#timeouts}
	Timeouts *GooglePrivatecaCertificateTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

