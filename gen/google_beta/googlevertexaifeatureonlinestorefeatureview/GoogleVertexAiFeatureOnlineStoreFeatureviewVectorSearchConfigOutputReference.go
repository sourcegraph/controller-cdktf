package googlevertexaifeatureonlinestorefeatureview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaifeatureonlinestorefeatureview/internal"
)

type GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference interface {
	cdktf.ComplexObject
	BruteForceConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfigOutputReference
	BruteForceConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfig
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
	CrowdingColumn() *string
	SetCrowdingColumn(val *string)
	CrowdingColumnInput() *string
	DistanceMeasureType() *string
	SetDistanceMeasureType(val *string)
	DistanceMeasureTypeInput() *string
	EmbeddingColumn() *string
	SetEmbeddingColumn(val *string)
	EmbeddingColumnInput() *string
	EmbeddingDimension() *float64
	SetEmbeddingDimension(val *float64)
	EmbeddingDimensionInput() *float64
	FilterColumns() *[]*string
	SetFilterColumns(val *[]*string)
	FilterColumnsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig
	SetInternalValue(val *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TreeAhConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfigOutputReference
	TreeAhConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfig
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
	PutBruteForceConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfig)
	PutTreeAhConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfig)
	ResetBruteForceConfig()
	ResetCrowdingColumn()
	ResetDistanceMeasureType()
	ResetEmbeddingDimension()
	ResetFilterColumns()
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

// The jsii proxy struct for GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference
type jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) BruteForceConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfigOutputReference {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfigOutputReference
	_jsii_.Get(
		j,
		"bruteForceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) BruteForceConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfig {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfig
	_jsii_.Get(
		j,
		"bruteForceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) CrowdingColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crowdingColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) CrowdingColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crowdingColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) DistanceMeasureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceMeasureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) DistanceMeasureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceMeasureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) EmbeddingColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) EmbeddingColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) EmbeddingDimension() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"embeddingDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) EmbeddingDimensionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"embeddingDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) FilterColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) FilterColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) InternalValue() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) TreeAhConfig() GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfigOutputReference {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfigOutputReference
	_jsii_.Get(
		j,
		"treeAhConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) TreeAhConfigInput() *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfig {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfig
	_jsii_.Get(
		j,
		"treeAhConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference_Override(g GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetCrowdingColumn(val *string) {
	if err := j.validateSetCrowdingColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crowdingColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetDistanceMeasureType(val *string) {
	if err := j.validateSetDistanceMeasureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distanceMeasureType",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetEmbeddingColumn(val *string) {
	if err := j.validateSetEmbeddingColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingColumn",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetEmbeddingDimension(val *float64) {
	if err := j.validateSetEmbeddingDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingDimension",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetFilterColumns(val *[]*string) {
	if err := j.validateSetFilterColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterColumns",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetInternalValue(val *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) PutBruteForceConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfig) {
	if err := g.validatePutBruteForceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBruteForceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) PutTreeAhConfig(value *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfig) {
	if err := g.validatePutTreeAhConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTreeAhConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ResetBruteForceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBruteForceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ResetCrowdingColumn() {
	_jsii_.InvokeVoid(
		g,
		"resetCrowdingColumn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ResetDistanceMeasureType() {
	_jsii_.InvokeVoid(
		g,
		"resetDistanceMeasureType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ResetEmbeddingDimension() {
	_jsii_.InvokeVoid(
		g,
		"resetEmbeddingDimension",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ResetFilterColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ResetTreeAhConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTreeAhConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

