package networkserviceslbedgeextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/networkserviceslbedgeextension/internal"
)

type NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference interface {
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
	FailOpen() interface{}
	SetFailOpen(val interface{})
	FailOpenInput() interface{}
	ForwardHeaders() *[]*string
	SetForwardHeaders(val *[]*string)
	ForwardHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	SupportedEvents() *[]*string
	SetSupportedEvents(val *[]*string)
	SupportedEventsInput() *[]*string
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
	ResetFailOpen()
	ResetForwardHeaders()
	ResetSupportedEvents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference
type jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) FailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) FailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ForwardHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forwardHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ForwardHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forwardHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) SupportedEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) SupportedEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesLbEdgeExtension.NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference_Override(n NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesLbEdgeExtension.NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetFailOpen(val interface{}) {
	if err := j.validateSetFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOpen",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetForwardHeaders(val *[]*string) {
	if err := j.validateSetForwardHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardHeaders",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetSupportedEvents(val *[]*string) {
	if err := j.validateSetSupportedEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedEvents",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ResetFailOpen() {
	_jsii_.InvokeVoid(
		n,
		"resetFailOpen",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ResetForwardHeaders() {
	_jsii_.InvokeVoid(
		n,
		"resetForwardHeaders",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ResetSupportedEvents() {
	_jsii_.InvokeVoid(
		n,
		"resetSupportedEvents",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesLbEdgeExtensionExtensionChainsExtensionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

