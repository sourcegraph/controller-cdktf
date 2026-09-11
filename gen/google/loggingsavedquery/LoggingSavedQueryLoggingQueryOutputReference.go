package loggingsavedquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/loggingsavedquery/internal"
)

type LoggingSavedQueryLoggingQueryOutputReference interface {
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
	InternalValue() *LoggingSavedQueryLoggingQuery
	SetInternalValue(val *LoggingSavedQueryLoggingQuery)
	SummaryFieldEnd() *float64
	SetSummaryFieldEnd(val *float64)
	SummaryFieldEndInput() *float64
	SummaryFields() LoggingSavedQueryLoggingQuerySummaryFieldsList
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

// The jsii proxy struct for LoggingSavedQueryLoggingQueryOutputReference
type jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) InternalValue() *LoggingSavedQueryLoggingQuery {
	var returns *LoggingSavedQueryLoggingQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) SummaryFieldEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) SummaryFieldEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) SummaryFields() LoggingSavedQueryLoggingQuerySummaryFieldsList {
	var returns LoggingSavedQueryLoggingQuerySummaryFieldsList
	_jsii_.Get(
		j,
		"summaryFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) SummaryFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"summaryFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) SummaryFieldStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) SummaryFieldStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"summaryFieldStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLoggingSavedQueryLoggingQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LoggingSavedQueryLoggingQueryOutputReference {
	_init_.Initialize()

	if err := validateNewLoggingSavedQueryLoggingQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.loggingSavedQuery.LoggingSavedQueryLoggingQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLoggingSavedQueryLoggingQueryOutputReference_Override(l LoggingSavedQueryLoggingQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.loggingSavedQuery.LoggingSavedQueryLoggingQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetInternalValue(val *LoggingSavedQueryLoggingQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetSummaryFieldEnd(val *float64) {
	if err := j.validateSetSummaryFieldEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summaryFieldEnd",
		val,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetSummaryFieldStart(val *float64) {
	if err := j.validateSetSummaryFieldStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summaryFieldStart",
		val,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) PutSummaryFields(value interface{}) {
	if err := l.validatePutSummaryFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSummaryFields",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) ResetSummaryFieldEnd() {
	_jsii_.InvokeVoid(
		l,
		"resetSummaryFieldEnd",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) ResetSummaryFields() {
	_jsii_.InvokeVoid(
		l,
		"resetSummaryFields",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) ResetSummaryFieldStart() {
	_jsii_.InvokeVoid(
		l,
		"resetSummaryFieldStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingSavedQueryLoggingQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

