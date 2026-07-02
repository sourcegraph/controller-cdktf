package cloudsecuritycompliancecloudcontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudsecuritycompliancecloudcontrol/internal"
)

type CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference interface {
	cdktf.ComplexObject
	BoolValue() any
	SetBoolValue(val any)
	BoolValueInput() any
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue
	SetInternalValue(val *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue)
	NumberValue() *float64
	SetNumberValue(val *float64)
	NumberValueInput() *float64
	StringListValue() CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference
	StringListValueInput() *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue
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
	PutStringListValue(value *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue)
	ResetBoolValue()
	ResetNumberValue()
	ResetStringListValue()
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference
type jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) BoolValue() any {
	var returns any
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) BoolValueInput() any {
	var returns any
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InternalValue() *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue {
	var returns *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) NumberValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) NumberValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringListValue() CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValueOutputReference
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringListValueInput() *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue {
	var returns *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue
	_jsii_.Get(
		j,
		"stringListValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference {
	_init_.Initialize()

	if err := validateNewCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference_Override(c CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetBoolValue(val any) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetInternalValue(val *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetNumberValue(val *float64) {
	if err := j.validateSetNumberValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) PutStringListValue(value *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueStringListValue) {
	if err := c.validatePutStringListValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStringListValue",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetBoolValue() {
	_jsii_.InvokeVoid(
		c,
		"resetBoolValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetNumberValue() {
	_jsii_.InvokeVoid(
		c,
		"resetNumberValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetStringListValue() {
	_jsii_.InvokeVoid(
		c,
		"resetStringListValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		c,
		"resetStringValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
