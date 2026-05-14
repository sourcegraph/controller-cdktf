package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterComputeResourcesConfigOutputReference interface {
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
	InternalValue() *HypercomputeclusterClusterComputeResourcesConfig
	SetInternalValue(val *HypercomputeclusterClusterComputeResourcesConfig)
	NewFlexStartInstances() HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstancesOutputReference
	NewFlexStartInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances
	NewOnDemandInstances() HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstancesOutputReference
	NewOnDemandInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances
	NewReservedInstances() HypercomputeclusterClusterComputeResourcesConfigNewReservedInstancesOutputReference
	NewReservedInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewReservedInstances
	NewSpotInstances() HypercomputeclusterClusterComputeResourcesConfigNewSpotInstancesOutputReference
	NewSpotInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewSpotInstances
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
	PutNewFlexStartInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances)
	PutNewOnDemandInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances)
	PutNewReservedInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewReservedInstances)
	PutNewSpotInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewSpotInstances)
	ResetNewFlexStartInstances()
	ResetNewOnDemandInstances()
	ResetNewReservedInstances()
	ResetNewSpotInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HypercomputeclusterClusterComputeResourcesConfigOutputReference
type jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) InternalValue() *HypercomputeclusterClusterComputeResourcesConfig {
	var returns *HypercomputeclusterClusterComputeResourcesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewFlexStartInstances() HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstancesOutputReference {
	var returns HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstancesOutputReference
	_jsii_.Get(
		j,
		"newFlexStartInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewFlexStartInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances {
	var returns *HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances
	_jsii_.Get(
		j,
		"newFlexStartInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewOnDemandInstances() HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstancesOutputReference {
	var returns HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstancesOutputReference
	_jsii_.Get(
		j,
		"newOnDemandInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewOnDemandInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances {
	var returns *HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances
	_jsii_.Get(
		j,
		"newOnDemandInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewReservedInstances() HypercomputeclusterClusterComputeResourcesConfigNewReservedInstancesOutputReference {
	var returns HypercomputeclusterClusterComputeResourcesConfigNewReservedInstancesOutputReference
	_jsii_.Get(
		j,
		"newReservedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewReservedInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewReservedInstances {
	var returns *HypercomputeclusterClusterComputeResourcesConfigNewReservedInstances
	_jsii_.Get(
		j,
		"newReservedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewSpotInstances() HypercomputeclusterClusterComputeResourcesConfigNewSpotInstancesOutputReference {
	var returns HypercomputeclusterClusterComputeResourcesConfigNewSpotInstancesOutputReference
	_jsii_.Get(
		j,
		"newSpotInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) NewSpotInstancesInput() *HypercomputeclusterClusterComputeResourcesConfigNewSpotInstances {
	var returns *HypercomputeclusterClusterComputeResourcesConfigNewSpotInstances
	_jsii_.Get(
		j,
		"newSpotInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHypercomputeclusterClusterComputeResourcesConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HypercomputeclusterClusterComputeResourcesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterComputeResourcesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterComputeResourcesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterComputeResourcesConfigOutputReference_Override(h HypercomputeclusterClusterComputeResourcesConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterComputeResourcesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference)SetInternalValue(val *HypercomputeclusterClusterComputeResourcesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewFlexStartInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances) {
	if err := h.validatePutNewFlexStartInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNewFlexStartInstances",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewOnDemandInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances) {
	if err := h.validatePutNewOnDemandInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNewOnDemandInstances",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewReservedInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewReservedInstances) {
	if err := h.validatePutNewReservedInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNewReservedInstances",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewSpotInstances(value *HypercomputeclusterClusterComputeResourcesConfigNewSpotInstances) {
	if err := h.validatePutNewSpotInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNewSpotInstances",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewFlexStartInstances() {
	_jsii_.InvokeVoid(
		h,
		"resetNewFlexStartInstances",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewOnDemandInstances() {
	_jsii_.InvokeVoid(
		h,
		"resetNewOnDemandInstances",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewReservedInstances() {
	_jsii_.InvokeVoid(
		h,
		"resetNewReservedInstances",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewSpotInstances() {
	_jsii_.InvokeVoid(
		h,
		"resetNewSpotInstances",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterComputeResourcesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

