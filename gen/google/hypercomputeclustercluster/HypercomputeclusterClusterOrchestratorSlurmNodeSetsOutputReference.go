package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference interface {
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
	ComputeId() *string
	SetComputeId(val *string)
	ComputeIdInput() *string
	ComputeInstance() HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstanceOutputReference
	ComputeInstanceInput() *HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxDynamicNodeCount() *string
	SetMaxDynamicNodeCount(val *string)
	MaxDynamicNodeCountInput() *string
	StaticNodeCount() *string
	SetStaticNodeCount(val *string)
	StaticNodeCountInput() *string
	StorageConfigs() HypercomputeclusterClusterOrchestratorSlurmNodeSetsStorageConfigsList
	StorageConfigsInput() interface{}
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
	PutComputeInstance(value *HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance)
	PutStorageConfigs(value interface{})
	ResetComputeId()
	ResetComputeInstance()
	ResetMaxDynamicNodeCount()
	ResetStaticNodeCount()
	ResetStorageConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference
type jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeInstance() HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstanceOutputReference {
	var returns HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstanceOutputReference
	_jsii_.Get(
		j,
		"computeInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeInstanceInput() *HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance {
	var returns *HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance
	_jsii_.Get(
		j,
		"computeInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) MaxDynamicNodeCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDynamicNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) MaxDynamicNodeCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDynamicNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StaticNodeCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StaticNodeCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StorageConfigs() HypercomputeclusterClusterOrchestratorSlurmNodeSetsStorageConfigsList {
	var returns HypercomputeclusterClusterOrchestratorSlurmNodeSetsStorageConfigsList
	_jsii_.Get(
		j,
		"storageConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StorageConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference_Override(h HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetComputeId(val *string) {
	if err := j.validateSetComputeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeId",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetMaxDynamicNodeCount(val *string) {
	if err := j.validateSetMaxDynamicNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDynamicNodeCount",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetStaticNodeCount(val *string) {
	if err := j.validateSetStaticNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticNodeCount",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) PutComputeInstance(value *HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance) {
	if err := h.validatePutComputeInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putComputeInstance",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) PutStorageConfigs(value interface{}) {
	if err := h.validatePutStorageConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putStorageConfigs",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetComputeId() {
	_jsii_.InvokeVoid(
		h,
		"resetComputeId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetComputeInstance() {
	_jsii_.InvokeVoid(
		h,
		"resetComputeInstance",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetMaxDynamicNodeCount() {
	_jsii_.InvokeVoid(
		h,
		"resetMaxDynamicNodeCount",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetStaticNodeCount() {
	_jsii_.InvokeVoid(
		h,
		"resetStaticNodeCount",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetStorageConfigs() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageConfigs",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := h.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

