package firebaseremoteconfigremoteconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/firebaseremoteconfigremoteconfig/internal"
)

type FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference interface {
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
	ConditionalValues() FirebaseRemoteConfigRemoteConfigParameterGroupsParametersConditionalValuesList
	ConditionalValuesInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultValue() FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValueOutputReference
	DefaultValueInput() *FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	ParameterName() *string
	SetParameterName(val *string)
	ParameterNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValueType() *string
	SetValueType(val *string)
	ValueTypeInput() *string
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
	PutConditionalValues(value any)
	PutDefaultValue(value *FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue)
	ResetConditionalValues()
	ResetDefaultValue()
	ResetDescription()
	ResetValueType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference
type jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ConditionalValues() FirebaseRemoteConfigRemoteConfigParameterGroupsParametersConditionalValuesList {
	var returns FirebaseRemoteConfigRemoteConfigParameterGroupsParametersConditionalValuesList
	_jsii_.Get(
		j,
		"conditionalValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ConditionalValuesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"conditionalValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) DefaultValue() FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValueOutputReference {
	var returns FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValueOutputReference
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) DefaultValueInput() *FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue {
	var returns *FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ParameterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueTypeInput",
		&returns,
	)
	return returns
}

func NewFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference {
	_init_.Initialize()

	if err := validateNewFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.firebaseRemoteConfigRemoteConfig.FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference_Override(f FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.firebaseRemoteConfigRemoteConfig.FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetParameterName(val *string) {
	if err := j.validateSetParameterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterName",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) SetValueType(val *string) {
	if err := j.validateSetValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueType",
		val,
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) PutConditionalValues(value any) {
	if err := f.validatePutConditionalValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putConditionalValues",
		[]any{value},
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) PutDefaultValue(value *FirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue) {
	if err := f.validatePutDefaultValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDefaultValue",
		[]any{value},
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetConditionalValues() {
	_jsii_.InvokeVoid(
		f,
		"resetConditionalValues",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetValueType() {
	_jsii_.InvokeVoid(
		f,
		"resetValueType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		f,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
