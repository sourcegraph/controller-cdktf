package groupschemaproperty

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupSchemaPropertyConfig struct {
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
	// Subschema unique string identifier.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#index GroupSchemaProperty#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// Subschema title (display name).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#title GroupSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Subschema type: string, boolean, number, integer, array, or object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#type GroupSchemaProperty#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Custom Subschema enumerated value of a property of type array.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#array_enum GroupSchemaProperty#array_enum}
	ArrayEnum *[]*string `field:"optional" json:"arrayEnum" yaml:"arrayEnum"`
	// array_one_of block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#array_one_of GroupSchemaProperty#array_one_of}
	ArrayOneOf interface{} `field:"optional" json:"arrayOneOf" yaml:"arrayOneOf"`
	// Subschema array type: string, number, integer, reference. Type field must be an array.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#array_type GroupSchemaProperty#array_type}
	ArrayType *string `field:"optional" json:"arrayType" yaml:"arrayType"`
	// Custom Subschema description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#description GroupSchemaProperty#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Custom Subschema enumerated value of the property. see: developer.okta.com/docs/api/resources/schemas#user-profile-schema-property-object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#enum GroupSchemaProperty#enum}
	Enum *[]*string `field:"optional" json:"enum" yaml:"enum"`
	// Subschema external name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#external_name GroupSchemaProperty#external_name}
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// Subschema external namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#external_namespace GroupSchemaProperty#external_namespace}
	ExternalNamespace *string `field:"optional" json:"externalNamespace" yaml:"externalNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#id GroupSchemaProperty#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// SubSchema profile manager, if not set it will inherit its setting.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#master GroupSchemaProperty#master}
	Master *string `field:"optional" json:"master" yaml:"master"`
	// master_override_priority block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#master_override_priority GroupSchemaProperty#master_override_priority}
	MasterOverridePriority interface{} `field:"optional" json:"masterOverridePriority" yaml:"masterOverridePriority"`
	// Subschema of type string maximum length.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#max_length GroupSchemaProperty#max_length}
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Subschema of type string minimum length.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#min_length GroupSchemaProperty#min_length}
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// one_of block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#one_of GroupSchemaProperty#one_of}
	OneOf interface{} `field:"optional" json:"oneOf" yaml:"oneOf"`
	// SubSchema permissions: HIDE, READ_ONLY, or READ_WRITE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#permissions GroupSchemaProperty#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
	// Whether the subschema is required.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#required GroupSchemaProperty#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#scope GroupSchemaProperty#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Subschema unique restriction.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property#unique GroupSchemaProperty#unique}
	Unique *string `field:"optional" json:"unique" yaml:"unique"`
}

