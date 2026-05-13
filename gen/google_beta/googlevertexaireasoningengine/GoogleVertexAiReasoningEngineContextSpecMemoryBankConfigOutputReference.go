package googlevertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaireasoningengine/internal"
)

type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference interface {
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
	DisableMemoryRevisions() interface{}
	SetDisableMemoryRevisions(val interface{})
	DisableMemoryRevisionsInput() interface{}
	// Experimental.
	Fqn() *string
	GenerationConfig() GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigOutputReference
	GenerationConfigInput() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig
	InternalValue() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfig
	SetInternalValue(val *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfig)
	SimilaritySearchConfig() GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfigOutputReference
	SimilaritySearchConfigInput() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TtlConfig() GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigOutputReference
	TtlConfigInput() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig
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
	PutGenerationConfig(value *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig)
	PutSimilaritySearchConfig(value *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig)
	PutTtlConfig(value *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig)
	ResetDisableMemoryRevisions()
	ResetGenerationConfig()
	ResetSimilaritySearchConfig()
	ResetTtlConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference
type jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) DisableMemoryRevisions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableMemoryRevisions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) DisableMemoryRevisionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableMemoryRevisionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GenerationConfig() GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigOutputReference {
	var returns GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigOutputReference
	_jsii_.Get(
		j,
		"generationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GenerationConfigInput() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig {
	var returns *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig
	_jsii_.Get(
		j,
		"generationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) InternalValue() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfig {
	var returns *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) SimilaritySearchConfig() GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfigOutputReference {
	var returns GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfigOutputReference
	_jsii_.Get(
		j,
		"similaritySearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) SimilaritySearchConfigInput() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig {
	var returns *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig
	_jsii_.Get(
		j,
		"similaritySearchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TtlConfig() GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigOutputReference {
	var returns GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigOutputReference
	_jsii_.Get(
		j,
		"ttlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TtlConfigInput() *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig {
	var returns *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig
	_jsii_.Get(
		j,
		"ttlConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference_Override(g GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetDisableMemoryRevisions(val interface{}) {
	if err := j.validateSetDisableMemoryRevisionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableMemoryRevisions",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetInternalValue(val *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutGenerationConfig(value *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig) {
	if err := g.validatePutGenerationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGenerationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutSimilaritySearchConfig(value *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig) {
	if err := g.validatePutSimilaritySearchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSimilaritySearchConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutTtlConfig(value *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig) {
	if err := g.validatePutTtlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTtlConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetDisableMemoryRevisions() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableMemoryRevisions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetGenerationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGenerationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetSimilaritySearchConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSimilaritySearchConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetTtlConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTtlConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

