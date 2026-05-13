package cestool


type CesToolClientFunctionParameters struct {
	// The type of the data. Possible values: STRING INTEGER NUMBER BOOLEAN OBJECT ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#type CesTool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Defines the schema for additional properties allowed in an object.
	//
	// The value must be a valid JSON string representing the Schema object.
	// (Note: OpenAPI also allows a boolean, this definition expects a Schema JSON).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#additional_properties CesTool#additional_properties}
	AdditionalProperties *string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// The instance value should be valid against at least one of the schemas in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#any_of CesTool#any_of}
	AnyOf *string `field:"optional" json:"anyOf" yaml:"anyOf"`
	// Default value of the data.
	//
	// Represents a dynamically typed value
	// which can be either null, a number, a string, a boolean, a struct,
	// or a list of values. The provided default value must be compatible
	// with the defined 'type' and other schema constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#default CesTool#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// A map of definitions for use by ref. Only allowed at the root of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#defs CesTool#defs}
	Defs *string `field:"optional" json:"defs" yaml:"defs"`
	// The description of the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#description CesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Possible values of the element of primitive type with enum format.
	//
	// Examples:
	// 1. We can define direction as :
	// {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	// 2. We can define apartment number as :
	// {type:INTEGER, format:enum, enum:["101", "201", "301"]}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#enum CesTool#enum}
	Enum *[]*string `field:"optional" json:"enum" yaml:"enum"`
	// Schema of the elements of Type.ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#items CesTool#items}
	Items *string `field:"optional" json:"items" yaml:"items"`
	// Maximum value for Type.INTEGER and Type.NUMBER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#maximum CesTool#maximum}
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// Maximum number of the elements for Type.ARRAY. (int64 format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#max_items CesTool#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Minimum value for Type.INTEGER and Type.NUMBER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#minimum CesTool#minimum}
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// Minimum number of the elements for Type.ARRAY. (int64 format).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#min_items CesTool#min_items}
	MinItems *float64 `field:"optional" json:"minItems" yaml:"minItems"`
	// Indicates if the value may be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#nullable CesTool#nullable}
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// Schemas of initial elements of Type.ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#prefix_items CesTool#prefix_items}
	PrefixItems *string `field:"optional" json:"prefixItems" yaml:"prefixItems"`
	// Properties of Type.OBJECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#properties CesTool#properties}
	Properties *string `field:"optional" json:"properties" yaml:"properties"`
	// Allows indirect references between schema nodes.
	//
	// The value should be a
	// valid reference to a child of the root 'defs'.
	// For example, the following schema defines a reference to a schema node
	// named "Pet":
	// type: object
	// properties:
	//   pet:
	//     ref: #/defs/Pet
	// defs:
	//   Pet:
	//     type: object
	//     properties:
	//       name:
	//         type: string
	// The value of the "pet" property is a reference to the schema node
	// named "Pet".
	// See details in
	// https://json-schema.org/understanding-json-schema/structuring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#ref CesTool#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// Required properties of Type.OBJECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#required CesTool#required}
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
	// The title of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#title CesTool#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Indicate the items in the array must be unique. Only applies to TYPE.ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#unique_items CesTool#unique_items}
	UniqueItems interface{} `field:"optional" json:"uniqueItems" yaml:"uniqueItems"`
}

