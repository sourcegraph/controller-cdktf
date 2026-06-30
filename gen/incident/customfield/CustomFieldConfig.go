package customfield

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomFieldConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Description of the custom field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field#description CustomField#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Type of custom field. Possible values are: `single_select`, `multi_select`, `text`, `link`, `numeric`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field#field_type CustomField#field_type}
	FieldType *string `field:"required" json:"fieldType" yaml:"fieldType"`
	// Human readable name for the custom field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field#name CustomField#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// For catalog fields, the ID of the associated catalog type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field#catalog_type_id CustomField#catalog_type_id}
	CatalogTypeId *string `field:"optional" json:"catalogTypeId" yaml:"catalogTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field#filter_by CustomField#filter_by}.
	FilterBy *CustomFieldFilterBy `field:"optional" json:"filterBy" yaml:"filterBy"`
	// For catalog fields, the ID of the attribute used to group catalog entries (if applicable).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field#group_by_catalog_attribute_id CustomField#group_by_catalog_attribute_id}
	GroupByCatalogAttributeId *string `field:"optional" json:"groupByCatalogAttributeId" yaml:"groupByCatalogAttributeId"`
	// Which catalog attribute provides helptext for the options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field#helptext_catalog_attribute_id CustomField#helptext_catalog_attribute_id}
	HelptextCatalogAttributeId *string `field:"optional" json:"helptextCatalogAttributeId" yaml:"helptextCatalogAttributeId"`
}
