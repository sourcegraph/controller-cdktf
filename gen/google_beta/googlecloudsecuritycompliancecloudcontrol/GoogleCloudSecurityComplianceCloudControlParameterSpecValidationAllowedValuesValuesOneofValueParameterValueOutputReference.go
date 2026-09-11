package googlecloudsecuritycompliancecloudcontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudsecuritycompliancecloudcontrol/internal"
)

type GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference interface {
	cdktn.ComplexObject
	BoolValue() interface{}
	SetBoolValue(val interface{})
	BoolValueInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue
	SetInternalValue(val *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue)
	NumberValue() *float64
	SetNumberValue(val *float64)
	NumberValueInput() *float64
	StringListValue() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference
	StringListValueInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue
	StringValue() *string
	SetStringValue(val *string)
	StringValueInput() *string
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
	PutStringListValue(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue)
	ResetBoolValue()
	ResetNumberValue()
	ResetStringListValue()
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference
type jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) BoolValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) BoolValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InternalValue() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) NumberValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) NumberValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringListValue() GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference {
	var returns GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringListValueInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue
	_jsii_.Get(
		j,
		"stringListValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudSecurityComplianceCloudControl.GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference_Override(g GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudSecurityComplianceCloudControl.GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetBoolValue(val interface{}) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetInternalValue(val *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetNumberValue(val *float64) {
	if err := j.validateSetNumberValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) PutStringListValue(value *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue) {
	if err := g.validatePutStringListValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStringListValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetBoolValue() {
	_jsii_.InvokeVoid(
		g,
		"resetBoolValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetNumberValue() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetStringListValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStringListValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStringValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

