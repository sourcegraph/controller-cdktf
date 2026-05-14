package googlehypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehypercomputeclustercluster/internal"
)

type GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference interface {
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
	DefaultPartition() *string
	SetDefaultPartition(val *string)
	DefaultPartitionInput() *string
	EpilogBashScripts() *[]*string
	SetEpilogBashScripts(val *[]*string)
	EpilogBashScriptsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleHypercomputeclusterClusterOrchestratorSlurm
	SetInternalValue(val *GoogleHypercomputeclusterClusterOrchestratorSlurm)
	LoginNodes() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference
	LoginNodesInput() *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes
	NodeSets() GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsList
	NodeSetsInput() interface{}
	Partitions() GoogleHypercomputeclusterClusterOrchestratorSlurmPartitionsList
	PartitionsInput() interface{}
	PrologBashScripts() *[]*string
	SetPrologBashScripts(val *[]*string)
	PrologBashScriptsInput() *[]*string
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
	PutLoginNodes(value *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes)
	PutNodeSets(value interface{})
	PutPartitions(value interface{})
	ResetDefaultPartition()
	ResetEpilogBashScripts()
	ResetPrologBashScripts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference
type jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) DefaultPartition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) DefaultPartitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) EpilogBashScripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"epilogBashScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) EpilogBashScriptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"epilogBashScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) InternalValue() *GoogleHypercomputeclusterClusterOrchestratorSlurm {
	var returns *GoogleHypercomputeclusterClusterOrchestratorSlurm
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) LoginNodes() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference
	_jsii_.Get(
		j,
		"loginNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) LoginNodesInput() *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes {
	var returns *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes
	_jsii_.Get(
		j,
		"loginNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) NodeSets() GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsList {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsList
	_jsii_.Get(
		j,
		"nodeSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) NodeSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) Partitions() GoogleHypercomputeclusterClusterOrchestratorSlurmPartitionsList {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmPartitionsList
	_jsii_.Get(
		j,
		"partitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) PartitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) PrologBashScripts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prologBashScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) PrologBashScriptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prologBashScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHypercomputeclusterClusterOrchestratorSlurmOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference_Override(g GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetDefaultPartition(val *string) {
	if err := j.validateSetDefaultPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPartition",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetEpilogBashScripts(val *[]*string) {
	if err := j.validateSetEpilogBashScriptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"epilogBashScripts",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetInternalValue(val *GoogleHypercomputeclusterClusterOrchestratorSlurm) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetPrologBashScripts(val *[]*string) {
	if err := j.validateSetPrologBashScriptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prologBashScripts",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) PutLoginNodes(value *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes) {
	if err := g.validatePutLoginNodesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoginNodes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) PutNodeSets(value interface{}) {
	if err := g.validatePutNodeSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeSets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) PutPartitions(value interface{}) {
	if err := g.validatePutPartitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartitions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) ResetDefaultPartition() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultPartition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) ResetEpilogBashScripts() {
	_jsii_.InvokeVoid(
		g,
		"resetEpilogBashScripts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) ResetPrologBashScripts() {
	_jsii_.InvokeVoid(
		g,
		"resetPrologBashScripts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

