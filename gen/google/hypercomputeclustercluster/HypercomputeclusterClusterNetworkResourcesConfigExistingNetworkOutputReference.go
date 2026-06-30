package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *HypercomputeclusterClusterNetworkResourcesConfigExistingNetwork
	SetInternalValue(val *HypercomputeclusterClusterNetworkResourcesConfigExistingNetwork)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
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

// The jsii proxy struct for HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference
type jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) InternalValue() *HypercomputeclusterClusterNetworkResourcesConfigExistingNetwork {
	var returns *HypercomputeclusterClusterNetworkResourcesConfigExistingNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewHypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference_Override(h HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference",
		[]any{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SetInternalValue(val *HypercomputeclusterClusterNetworkResourcesConfigExistingNetwork) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		h,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterNetworkResourcesConfigExistingNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
