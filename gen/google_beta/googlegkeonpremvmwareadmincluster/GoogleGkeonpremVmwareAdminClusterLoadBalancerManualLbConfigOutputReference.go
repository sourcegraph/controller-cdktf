package googlegkeonpremvmwareadmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwareadmincluster/internal"
)

type GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference interface {
	cdktn.ComplexObject
	AddonsNodePort() *float64
	SetAddonsNodePort(val *float64)
	AddonsNodePortInput() *float64
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
	InternalValue() *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig
	SetInternalValue(val *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig)
	KonnectivityServerNodePort() *float64
	SetKonnectivityServerNodePort(val *float64)
	KonnectivityServerNodePortInput() *float64
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
	ResetAddonsNodePort()
	ResetControlPlaneNodePort()
	ResetIngressHttpNodePort()
	ResetIngressHttpsNodePort()
	ResetKonnectivityServerNodePort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference
type jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) AddonsNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addonsNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) AddonsNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addonsNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ControlPlaneNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controlPlaneNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ControlPlaneNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controlPlaneNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpsNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpsNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpsNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpsNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) InternalValue() *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig {
	var returns *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) KonnectivityServerNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"konnectivityServerNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) KonnectivityServerNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"konnectivityServerNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference_Override(g GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetAddonsNodePort(val *float64) {
	if err := j.validateSetAddonsNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addonsNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetControlPlaneNodePort(val *float64) {
	if err := j.validateSetControlPlaneNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetIngressHttpNodePort(val *float64) {
	if err := j.validateSetIngressHttpNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressHttpNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetIngressHttpsNodePort(val *float64) {
	if err := j.validateSetIngressHttpsNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressHttpsNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetInternalValue(val *GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetKonnectivityServerNodePort(val *float64) {
	if err := j.validateSetKonnectivityServerNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"konnectivityServerNodePort",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetAddonsNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetAddonsNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetControlPlaneNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetIngressHttpNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressHttpNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetIngressHttpsNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressHttpsNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetKonnectivityServerNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetKonnectivityServerNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

