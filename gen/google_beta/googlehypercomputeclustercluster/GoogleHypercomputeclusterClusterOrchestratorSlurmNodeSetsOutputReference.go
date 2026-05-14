package googlehypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehypercomputeclustercluster/internal"
)

type GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference interface {
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
	ComputeId() *string
	SetComputeId(val *string)
	ComputeIdInput() *string
	ComputeInstance() GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstanceOutputReference
	ComputeInstanceInput() *GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance
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
	StorageConfigs() GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsStorageConfigsList
	StorageConfigsInput() interface{}
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
	PutComputeInstance(value *GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance)
	PutStorageConfigs(value interface{})
	ResetComputeId()
	ResetComputeInstance()
	ResetMaxDynamicNodeCount()
	ResetStaticNodeCount()
	ResetStorageConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference
type jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeInstance() GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstanceOutputReference {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstanceOutputReference
	_jsii_.Get(
		j,
		"computeInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeInstanceInput() *GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance {
	var returns *GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance
	_jsii_.Get(
		j,
		"computeInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) MaxDynamicNodeCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDynamicNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) MaxDynamicNodeCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDynamicNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StaticNodeCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StaticNodeCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StorageConfigs() GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsStorageConfigsList {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsStorageConfigsList
	_jsii_.Get(
		j,
		"storageConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) StorageConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference_Override(g GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetComputeId(val *string) {
	if err := j.validateSetComputeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeId",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetMaxDynamicNodeCount(val *string) {
	if err := j.validateSetMaxDynamicNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDynamicNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetStaticNodeCount(val *string) {
	if err := j.validateSetStaticNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) PutComputeInstance(value *GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance) {
	if err := g.validatePutComputeInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putComputeInstance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) PutStorageConfigs(value interface{}) {
	if err := g.validatePutStorageConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetComputeId() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetComputeInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetMaxDynamicNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxDynamicNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetStaticNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ResetStorageConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

