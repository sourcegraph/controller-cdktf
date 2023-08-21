package googledocumentaiwarehousedocumentschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledocumentaiwarehousedocumentschema/internal"
)

type GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DateTimeTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference
	DateTimeTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnumTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference
	EnumTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions
	FloatTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference
	FloatTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions
	// Experimental.
	Fqn() *string
	IntegerTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference
	IntegerTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsFilterable() interface{}
	SetIsFilterable(val interface{})
	IsFilterableInput() interface{}
	IsMetadata() interface{}
	SetIsMetadata(val interface{})
	IsMetadataInput() interface{}
	IsRepeatable() interface{}
	SetIsRepeatable(val interface{})
	IsRepeatableInput() interface{}
	IsRequired() interface{}
	SetIsRequired(val interface{})
	IsRequiredInput() interface{}
	IsSearchable() interface{}
	SetIsSearchable(val interface{})
	IsSearchableInput() interface{}
	MapTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference
	MapTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions
	Name() *string
	SetName(val *string)
	NameInput() *string
	PropertyTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference
	PropertyTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions
	RetrievalImportance() *string
	SetRetrievalImportance(val *string)
	RetrievalImportanceInput() *string
	SchemaSources() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList
	SchemaSourcesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference
	TextTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions
	TimestampTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference
	TimestampTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDateTimeTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions)
	PutEnumTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions)
	PutFloatTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions)
	PutIntegerTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions)
	PutMapTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions)
	PutPropertyTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions)
	PutSchemaSources(value interface{})
	PutTextTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions)
	PutTimestampTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions)
	ResetDateTimeTypeOptions()
	ResetDisplayName()
	ResetEnumTypeOptions()
	ResetFloatTypeOptions()
	ResetIntegerTypeOptions()
	ResetIsFilterable()
	ResetIsMetadata()
	ResetIsRepeatable()
	ResetIsRequired()
	ResetIsSearchable()
	ResetMapTypeOptions()
	ResetPropertyTypeOptions()
	ResetRetrievalImportance()
	ResetSchemaSources()
	ResetTextTypeOptions()
	ResetTimestampTypeOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference
type jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DateTimeTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"dateTimeTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DateTimeTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions
	_jsii_.Get(
		j,
		"dateTimeTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) EnumTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"enumTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) EnumTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions
	_jsii_.Get(
		j,
		"enumTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) FloatTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"floatTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) FloatTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions
	_jsii_.Get(
		j,
		"floatTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IntegerTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"integerTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IntegerTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions
	_jsii_.Get(
		j,
		"integerTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsFilterable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFilterable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsFilterableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFilterableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRepeatable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRepeatable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRepeatableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRepeatableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsSearchable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSearchable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsSearchableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSearchableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) MapTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"mapTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) MapTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions
	_jsii_.Get(
		j,
		"mapTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PropertyTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"propertyTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PropertyTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions
	_jsii_.Get(
		j,
		"propertyTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) RetrievalImportance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalImportance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) RetrievalImportanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalImportanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) SchemaSources() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList
	_jsii_.Get(
		j,
		"schemaSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) SchemaSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TextTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"textTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TextTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions
	_jsii_.Get(
		j,
		"textTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TimestampTypeOptions() GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference {
	var returns GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"timestampTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TimestampTypeOptionsInput() *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions {
	var returns *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions
	_jsii_.Get(
		j,
		"timestampTypeOptionsInput",
		&returns,
	)
	return returns
}


func NewGoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference_Override(g GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsFilterable(val interface{}) {
	if err := j.validateSetIsFilterableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFilterable",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsMetadata(val interface{}) {
	if err := j.validateSetIsMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMetadata",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsRepeatable(val interface{}) {
	if err := j.validateSetIsRepeatableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRepeatable",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsRequired(val interface{}) {
	if err := j.validateSetIsRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRequired",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsSearchable(val interface{}) {
	if err := j.validateSetIsSearchableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSearchable",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetRetrievalImportance(val *string) {
	if err := j.validateSetRetrievalImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrievalImportance",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutDateTimeTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions) {
	if err := g.validatePutDateTimeTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDateTimeTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutEnumTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions) {
	if err := g.validatePutEnumTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnumTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutFloatTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions) {
	if err := g.validatePutFloatTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFloatTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutIntegerTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions) {
	if err := g.validatePutIntegerTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIntegerTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutMapTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions) {
	if err := g.validatePutMapTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMapTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutPropertyTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions) {
	if err := g.validatePutPropertyTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPropertyTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutSchemaSources(value interface{}) {
	if err := g.validatePutSchemaSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaSources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutTextTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions) {
	if err := g.validatePutTextTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTextTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutTimestampTypeOptions(value *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions) {
	if err := g.validatePutTimestampTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimestampTypeOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetDateTimeTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetDateTimeTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetEnumTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetEnumTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetFloatTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetFloatTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIntegerTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetIntegerTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsFilterable() {
	_jsii_.InvokeVoid(
		g,
		"resetIsFilterable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetIsMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsRepeatable() {
	_jsii_.InvokeVoid(
		g,
		"resetIsRepeatable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetIsRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsSearchable() {
	_jsii_.InvokeVoid(
		g,
		"resetIsSearchable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetMapTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetMapTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetPropertyTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetPropertyTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetRetrievalImportance() {
	_jsii_.InvokeVoid(
		g,
		"resetRetrievalImportance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetSchemaSources() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetTextTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetTextTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetTimestampTypeOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetTimestampTypeOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

