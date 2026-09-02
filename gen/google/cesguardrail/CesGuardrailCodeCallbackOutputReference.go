package cesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesguardrail/internal"
)

type CesGuardrailCodeCallbackOutputReference interface {
	cdktn.ComplexObject
	AfterAgentCallback() CesGuardrailCodeCallbackAfterAgentCallbackOutputReference
	AfterAgentCallbackInput() *CesGuardrailCodeCallbackAfterAgentCallback
	AfterModelCallback() CesGuardrailCodeCallbackAfterModelCallbackOutputReference
	AfterModelCallbackInput() *CesGuardrailCodeCallbackAfterModelCallback
	BeforeAgentCallback() CesGuardrailCodeCallbackBeforeAgentCallbackOutputReference
	BeforeAgentCallbackInput() *CesGuardrailCodeCallbackBeforeAgentCallback
	BeforeModelCallback() CesGuardrailCodeCallbackBeforeModelCallbackOutputReference
	BeforeModelCallbackInput() *CesGuardrailCodeCallbackBeforeModelCallback
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
	InternalValue() *CesGuardrailCodeCallback
	SetInternalValue(val *CesGuardrailCodeCallback)
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
	PutAfterAgentCallback(value *CesGuardrailCodeCallbackAfterAgentCallback)
	PutAfterModelCallback(value *CesGuardrailCodeCallbackAfterModelCallback)
	PutBeforeAgentCallback(value *CesGuardrailCodeCallbackBeforeAgentCallback)
	PutBeforeModelCallback(value *CesGuardrailCodeCallbackBeforeModelCallback)
	ResetAfterAgentCallback()
	ResetAfterModelCallback()
	ResetBeforeAgentCallback()
	ResetBeforeModelCallback()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesGuardrailCodeCallbackOutputReference
type jsiiProxy_CesGuardrailCodeCallbackOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) AfterAgentCallback() CesGuardrailCodeCallbackAfterAgentCallbackOutputReference {
	var returns CesGuardrailCodeCallbackAfterAgentCallbackOutputReference
	_jsii_.Get(
		j,
		"afterAgentCallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) AfterAgentCallbackInput() *CesGuardrailCodeCallbackAfterAgentCallback {
	var returns *CesGuardrailCodeCallbackAfterAgentCallback
	_jsii_.Get(
		j,
		"afterAgentCallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) AfterModelCallback() CesGuardrailCodeCallbackAfterModelCallbackOutputReference {
	var returns CesGuardrailCodeCallbackAfterModelCallbackOutputReference
	_jsii_.Get(
		j,
		"afterModelCallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) AfterModelCallbackInput() *CesGuardrailCodeCallbackAfterModelCallback {
	var returns *CesGuardrailCodeCallbackAfterModelCallback
	_jsii_.Get(
		j,
		"afterModelCallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) BeforeAgentCallback() CesGuardrailCodeCallbackBeforeAgentCallbackOutputReference {
	var returns CesGuardrailCodeCallbackBeforeAgentCallbackOutputReference
	_jsii_.Get(
		j,
		"beforeAgentCallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) BeforeAgentCallbackInput() *CesGuardrailCodeCallbackBeforeAgentCallback {
	var returns *CesGuardrailCodeCallbackBeforeAgentCallback
	_jsii_.Get(
		j,
		"beforeAgentCallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) BeforeModelCallback() CesGuardrailCodeCallbackBeforeModelCallbackOutputReference {
	var returns CesGuardrailCodeCallbackBeforeModelCallbackOutputReference
	_jsii_.Get(
		j,
		"beforeModelCallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) BeforeModelCallbackInput() *CesGuardrailCodeCallbackBeforeModelCallback {
	var returns *CesGuardrailCodeCallbackBeforeModelCallback
	_jsii_.Get(
		j,
		"beforeModelCallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) InternalValue() *CesGuardrailCodeCallback {
	var returns *CesGuardrailCodeCallback
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesGuardrailCodeCallbackOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesGuardrailCodeCallbackOutputReference {
	_init_.Initialize()

	if err := validateNewCesGuardrailCodeCallbackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesGuardrailCodeCallbackOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesGuardrail.CesGuardrailCodeCallbackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesGuardrailCodeCallbackOutputReference_Override(c CesGuardrailCodeCallbackOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesGuardrail.CesGuardrailCodeCallbackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference)SetInternalValue(val *CesGuardrailCodeCallback) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailCodeCallbackOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) PutAfterAgentCallback(value *CesGuardrailCodeCallbackAfterAgentCallback) {
	if err := c.validatePutAfterAgentCallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAfterAgentCallback",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) PutAfterModelCallback(value *CesGuardrailCodeCallbackAfterModelCallback) {
	if err := c.validatePutAfterModelCallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAfterModelCallback",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) PutBeforeAgentCallback(value *CesGuardrailCodeCallbackBeforeAgentCallback) {
	if err := c.validatePutBeforeAgentCallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBeforeAgentCallback",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) PutBeforeModelCallback(value *CesGuardrailCodeCallbackBeforeModelCallback) {
	if err := c.validatePutBeforeModelCallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBeforeModelCallback",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ResetAfterAgentCallback() {
	_jsii_.InvokeVoid(
		c,
		"resetAfterAgentCallback",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ResetAfterModelCallback() {
	_jsii_.InvokeVoid(
		c,
		"resetAfterModelCallback",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ResetBeforeAgentCallback() {
	_jsii_.InvokeVoid(
		c,
		"resetBeforeAgentCallback",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ResetBeforeModelCallback() {
	_jsii_.InvokeVoid(
		c,
		"resetBeforeModelCallback",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesGuardrailCodeCallbackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

