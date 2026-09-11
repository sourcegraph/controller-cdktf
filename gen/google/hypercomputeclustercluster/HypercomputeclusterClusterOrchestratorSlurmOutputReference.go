package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterOrchestratorSlurmOutputReference interface {
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
	DefaultPartition() *string
	SetDefaultPartition(val *string)
	DefaultPartitionInput() *string
	EpilogBashScripts() *[]*string
	SetEpilogBashScripts(val *[]*string)
	EpilogBashScriptsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *HypercomputeclusterClusterOrchestratorSlurm
	SetInternalValue(val *HypercomputeclusterClusterOrchestratorSlurm)
	LoginNodes() HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference
	LoginNodesInput() *HypercomputeclusterClusterOrchestratorSlurmLoginNodes
	NodeSets() HypercomputeclusterClusterOrchestratorSlurmNodeSetsList
	NodeSetsInput() interface{}
	Partitions() HypercomputeclusterClusterOrchestratorSlurmPartitionsList
	PartitionsInput() interface{}
	PrologBashScripts() *[]*string
	SetPrologBashScripts(val *[]*string)
	PrologBashScriptsInput() *[]*string
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
	PutLoginNodes(value *HypercomputeclusterClusterOrchestratorSlurmLoginNodes)
	PutNodeSets(value interface{})
	PutPartitions(value interface{})
	ResetDefaultPartition()
	ResetEpilogBashScripts()
	ResetPrologBashScripts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HypercomputeclusterClusterOrchestratorSlurmOutputReference
type jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) DefaultPartition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) DefaultPartitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) EpilogBashScripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"epilogBashScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) EpilogBashScriptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"epilogBashScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) InternalValue() *HypercomputeclusterClusterOrchestratorSlurm {
	var returns *HypercomputeclusterClusterOrchestratorSlurm
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) LoginNodes() HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference {
	var returns HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference
	_jsii_.Get(
		j,
		"loginNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) LoginNodesInput() *HypercomputeclusterClusterOrchestratorSlurmLoginNodes {
	var returns *HypercomputeclusterClusterOrchestratorSlurmLoginNodes
	_jsii_.Get(
		j,
		"loginNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) NodeSets() HypercomputeclusterClusterOrchestratorSlurmNodeSetsList {
	var returns HypercomputeclusterClusterOrchestratorSlurmNodeSetsList
	_jsii_.Get(
		j,
		"nodeSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) NodeSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) Partitions() HypercomputeclusterClusterOrchestratorSlurmPartitionsList {
	var returns HypercomputeclusterClusterOrchestratorSlurmPartitionsList
	_jsii_.Get(
		j,
		"partitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) PartitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) PrologBashScripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prologBashScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) PrologBashScriptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prologBashScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHypercomputeclusterClusterOrchestratorSlurmOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) HypercomputeclusterClusterOrchestratorSlurmOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterOrchestratorSlurmOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterOrchestratorSlurmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterOrchestratorSlurmOutputReference_Override(h HypercomputeclusterClusterOrchestratorSlurmOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterOrchestratorSlurmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetDefaultPartition(val *string) {
	if err := j.validateSetDefaultPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPartition",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetEpilogBashScripts(val *[]*string) {
	if err := j.validateSetEpilogBashScriptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"epilogBashScripts",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetInternalValue(val *HypercomputeclusterClusterOrchestratorSlurm) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetPrologBashScripts(val *[]*string) {
	if err := j.validateSetPrologBashScriptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prologBashScripts",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) PutLoginNodes(value *HypercomputeclusterClusterOrchestratorSlurmLoginNodes) {
	if err := h.validatePutLoginNodesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putLoginNodes",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) PutNodeSets(value interface{}) {
	if err := h.validatePutNodeSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNodeSets",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) PutPartitions(value interface{}) {
	if err := h.validatePutPartitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putPartitions",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) ResetDefaultPartition() {
	_jsii_.InvokeVoid(
		h,
		"resetDefaultPartition",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) ResetEpilogBashScripts() {
	_jsii_.InvokeVoid(
		h,
		"resetEpilogBashScripts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) ResetPrologBashScripts() {
	_jsii_.InvokeVoid(
		h,
		"resetPrologBashScripts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

