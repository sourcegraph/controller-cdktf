package googlevertexaiindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiindex/internal"
)

type GoogleVertexAiIndexMetadataConfigOutputReference interface {
	cdktf.ComplexObject
	AlgorithmConfig() GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference
	AlgorithmConfigInput() *GoogleVertexAiIndexMetadataConfigAlgorithmConfig
	ApproximateNeighborsCount() *float64
	SetApproximateNeighborsCount(val *float64)
	ApproximateNeighborsCountInput() *float64
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
	Dimensions() *float64
	SetDimensions(val *float64)
	DimensionsInput() *float64
	DistanceMeasureType() *string
	SetDistanceMeasureType(val *string)
	DistanceMeasureTypeInput() *string
	FeatureNormType() *string
	SetFeatureNormType(val *string)
	FeatureNormTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiIndexMetadataConfig
	SetInternalValue(val *GoogleVertexAiIndexMetadataConfig)
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
	PutAlgorithmConfig(value *GoogleVertexAiIndexMetadataConfigAlgorithmConfig)
	ResetAlgorithmConfig()
	ResetApproximateNeighborsCount()
	ResetDistanceMeasureType()
	ResetFeatureNormType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiIndexMetadataConfigOutputReference
type jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) AlgorithmConfig() GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference {
	var returns GoogleVertexAiIndexMetadataConfigAlgorithmConfigOutputReference
	_jsii_.Get(
		j,
		"algorithmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) AlgorithmConfigInput() *GoogleVertexAiIndexMetadataConfigAlgorithmConfig {
	var returns *GoogleVertexAiIndexMetadataConfigAlgorithmConfig
	_jsii_.Get(
		j,
		"algorithmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ApproximateNeighborsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approximateNeighborsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ApproximateNeighborsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approximateNeighborsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) Dimensions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) DimensionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) DistanceMeasureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceMeasureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) DistanceMeasureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceMeasureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) FeatureNormType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureNormType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) FeatureNormTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureNormTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) InternalValue() *GoogleVertexAiIndexMetadataConfig {
	var returns *GoogleVertexAiIndexMetadataConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiIndexMetadataConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiIndexMetadataConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiIndexMetadataConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndex.GoogleVertexAiIndexMetadataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiIndexMetadataConfigOutputReference_Override(g GoogleVertexAiIndexMetadataConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiIndex.GoogleVertexAiIndexMetadataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetApproximateNeighborsCount(val *float64) {
	if err := j.validateSetApproximateNeighborsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approximateNeighborsCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetDimensions(val *float64) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetDistanceMeasureType(val *string) {
	if err := j.validateSetDistanceMeasureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distanceMeasureType",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetFeatureNormType(val *string) {
	if err := j.validateSetFeatureNormTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureNormType",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetInternalValue(val *GoogleVertexAiIndexMetadataConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) PutAlgorithmConfig(value *GoogleVertexAiIndexMetadataConfigAlgorithmConfig) {
	if err := g.validatePutAlgorithmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAlgorithmConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ResetAlgorithmConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAlgorithmConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ResetApproximateNeighborsCount() {
	_jsii_.InvokeVoid(
		g,
		"resetApproximateNeighborsCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ResetDistanceMeasureType() {
	_jsii_.InvokeVoid(
		g,
		"resetDistanceMeasureType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ResetFeatureNormType() {
	_jsii_.InvokeVoid(
		g,
		"resetFeatureNormType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiIndexMetadataConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

