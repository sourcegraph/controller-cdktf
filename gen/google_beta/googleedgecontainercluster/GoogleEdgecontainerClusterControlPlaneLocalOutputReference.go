package googleedgecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleedgecontainercluster/internal"
)

type GoogleEdgecontainerClusterControlPlaneLocalOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleEdgecontainerClusterControlPlaneLocal
	SetInternalValue(val *GoogleEdgecontainerClusterControlPlaneLocal)
	MachineFilter() *string
	SetMachineFilter(val *string)
	MachineFilterInput() *string
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeLocation() *string
	SetNodeLocation(val *string)
	NodeLocationInput() *string
	SharedDeploymentPolicy() *string
	SetSharedDeploymentPolicy(val *string)
	SharedDeploymentPolicyInput() *string
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
	ResetMachineFilter()
	ResetNodeCount()
	ResetNodeLocation()
	ResetSharedDeploymentPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleEdgecontainerClusterControlPlaneLocalOutputReference
type jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) InternalValue() *GoogleEdgecontainerClusterControlPlaneLocal {
	var returns *GoogleEdgecontainerClusterControlPlaneLocal
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) MachineFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) MachineFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) NodeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) NodeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) SharedDeploymentPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedDeploymentPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) SharedDeploymentPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedDeploymentPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleEdgecontainerClusterControlPlaneLocalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleEdgecontainerClusterControlPlaneLocalOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleEdgecontainerClusterControlPlaneLocalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEdgecontainerCluster.GoogleEdgecontainerClusterControlPlaneLocalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleEdgecontainerClusterControlPlaneLocalOutputReference_Override(g GoogleEdgecontainerClusterControlPlaneLocalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEdgecontainerCluster.GoogleEdgecontainerClusterControlPlaneLocalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetInternalValue(val *GoogleEdgecontainerClusterControlPlaneLocal) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetMachineFilter(val *string) {
	if err := j.validateSetMachineFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetNodeLocation(val *string) {
	if err := j.validateSetNodeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetSharedDeploymentPolicy(val *string) {
	if err := j.validateSetSharedDeploymentPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedDeploymentPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ResetMachineFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ResetNodeLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ResetSharedDeploymentPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSharedDeploymentPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterControlPlaneLocalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

