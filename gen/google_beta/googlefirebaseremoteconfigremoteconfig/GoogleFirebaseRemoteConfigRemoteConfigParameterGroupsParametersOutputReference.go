package googlefirebaseremoteconfigremoteconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlefirebaseremoteconfigremoteconfig/internal"
)

type GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference interface {
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
	ConditionalValues() GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersConditionalValuesList
	ConditionalValuesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultValue() GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValueOutputReference
	DefaultValueInput() *GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ParameterName() *string
	SetParameterName(val *string)
	ParameterNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ValueType() *string
	SetValueType(val *string)
	ValueTypeInput() *string
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
	PutConditionalValues(value interface{})
	PutDefaultValue(value *GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue)
	ResetConditionalValues()
	ResetDefaultValue()
	ResetDescription()
	ResetValueType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference
type jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ConditionalValues() GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersConditionalValuesList {
	var returns GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersConditionalValuesList
	_jsii_.Get(
		j,
		"conditionalValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ConditionalValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) DefaultValue() GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValueOutputReference {
	var returns GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValueOutputReference
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) DefaultValueInput() *GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue {
	var returns *GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ParameterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueTypeInput",
		&returns,
	)
	return returns
}


func NewGoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirebaseRemoteConfigRemoteConfig.GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference_Override(g GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirebaseRemoteConfigRemoteConfig.GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetParameterName(val *string) {
	if err := j.validateSetParameterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterName",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference)SetValueType(val *string) {
	if err := j.validateSetValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueType",
		val,
	)
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) PutConditionalValues(value interface{}) {
	if err := g.validatePutConditionalValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditionalValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) PutDefaultValue(value *GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue) {
	if err := g.validatePutDefaultValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetConditionalValues() {
	_jsii_.InvokeVoid(
		g,
		"resetConditionalValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ResetValueType() {
	_jsii_.InvokeVoid(
		g,
		"resetValueType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

