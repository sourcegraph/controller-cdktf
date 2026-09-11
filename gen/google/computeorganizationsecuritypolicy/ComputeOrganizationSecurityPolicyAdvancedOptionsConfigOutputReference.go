package computeorganizationsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeorganizationsecuritypolicy/internal"
)

type ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeOrganizationSecurityPolicyAdvancedOptionsConfig
	SetInternalValue(val *ComputeOrganizationSecurityPolicyAdvancedOptionsConfig)
	JsonCustomConfig() ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	JsonCustomConfigInput() *ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	JsonParsing() *string
	SetJsonParsing(val *string)
	JsonParsingInput() *string
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	RequestBodyInspectionSize() *string
	SetRequestBodyInspectionSize(val *string)
	RequestBodyInspectionSizeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserIpRequestHeaders() *[]*string
	SetUserIpRequestHeaders(val *[]*string)
	UserIpRequestHeadersInput() *[]*string
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
	PutJsonCustomConfig(value *ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig)
	ResetJsonCustomConfig()
	ResetJsonParsing()
	ResetLogLevel()
	ResetRequestBodyInspectionSize()
	ResetUserIpRequestHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference
type jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) InternalValue() *ComputeOrganizationSecurityPolicyAdvancedOptionsConfig {
	var returns *ComputeOrganizationSecurityPolicyAdvancedOptionsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfig() ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference {
	var returns ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	_jsii_.Get(
		j,
		"jsonCustomConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfigInput() *ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig {
	var returns *ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	_jsii_.Get(
		j,
		"jsonCustomConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) RequestBodyInspectionSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInspectionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) RequestBodyInspectionSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInspectionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeadersInput",
		&returns,
	)
	return returns
}


func NewComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeOrganizationSecurityPolicy.ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference_Override(c ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeOrganizationSecurityPolicy.ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetInternalValue(val *ComputeOrganizationSecurityPolicyAdvancedOptionsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetJsonParsing(val *string) {
	if err := j.validateSetJsonParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonParsing",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetRequestBodyInspectionSize(val *string) {
	if err := j.validateSetRequestBodyInspectionSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBodyInspectionSize",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference)SetUserIpRequestHeaders(val *[]*string) {
	if err := j.validateSetUserIpRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIpRequestHeaders",
		val,
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) PutJsonCustomConfig(value *ComputeOrganizationSecurityPolicyAdvancedOptionsConfigJsonCustomConfig) {
	if err := c.validatePutJsonCustomConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putJsonCustomConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonCustomConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetJsonCustomConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonParsing() {
	_jsii_.InvokeVoid(
		c,
		"resetJsonParsing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ResetRequestBodyInspectionSize() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestBodyInspectionSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ResetUserIpRequestHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetUserIpRequestHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyAdvancedOptionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

