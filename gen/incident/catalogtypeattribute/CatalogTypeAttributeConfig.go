package catalogtypeattribute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogTypeAttributeConfig struct {
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
	// ID of this catalog type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type_attribute#catalog_type_id CatalogTypeAttribute#catalog_type_id}
	CatalogTypeId *string `field:"required" json:"catalogTypeId" yaml:"catalogTypeId"`
	// The name of this attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type_attribute#name CatalogTypeAttribute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of this attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type_attribute#type CatalogTypeAttribute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether this attribute is an array or scalar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type_attribute#array CatalogTypeAttribute#array}
	Array any `field:"optional" json:"array" yaml:"array"`
	// If this is a backlink, the id of the attribute that it's linked from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type_attribute#backlink_attribute CatalogTypeAttribute#backlink_attribute}
	BacklinkAttribute *string `field:"optional" json:"backlinkAttribute" yaml:"backlinkAttribute"`
	// If this is a path attribute, the path that we should use to pull the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type_attribute#path CatalogTypeAttribute#path}
	Path *[]*string `field:"optional" json:"path" yaml:"path"`
	// If true, Terraform will only manage the schema of the attribute.
	//
	// Values for this attribute can be managed from the incident.io web dashboard.
	//
	// NOTE: When enabled, you should use the `managed_attributes` argument on either `incident_catalog_entry` or `incident_catalog_entries` to manage the values of other attributes on this type, without Terraform overwriting values set in the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type_attribute#schema_only CatalogTypeAttribute#schema_only}
	SchemaOnly any `field:"optional" json:"schemaOnly" yaml:"schemaOnly"`
}
