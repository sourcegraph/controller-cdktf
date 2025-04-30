package googleeventarcpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleeventarcpipeline/internal"
)

type GoogleEventarcPipelineDestinationsOutputReference interface {
	cdktf.ComplexObject
	AuthenticationConfig() GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference
	AuthenticationConfigInput() *GoogleEventarcPipelineDestinationsAuthenticationConfig
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
	HttpEndpoint() GoogleEventarcPipelineDestinationsHttpEndpointOutputReference
	HttpEndpointInput() *GoogleEventarcPipelineDestinationsHttpEndpoint
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MessageBus() *string
	SetMessageBus(val *string)
	MessageBusInput() *string
	NetworkConfig() GoogleEventarcPipelineDestinationsNetworkConfigOutputReference
	NetworkConfigInput() *GoogleEventarcPipelineDestinationsNetworkConfig
	OutputPayloadFormat() GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference
	OutputPayloadFormatInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormat
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
	Workflow() *string
	SetWorkflow(val *string)
	WorkflowInput() *string
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
	PutAuthenticationConfig(value *GoogleEventarcPipelineDestinationsAuthenticationConfig)
	PutHttpEndpoint(value *GoogleEventarcPipelineDestinationsHttpEndpoint)
	PutNetworkConfig(value *GoogleEventarcPipelineDestinationsNetworkConfig)
	PutOutputPayloadFormat(value *GoogleEventarcPipelineDestinationsOutputPayloadFormat)
	ResetAuthenticationConfig()
	ResetHttpEndpoint()
	ResetMessageBus()
	ResetNetworkConfig()
	ResetOutputPayloadFormat()
	ResetTopic()
	ResetWorkflow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleEventarcPipelineDestinationsOutputReference
type jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) AuthenticationConfig() GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference {
	var returns GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference
	_jsii_.Get(
		j,
		"authenticationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) AuthenticationConfigInput() *GoogleEventarcPipelineDestinationsAuthenticationConfig {
	var returns *GoogleEventarcPipelineDestinationsAuthenticationConfig
	_jsii_.Get(
		j,
		"authenticationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) HttpEndpoint() GoogleEventarcPipelineDestinationsHttpEndpointOutputReference {
	var returns GoogleEventarcPipelineDestinationsHttpEndpointOutputReference
	_jsii_.Get(
		j,
		"httpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) HttpEndpointInput() *GoogleEventarcPipelineDestinationsHttpEndpoint {
	var returns *GoogleEventarcPipelineDestinationsHttpEndpoint
	_jsii_.Get(
		j,
		"httpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) MessageBus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) MessageBusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageBusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) NetworkConfig() GoogleEventarcPipelineDestinationsNetworkConfigOutputReference {
	var returns GoogleEventarcPipelineDestinationsNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) NetworkConfigInput() *GoogleEventarcPipelineDestinationsNetworkConfig {
	var returns *GoogleEventarcPipelineDestinationsNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) OutputPayloadFormat() GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference {
	var returns GoogleEventarcPipelineDestinationsOutputPayloadFormatOutputReference
	_jsii_.Get(
		j,
		"outputPayloadFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) OutputPayloadFormatInput() *GoogleEventarcPipelineDestinationsOutputPayloadFormat {
	var returns *GoogleEventarcPipelineDestinationsOutputPayloadFormat
	_jsii_.Get(
		j,
		"outputPayloadFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) Workflow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) WorkflowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


func NewGoogleEventarcPipelineDestinationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleEventarcPipelineDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleEventarcPipelineDestinationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleEventarcPipelineDestinationsOutputReference_Override(g GoogleEventarcPipelineDestinationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetMessageBus(val *string) {
	if err := j.validateSetMessageBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageBus",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference)SetWorkflow(val *string) {
	if err := j.validateSetWorkflowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflow",
		val,
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) PutAuthenticationConfig(value *GoogleEventarcPipelineDestinationsAuthenticationConfig) {
	if err := g.validatePutAuthenticationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthenticationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) PutHttpEndpoint(value *GoogleEventarcPipelineDestinationsHttpEndpoint) {
	if err := g.validatePutHttpEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpEndpoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) PutNetworkConfig(value *GoogleEventarcPipelineDestinationsNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) PutOutputPayloadFormat(value *GoogleEventarcPipelineDestinationsOutputPayloadFormat) {
	if err := g.validatePutOutputPayloadFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutputPayloadFormat",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ResetAuthenticationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthenticationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ResetHttpEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ResetMessageBus() {
	_jsii_.InvokeVoid(
		g,
		"resetMessageBus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ResetOutputPayloadFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputPayloadFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ResetTopic() {
	_jsii_.InvokeVoid(
		g,
		"resetTopic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ResetWorkflow() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

