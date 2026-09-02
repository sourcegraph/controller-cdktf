package googleloggingsavedquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleloggingsavedquery/internal"
)

type GoogleLoggingSavedQueryLoggingQueryOutputReference interface {
	cdktn.ComplexObject
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
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleLoggingSavedQueryLoggingQuery
	SetInternalValue(val *GoogleLoggingSavedQueryLoggingQuery)
	SummaryFieldEnd() *float64
	SetSummaryFieldEnd(val *float64)
	SummaryFieldEndInput() *float64
	SummaryFields() GoogleLoggingSavedQueryLoggingQuerySummaryFieldsList
	SummaryFieldsInput() interface{}
	SummaryFieldStart() *float64
	SetSummaryFieldStart(val *float64)
	SummaryFieldStartInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutSummaryFields(value interface{})
	ResetSummaryFieldEnd()
	ResetSummaryFields()
	ResetSummaryFieldStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleLoggingSavedQueryLoggingQueryOutputReference
type jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) InternalValue() *GoogleLoggingSavedQueryLoggingQuery {
	var returns *GoogleLoggingSavedQueryLoggingQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) SummaryFieldEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) SummaryFieldEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) SummaryFields() GoogleLoggingSavedQueryLoggingQuerySummaryFieldsList {
	var returns GoogleLoggingSavedQueryLoggingQuerySummaryFieldsList
	_jsii_.Get(
		j,
		"summaryFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) SummaryFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"summaryFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) SummaryFieldStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) SummaryFieldStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleLoggingSavedQueryLoggingQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleLoggingSavedQueryLoggingQueryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleLoggingSavedQueryLoggingQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleLoggingSavedQuery.GoogleLoggingSavedQueryLoggingQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleLoggingSavedQueryLoggingQueryOutputReference_Override(g GoogleLoggingSavedQueryLoggingQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleLoggingSavedQuery.GoogleLoggingSavedQueryLoggingQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetInternalValue(val *GoogleLoggingSavedQueryLoggingQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetSummaryFieldEnd(val *float64) {
	if err := j.validateSetSummaryFieldEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summaryFieldEnd",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetSummaryFieldStart(val *float64) {
	if err := j.validateSetSummaryFieldStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summaryFieldStart",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) PutSummaryFields(value interface{}) {
	if err := g.validatePutSummaryFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSummaryFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) ResetSummaryFieldEnd() {
	_jsii_.InvokeVoid(
		g,
		"resetSummaryFieldEnd",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) ResetSummaryFields() {
	_jsii_.InvokeVoid(
		g,
		"resetSummaryFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) ResetSummaryFieldStart() {
	_jsii_.InvokeVoid(
		g,
		"resetSummaryFieldStart",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingSavedQueryLoggingQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

