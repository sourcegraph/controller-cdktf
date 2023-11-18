package googledataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplexdatascan/internal"
)

type GoogleDataplexDatascanDataProfileSpecOutputReference interface {
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
	ExcludeFields() GoogleDataplexDatascanDataProfileSpecExcludeFieldsOutputReference
	ExcludeFieldsInput() *GoogleDataplexDatascanDataProfileSpecExcludeFields
	// Experimental.
	Fqn() *string
	IncludeFields() GoogleDataplexDatascanDataProfileSpecIncludeFieldsOutputReference
	IncludeFieldsInput() *GoogleDataplexDatascanDataProfileSpecIncludeFields
	InternalValue() *GoogleDataplexDatascanDataProfileSpec
	SetInternalValue(val *GoogleDataplexDatascanDataProfileSpec)
	PostScanActions() GoogleDataplexDatascanDataProfileSpecPostScanActionsOutputReference
	PostScanActionsInput() *GoogleDataplexDatascanDataProfileSpecPostScanActions
	RowFilter() *string
	SetRowFilter(val *string)
	RowFilterInput() *string
	SamplingPercent() *float64
	SetSamplingPercent(val *float64)
	SamplingPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutExcludeFields(value *GoogleDataplexDatascanDataProfileSpecExcludeFields)
	PutIncludeFields(value *GoogleDataplexDatascanDataProfileSpecIncludeFields)
	PutPostScanActions(value *GoogleDataplexDatascanDataProfileSpecPostScanActions)
	ResetExcludeFields()
	ResetIncludeFields()
	ResetPostScanActions()
	ResetRowFilter()
	ResetSamplingPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexDatascanDataProfileSpecOutputReference
type jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ExcludeFields() GoogleDataplexDatascanDataProfileSpecExcludeFieldsOutputReference {
	var returns GoogleDataplexDatascanDataProfileSpecExcludeFieldsOutputReference
	_jsii_.Get(
		j,
		"excludeFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ExcludeFieldsInput() *GoogleDataplexDatascanDataProfileSpecExcludeFields {
	var returns *GoogleDataplexDatascanDataProfileSpecExcludeFields
	_jsii_.Get(
		j,
		"excludeFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) IncludeFields() GoogleDataplexDatascanDataProfileSpecIncludeFieldsOutputReference {
	var returns GoogleDataplexDatascanDataProfileSpecIncludeFieldsOutputReference
	_jsii_.Get(
		j,
		"includeFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) IncludeFieldsInput() *GoogleDataplexDatascanDataProfileSpecIncludeFields {
	var returns *GoogleDataplexDatascanDataProfileSpecIncludeFields
	_jsii_.Get(
		j,
		"includeFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) InternalValue() *GoogleDataplexDatascanDataProfileSpec {
	var returns *GoogleDataplexDatascanDataProfileSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) PostScanActions() GoogleDataplexDatascanDataProfileSpecPostScanActionsOutputReference {
	var returns GoogleDataplexDatascanDataProfileSpecPostScanActionsOutputReference
	_jsii_.Get(
		j,
		"postScanActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) PostScanActionsInput() *GoogleDataplexDatascanDataProfileSpecPostScanActions {
	var returns *GoogleDataplexDatascanDataProfileSpecPostScanActions
	_jsii_.Get(
		j,
		"postScanActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) RowFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) RowFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) SamplingPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) SamplingPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexDatascanDataProfileSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexDatascanDataProfileSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexDatascanDataProfileSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataProfileSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexDatascanDataProfileSpecOutputReference_Override(g GoogleDataplexDatascanDataProfileSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataProfileSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference)SetInternalValue(val *GoogleDataplexDatascanDataProfileSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference)SetRowFilter(val *string) {
	if err := j.validateSetRowFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference)SetSamplingPercent(val *float64) {
	if err := j.validateSetSamplingPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) PutExcludeFields(value *GoogleDataplexDatascanDataProfileSpecExcludeFields) {
	if err := g.validatePutExcludeFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) PutIncludeFields(value *GoogleDataplexDatascanDataProfileSpecIncludeFields) {
	if err := g.validatePutIncludeFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludeFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) PutPostScanActions(value *GoogleDataplexDatascanDataProfileSpecPostScanActions) {
	if err := g.validatePutPostScanActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostScanActions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ResetExcludeFields() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ResetIncludeFields() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ResetPostScanActions() {
	_jsii_.InvokeVoid(
		g,
		"resetPostScanActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ResetRowFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetRowFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ResetSamplingPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetSamplingPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataProfileSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

