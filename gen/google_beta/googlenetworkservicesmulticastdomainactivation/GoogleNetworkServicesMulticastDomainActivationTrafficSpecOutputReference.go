package googlenetworkservicesmulticastdomainactivation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicesmulticastdomainactivation/internal"
)

type GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference interface {
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
	InternalValue() *GoogleNetworkServicesMulticastDomainActivationTrafficSpec
	SetInternalValue(val *GoogleNetworkServicesMulticastDomainActivationTrafficSpec)
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
	ResetAggrEgressPps()
	ResetAggrIngressPps()
	ResetAvgPacketSize()
	ResetMaxPerGroupIngressPps()
	ResetMaxPerGroupSubscribers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference
type jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrEgressPps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrEgressPps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrEgressPpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrEgressPpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrIngressPps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrIngressPps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AggrIngressPpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggrIngressPpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AvgPacketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"avgPacketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) AvgPacketSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"avgPacketSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) InternalValue() *GoogleNetworkServicesMulticastDomainActivationTrafficSpec {
	var returns *GoogleNetworkServicesMulticastDomainActivationTrafficSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupIngressPps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupIngressPps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupIngressPpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupIngressPpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupSubscribers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupSubscribers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) MaxPerGroupSubscribersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPerGroupSubscribersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesMulticastDomainActivation.GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference_Override(g GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesMulticastDomainActivation.GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetAggrEgressPps(val *string) {
	if err := j.validateSetAggrEgressPpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggrEgressPps",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetAggrIngressPps(val *string) {
	if err := j.validateSetAggrIngressPpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggrIngressPps",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetAvgPacketSize(val *float64) {
	if err := j.validateSetAvgPacketSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avgPacketSize",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetInternalValue(val *GoogleNetworkServicesMulticastDomainActivationTrafficSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetMaxPerGroupIngressPps(val *string) {
	if err := j.validateSetMaxPerGroupIngressPpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPerGroupIngressPps",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetMaxPerGroupSubscribers(val *string) {
	if err := j.validateSetMaxPerGroupSubscribersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPerGroupSubscribers",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetAggrEgressPps() {
	_jsii_.InvokeVoid(
		g,
		"resetAggrEgressPps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetAggrIngressPps() {
	_jsii_.InvokeVoid(
		g,
		"resetAggrIngressPps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetAvgPacketSize() {
	_jsii_.InvokeVoid(
		g,
		"resetAvgPacketSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetMaxPerGroupIngressPps() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxPerGroupIngressPps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ResetMaxPerGroupSubscribers() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxPerGroupSubscribers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesMulticastDomainActivationTrafficSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

