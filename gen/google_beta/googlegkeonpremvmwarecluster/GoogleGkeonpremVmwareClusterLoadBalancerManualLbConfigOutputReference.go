package googlegkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwarecluster/internal"
)

type GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference interface {
	cdktf.ComplexObject
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
	ControlPlaneNodePort() *float64
	SetControlPlaneNodePort(val *float64)
	ControlPlaneNodePortInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IngressHttpNodePort() *float64
	SetIngressHttpNodePort(val *float64)
	IngressHttpNodePortInput() *float64
	IngressHttpsNodePort() *float64
	SetIngressHttpsNodePort(val *float64)
	IngressHttpsNodePortInput() *float64
	InternalValue() *GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfig
	SetInternalValue(val *GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfig)
	KonnectivityServerNodePort() *float64
	SetKonnectivityServerNodePort(val *float64)
	KonnectivityServerNodePortInput() *float64
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
	ResetControlPlaneNodePort()
	ResetIngressHttpNodePort()
	ResetIngressHttpsNodePort()
	ResetKonnectivityServerNodePort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference
type jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ControlPlaneNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controlPlaneNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ControlPlaneNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controlPlaneNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) IngressHttpNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) IngressHttpNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) IngressHttpsNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpsNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) IngressHttpsNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpsNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) InternalValue() *GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfig {
	var returns *GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) KonnectivityServerNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"konnectivityServerNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) KonnectivityServerNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"konnectivityServerNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference_Override(g GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleGkeonpremVmwareCluster.GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetControlPlaneNodePort(val *float64) {
	if err := j.validateSetControlPlaneNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetIngressHttpNodePort(val *float64) {
	if err := j.validateSetIngressHttpNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressHttpNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetIngressHttpsNodePort(val *float64) {
	if err := j.validateSetIngressHttpsNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressHttpsNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetInternalValue(val *GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetKonnectivityServerNodePort(val *float64) {
	if err := j.validateSetKonnectivityServerNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"konnectivityServerNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ResetControlPlaneNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ResetIngressHttpNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressHttpNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ResetIngressHttpsNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressHttpsNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ResetKonnectivityServerNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetKonnectivityServerNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareClusterLoadBalancerManualLbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

