package googlecesapp


type GoogleCesAppVariableDeclarationsSchema struct {
	// The type of the data. Possible values: STRING INTEGER NUMBER BOOLEAN OBJECT ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#type GoogleCesApp#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Optional.
	//
	// Defines the schema for additional properties allowed in an object.
	// The value must be a valid JSON string representing the Schema object.
	// (Note: OpenAPI also allows a boolean, this definition expects a Schema JSON).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#additional_properties GoogleCesApp#additional_properties}
	AdditionalProperties *string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Optional. The instance value should be valid against at least one of the schemas in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#any_of GoogleCesApp#any_of}
	AnyOf *string `field:"optional" json:"anyOf" yaml:"anyOf"`
	// Optional.
	//
	// Default value of the data. Represents a dynamically typed value
	// which can be either null, a number, a string, a boolean, a struct,
	// or a list of values. The provided default value must be encoded as a JSON string.
	// Use 'jsonencode' in Terraform HCL to encode the default value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#default GoogleCesApp#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// A map of definitions for use by ref. Only allowed at the root of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#defs GoogleCesApp#defs}
	Defs *string `field:"optional" json:"defs" yaml:"defs"`
	// The description of the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#description GoogleCesApp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Possible values of the element of primitive type with enum format.
	//
	// Examples:
	// 1. We can define direction as :
	// {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	// 2. We can define apartment number as :
	// {type:INTEGER, format:enum, enum:["101", "201", "301"]}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#enum GoogleCesApp#enum}
	Enum *[]*string `field:"optional" json:"enum" yaml:"enum"`
	// Schema of the elements of Type.ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#items GoogleCesApp#items}
	Items *string `field:"optional" json:"items" yaml:"items"`
	// Indicates if the value may be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#nullable GoogleCesApp#nullable}
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// Optional. Schemas of initial elements of Type.ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#prefix_items GoogleCesApp#prefix_items}
	PrefixItems *string `field:"optional" json:"prefixItems" yaml:"prefixItems"`
	// Properties of Type.OBJECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#properties GoogleCesApp#properties}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#ref GoogleCesApp#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// Required properties of Type.OBJECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#required GoogleCesApp#required}
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
	// The title of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#title GoogleCesApp#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Indicate the items in the array must be unique. Only applies to TYPE.ARRAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#unique_items GoogleCesApp#unique_items}
	UniqueItems interface{} `field:"optional" json:"uniqueItems" yaml:"uniqueItems"`
}

