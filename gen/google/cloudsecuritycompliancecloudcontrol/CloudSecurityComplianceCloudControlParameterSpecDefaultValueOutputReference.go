package cloudsecuritycompliancecloudcontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudsecuritycompliancecloudcontrol/internal"
)

type CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference interface {
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
	InternalValue() *CloudSecurityComplianceCloudControlParameterSpecDefaultValue
	SetInternalValue(val *CloudSecurityComplianceCloudControlParameterSpecDefaultValue)
	NumberValue() *float64
	SetNumberValue(val *float64)
	NumberValueInput() *float64
	OneofValue() CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValueOutputReference
	OneofValueInput() *CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValue
	StringListValue() CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValueOutputReference
	StringListValueInput() *CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValue
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
	PutOneofValue(value *CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValue)
	PutStringListValue(value *CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValue)
	ResetBoolValue()
	ResetNumberValue()
	ResetOneofValue()
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

// The jsii proxy struct for CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference
type jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) BoolValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) BoolValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) InternalValue() *CloudSecurityComplianceCloudControlParameterSpecDefaultValue {
	var returns *CloudSecurityComplianceCloudControlParameterSpecDefaultValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) NumberValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) NumberValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) OneofValue() CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValueOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValueOutputReference
	_jsii_.Get(
		j,
		"oneofValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) OneofValueInput() *CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValue {
	var returns *CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValue
	_jsii_.Get(
		j,
		"oneofValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) StringListValue() CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValueOutputReference {
	var returns CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValueOutputReference
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) StringListValueInput() *CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValue {
	var returns *CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValue
	_jsii_.Get(
		j,
		"stringListValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference {
	_init_.Initialize()

	if err := validateNewCloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference_Override(c CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cloudSecurityComplianceCloudControl.CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetBoolValue(val interface{}) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetInternalValue(val *CloudSecurityComplianceCloudControlParameterSpecDefaultValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetNumberValue(val *float64) {
	if err := j.validateSetNumberValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) PutOneofValue(value *CloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValue) {
	if err := c.validatePutOneofValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOneofValue",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) PutStringListValue(value *CloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValue) {
	if err := c.validatePutStringListValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStringListValue",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ResetBoolValue() {
	_jsii_.InvokeVoid(
		c,
		"resetBoolValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ResetNumberValue() {
	_jsii_.InvokeVoid(
		c,
		"resetNumberValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ResetOneofValue() {
	_jsii_.InvokeVoid(
		c,
		"resetOneofValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ResetStringListValue() {
	_jsii_.InvokeVoid(
		c,
		"resetStringListValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		c,
		"resetStringValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceCloudControlParameterSpecDefaultValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

