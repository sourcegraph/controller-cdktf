package googlevertexaiindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiindex/internal"
)

type GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference interface {
	cdktf.ComplexObject
	BruteForceConfig() GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfigOutputReference
	BruteForceConfigInput() *GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig
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
	InternalValue() *GoogleVertexAiIndexMetadataConfigAlgorithmConfig
	SetInternalValue(val *GoogleVertexAiIndexMetadataConfigAlgorithmConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TreeAhConfig() GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference
	TreeAhConfigInput() *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
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
	PutBruteForceConfig(value *GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig)
	PutTreeAhConfig(value *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig)
	ResetBruteForceConfig()
	ResetTreeAhConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference
type jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) BruteForceConfig() GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfigOutputReference {
	var returns GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfigOutputReference
	_jsii_.Get(
		j,
		"bruteForceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) BruteForceConfigInput() *GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig {
	var returns *GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig
	_jsii_.Get(
		j,
		"bruteForceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) InternalValue() *GoogleVertexAiIndexMetadataConfigAlgorithmConfig {
	var returns *GoogleVertexAiIndexMetadataConfigAlgorithmConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TreeAhConfig() GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference {
	var returns GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference
	_jsii_.Get(
		j,
		"treeAhConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TreeAhConfigInput() *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig {
	var returns *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
	_jsii_.Get(
		j,
		"treeAhConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndex.GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference_Override(g GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndex.GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetInternalValue(val *GoogleVertexAiIndexMetadataConfigAlgorithmConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) PutBruteForceConfig(value *GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig) {
	if err := g.validatePutBruteForceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBruteForceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) PutTreeAhConfig(value *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig) {
	if err := g.validatePutTreeAhConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTreeAhConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ResetBruteForceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBruteForceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ResetTreeAhConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTreeAhConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

