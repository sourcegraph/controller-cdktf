package googlevertexaiindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiindex/internal"
)

type GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference interface {
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
	InternalValue() *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
	SetInternalValue(val *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig)
	LeafNodeEmbeddingCount() *float64
	SetLeafNodeEmbeddingCount(val *float64)
	LeafNodeEmbeddingCountInput() *float64
	LeafNodesToSearchPercent() *float64
	SetLeafNodesToSearchPercent(val *float64)
	LeafNodesToSearchPercentInput() *float64
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
	ResetLeafNodeEmbeddingCount()
	ResetLeafNodesToSearchPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference
type jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) InternalValue() *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig {
	var returns *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodeEmbeddingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodeEmbeddingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodeEmbeddingCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodeEmbeddingCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodesToSearchPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodesToSearchPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodesToSearchPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodesToSearchPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndex.GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference_Override(g GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndex.GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetInternalValue(val *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetLeafNodeEmbeddingCount(val *float64) {
	if err := j.validateSetLeafNodeEmbeddingCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leafNodeEmbeddingCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetLeafNodesToSearchPercent(val *float64) {
	if err := j.validateSetLeafNodesToSearchPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leafNodesToSearchPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ResetLeafNodeEmbeddingCount() {
	_jsii_.InvokeVoid(
		g,
		"resetLeafNodeEmbeddingCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ResetLeafNodesToSearchPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetLeafNodesToSearchPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

