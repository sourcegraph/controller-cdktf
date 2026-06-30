package gkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkeonpremvmwarecluster/internal"
)

type GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference interface {
	cdktf.ComplexObject
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
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig
	SetInternalValue(val *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference
type jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) InternalValue() *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig {
	var returns *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference_Override(g GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) SetInternalValue(val *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
