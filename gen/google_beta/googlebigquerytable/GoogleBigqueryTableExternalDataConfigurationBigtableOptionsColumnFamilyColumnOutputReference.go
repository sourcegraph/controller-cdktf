package googlebigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerytable/internal"
)

type GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference interface {
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
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	FieldName() *string
	SetFieldName(val *string)
	FieldNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OnlyReadLatest() interface{}
	SetOnlyReadLatest(val interface{})
	OnlyReadLatestInput() interface{}
	QualifierEncoded() *string
	SetQualifierEncoded(val *string)
	QualifierEncodedInput() *string
	QualifierString() *string
	SetQualifierString(val *string)
	QualifierStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetEncoding()
	ResetFieldName()
	ResetOnlyReadLatest()
	ResetQualifierEncoded()
	ResetQualifierString()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference
type jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) OnlyReadLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyReadLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) OnlyReadLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyReadLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierEncoded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierEncoded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierEncodedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierEncodedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference_Override(g GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetOnlyReadLatest(val interface{}) {
	if err := j.validateSetOnlyReadLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyReadLatest",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetQualifierEncoded(val *string) {
	if err := j.validateSetQualifierEncodedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifierEncoded",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetQualifierString(val *string) {
	if err := j.validateSetQualifierStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifierString",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		g,
		"resetEncoding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetFieldName() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetOnlyReadLatest() {
	_jsii_.InvokeVoid(
		g,
		"resetOnlyReadLatest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetQualifierEncoded() {
	_jsii_.InvokeVoid(
		g,
		"resetQualifierEncoded",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetQualifierString() {
	_jsii_.InvokeVoid(
		g,
		"resetQualifierString",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

