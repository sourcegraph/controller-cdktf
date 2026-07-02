package networkservicesmulticastdomainactivation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/networkservicesmulticastdomainactivation/internal"
)

type NetworkServicesMulticastDomainActivationTrafficSpecOutputReference interface {
	cdktf.ComplexObject
	AggrEgressPps() *string
	SetAggrEgressPps(val *string)
	AggrEgressPpsInput() *string
	AggrIngressPps() *string
	SetAggrIngressPps(val *string)
	AggrIngressPpsInput() *string
	AvgPacketSize() *float64
	SetAvgPacketSize(val *float64)
	AvgPacketSizeInput() *float64
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
	InternalValue() *NetworkServicesMulticastDomainActivationTrafficSpec
	SetInternalValue(val *NetworkServicesMulticastDomainActivationTrafficSpec)
	MaxPerGroupIngressPps() *string
	SetMaxPerGroupIngressPps(val *string)
	MaxPerGroupIngressPpsInput() *string
	MaxPerGroupSubscribers() *string
	SetMaxPerGroupSubscribers(val *string)
	MaxPerGroupSubscribersInput() *string
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
	ResetAggrEgressPps()
	ResetAggrIngressPps()
	ResetAvgPacketSize()
	ResetMaxPerGroupIngressPps()
	ResetMaxPerGroupSubscribers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesMulticastDomainActivationTrafficSpecOutputReference
type jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrEgressPps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrEgressPps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrEgressPpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrEgressPpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrIngressPps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrIngressPps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrIngressPpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrIngressPpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AvgPacketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"avgPacketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AvgPacketSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"avgPacketSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) InternalValue() *NetworkServicesMulticastDomainActivationTrafficSpec {
	var returns *NetworkServicesMulticastDomainActivationTrafficSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupIngressPps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupIngressPps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupIngressPpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupIngressPpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupSubscribers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupSubscribers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupSubscribersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupSubscribersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewNetworkServicesMulticastDomainActivationTrafficSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkServicesMulticastDomainActivationTrafficSpecOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesMulticastDomainActivationTrafficSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesMulticastDomainActivation.NetworkServicesMulticastDomainActivationTrafficSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesMulticastDomainActivationTrafficSpecOutputReference_Override(n NetworkServicesMulticastDomainActivationTrafficSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesMulticastDomainActivation.NetworkServicesMulticastDomainActivationTrafficSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetAggrEgressPps(val *string) {
	if err := j.validateSetAggrEgressPpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggrEgressPps",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetAggrIngressPps(val *string) {
	if err := j.validateSetAggrIngressPpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggrIngressPps",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetAvgPacketSize(val *float64) {
	if err := j.validateSetAvgPacketSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avgPacketSize",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetInternalValue(val *NetworkServicesMulticastDomainActivationTrafficSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetMaxPerGroupIngressPps(val *string) {
	if err := j.validateSetMaxPerGroupIngressPpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPerGroupIngressPps",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetMaxPerGroupSubscribers(val *string) {
	if err := j.validateSetMaxPerGroupSubscribersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPerGroupSubscribers",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetAggrEgressPps() {
	_jsii_.InvokeVoid(
		n,
		"resetAggrEgressPps",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetAggrIngressPps() {
	_jsii_.InvokeVoid(
		n,
		"resetAggrIngressPps",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetAvgPacketSize() {
	_jsii_.InvokeVoid(
		n,
		"resetAvgPacketSize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetMaxPerGroupIngressPps() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxPerGroupIngressPps",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetMaxPerGroupSubscribers() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxPerGroupSubscribers",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		n,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
