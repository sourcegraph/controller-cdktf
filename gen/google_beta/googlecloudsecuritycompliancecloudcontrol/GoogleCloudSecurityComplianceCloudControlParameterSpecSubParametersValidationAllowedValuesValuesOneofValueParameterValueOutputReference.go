package googlecloudsecuritycompliancecloudcontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudsecuritycompliancecloudcontrol/internal"
)

type GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValue
	SetInternalValue(val *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValue)
	NumberValue() *float64
	SetNumberValue(val *float64)
	NumberValueInput() *float64
	StringListValue() GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference
	StringListValueInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValue
	StringValue() *string
	SetStringValue(val *string)
	StringValueInput() *string
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
	PutStringListValue(value *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValue)
	ResetBoolValue()
	ResetNumberValue()
	ResetStringListValue()
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference
type jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) BoolValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) BoolValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InternalValue() *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValue {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) NumberValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) NumberValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringListValue() GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference {
	var returns GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringListValueInput() *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValue {
	var returns *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValue
	_jsii_.Get(
		j,
		"stringListValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudSecurityComplianceCloudControl.GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference_Override(g GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudSecurityComplianceCloudControl.GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetBoolValue(val interface{}) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetInternalValue(val *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetNumberValue(val *float64) {
	if err := j.validateSetNumberValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) PutStringListValue(value *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueStringListValue) {
	if err := g.validatePutStringListValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStringListValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetBoolValue() {
	_jsii_.InvokeVoid(
		g,
		"resetBoolValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetNumberValue() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetStringListValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStringListValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStringValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

