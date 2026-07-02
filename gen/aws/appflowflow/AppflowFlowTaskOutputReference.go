package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/appflowflow/internal"
)

type AppflowFlowTaskOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectorOperator() AppflowFlowTaskConnectorOperatorList
	ConnectorOperatorInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationField() *string
	SetDestinationField(val *string)
	DestinationFieldInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	SourceFields() *[]*string
	SetSourceFields(val *[]*string)
	SourceFieldsInput() *[]*string
	TaskProperties() *map[string]*string
	SetTaskProperties(val *map[string]*string)
	TaskPropertiesInput() *map[string]*string
	TaskType() *string
	SetTaskType(val *string)
	TaskTypeInput() *string
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutConnectorOperator(value any)
	ResetConnectorOperator()
	ResetDestinationField()
	ResetTaskProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowTaskOutputReference
type jsiiProxy_AppflowFlowTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) ConnectorOperator() AppflowFlowTaskConnectorOperatorList {
	var returns AppflowFlowTaskConnectorOperatorList
	_jsii_.Get(
		j,
		"connectorOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) ConnectorOperatorInput() any {
	var returns any
	_jsii_.Get(
		j,
		"connectorOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) DestinationField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) DestinationFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SourceFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SourceFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) TaskProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"taskProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) TaskPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"taskPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) TaskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppflowFlowTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppflowFlowTaskOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowTaskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppflowFlowTaskOutputReference_Override(a AppflowFlowTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTaskOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetDestinationField(val *string) {
	if err := j.validateSetDestinationFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationField",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetSourceFields(val *[]*string) {
	if err := j.validateSetSourceFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFields",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetTaskProperties(val *map[string]*string) {
	if err := j.validateSetTaskPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskProperties",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetTaskType(val *string) {
	if err := j.validateSetTaskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTaskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) PutConnectorOperator(value any) {
	if err := a.validatePutConnectorOperatorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectorOperator",
		[]any{value},
	)
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) ResetConnectorOperator() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectorOperator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) ResetDestinationField() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) ResetTaskProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
