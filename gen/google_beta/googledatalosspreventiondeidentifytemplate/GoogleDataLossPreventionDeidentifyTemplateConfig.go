package googledatalosspreventiondeidentifytemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataLossPreventionDeidentifyTemplateConfig struct {
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
	// deidentify_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#deidentify_config GoogleDataLossPreventionDeidentifyTemplate#deidentify_config}
	DeidentifyConfig *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfig `field:"required" json:"deidentifyConfig" yaml:"deidentifyConfig"`
	// The parent of the template in any of the following formats:.
	//
	// * 'projects/{{project}}'
	// * 'projects/{{project}}/locations/{{location}}'
	// * 'organizations/{{organization_id}}'
	// * 'organizations/{{organization_id}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#parent GoogleDataLossPreventionDeidentifyTemplate#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// A description of the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#description GoogleDataLossPreventionDeidentifyTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User set display name of the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#display_name GoogleDataLossPreventionDeidentifyTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#id GoogleDataLossPreventionDeidentifyTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The template id can contain uppercase and lowercase letters, numbers, and hyphens;
	//
	// that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is
	// 100 characters. Can be empty to allow the system to generate one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#template_id GoogleDataLossPreventionDeidentifyTemplate#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_deidentify_template#timeouts GoogleDataLossPreventionDeidentifyTemplate#timeouts}
	Timeouts *GoogleDataLossPreventionDeidentifyTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

