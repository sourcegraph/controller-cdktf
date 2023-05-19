package googlebigqueryjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryjob/internal"
)

type GoogleBigqueryJobExtractOutputReference interface {
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
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationFormat() *string
	SetDestinationFormat(val *string)
	DestinationFormatInput() *string
	DestinationUris() *[]*string
	SetDestinationUris(val *[]*string)
	DestinationUrisInput() *[]*string
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryJobExtract
	SetInternalValue(val *GoogleBigqueryJobExtract)
	PrintHeader() interface{}
	SetPrintHeader(val interface{})
	PrintHeaderInput() interface{}
	SourceModel() GoogleBigqueryJobExtractSourceModelOutputReference
	SourceModelInput() *GoogleBigqueryJobExtractSourceModel
	SourceTable() GoogleBigqueryJobExtractSourceTableOutputReference
	SourceTableInput() *GoogleBigqueryJobExtractSourceTable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseAvroLogicalTypes() interface{}
	SetUseAvroLogicalTypes(val interface{})
	UseAvroLogicalTypesInput() interface{}
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
	PutSourceModel(value *GoogleBigqueryJobExtractSourceModel)
	PutSourceTable(value *GoogleBigqueryJobExtractSourceTable)
	ResetCompression()
	ResetDestinationFormat()
	ResetFieldDelimiter()
	ResetPrintHeader()
	ResetSourceModel()
	ResetSourceTable()
	ResetUseAvroLogicalTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryJobExtractOutputReference
type jsiiProxy_GoogleBigqueryJobExtractOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) DestinationFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) DestinationFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) DestinationUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) DestinationUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) InternalValue() *GoogleBigqueryJobExtract {
	var returns *GoogleBigqueryJobExtract
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) PrintHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"printHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) PrintHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"printHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) SourceModel() GoogleBigqueryJobExtractSourceModelOutputReference {
	var returns GoogleBigqueryJobExtractSourceModelOutputReference
	_jsii_.Get(
		j,
		"sourceModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) SourceModelInput() *GoogleBigqueryJobExtractSourceModel {
	var returns *GoogleBigqueryJobExtractSourceModel
	_jsii_.Get(
		j,
		"sourceModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) SourceTable() GoogleBigqueryJobExtractSourceTableOutputReference {
	var returns GoogleBigqueryJobExtractSourceTableOutputReference
	_jsii_.Get(
		j,
		"sourceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) SourceTableInput() *GoogleBigqueryJobExtractSourceTable {
	var returns *GoogleBigqueryJobExtractSourceTable
	_jsii_.Get(
		j,
		"sourceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) UseAvroLogicalTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAvroLogicalTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference) UseAvroLogicalTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAvroLogicalTypesInput",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryJobExtractOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryJobExtractOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryJobExtractOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryJobExtractOutputReference{}

	_jsii_.Create(
		"google-beta.googleBigqueryJob.GoogleBigqueryJobExtractOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryJobExtractOutputReference_Override(g GoogleBigqueryJobExtractOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleBigqueryJob.GoogleBigqueryJobExtractOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetDestinationFormat(val *string) {
	if err := j.validateSetDestinationFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetDestinationUris(val *[]*string) {
	if err := j.validateSetDestinationUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationUris",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetInternalValue(val *GoogleBigqueryJobExtract) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetPrintHeader(val interface{}) {
	if err := j.validateSetPrintHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printHeader",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobExtractOutputReference)SetUseAvroLogicalTypes(val interface{}) {
	if err := j.validateSetUseAvroLogicalTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAvroLogicalTypes",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) PutSourceModel(value *GoogleBigqueryJobExtractSourceModel) {
	if err := g.validatePutSourceModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceModel",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) PutSourceTable(value *GoogleBigqueryJobExtractSourceTable) {
	if err := g.validatePutSourceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceTable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		g,
		"resetCompression",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ResetDestinationFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ResetPrintHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetPrintHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ResetSourceModel() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ResetSourceTable() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ResetUseAvroLogicalTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetUseAvroLogicalTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobExtractOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

