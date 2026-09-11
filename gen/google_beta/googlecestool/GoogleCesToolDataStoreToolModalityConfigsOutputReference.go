package googlecestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecestool/internal"
)

type GoogleCesToolDataStoreToolModalityConfigsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GroundingConfig() GoogleCesToolDataStoreToolModalityConfigsGroundingConfigOutputReference
	GroundingConfigInput() *GoogleCesToolDataStoreToolModalityConfigsGroundingConfig
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModalityType() *string
	SetModalityType(val *string)
	ModalityTypeInput() *string
	RewriterConfig() GoogleCesToolDataStoreToolModalityConfigsRewriterConfigOutputReference
	RewriterConfigInput() *GoogleCesToolDataStoreToolModalityConfigsRewriterConfig
	SummarizationConfig() GoogleCesToolDataStoreToolModalityConfigsSummarizationConfigOutputReference
	SummarizationConfigInput() *GoogleCesToolDataStoreToolModalityConfigsSummarizationConfig
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
	PutGroundingConfig(value *GoogleCesToolDataStoreToolModalityConfigsGroundingConfig)
	PutRewriterConfig(value *GoogleCesToolDataStoreToolModalityConfigsRewriterConfig)
	PutSummarizationConfig(value *GoogleCesToolDataStoreToolModalityConfigsSummarizationConfig)
	ResetGroundingConfig()
	ResetRewriterConfig()
	ResetSummarizationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesToolDataStoreToolModalityConfigsOutputReference
type jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GroundingConfig() GoogleCesToolDataStoreToolModalityConfigsGroundingConfigOutputReference {
	var returns GoogleCesToolDataStoreToolModalityConfigsGroundingConfigOutputReference
	_jsii_.Get(
		j,
		"groundingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GroundingConfigInput() *GoogleCesToolDataStoreToolModalityConfigsGroundingConfig {
	var returns *GoogleCesToolDataStoreToolModalityConfigsGroundingConfig
	_jsii_.Get(
		j,
		"groundingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ModalityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modalityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ModalityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modalityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) RewriterConfig() GoogleCesToolDataStoreToolModalityConfigsRewriterConfigOutputReference {
	var returns GoogleCesToolDataStoreToolModalityConfigsRewriterConfigOutputReference
	_jsii_.Get(
		j,
		"rewriterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) RewriterConfigInput() *GoogleCesToolDataStoreToolModalityConfigsRewriterConfig {
	var returns *GoogleCesToolDataStoreToolModalityConfigsRewriterConfig
	_jsii_.Get(
		j,
		"rewriterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) SummarizationConfig() GoogleCesToolDataStoreToolModalityConfigsSummarizationConfigOutputReference {
	var returns GoogleCesToolDataStoreToolModalityConfigsSummarizationConfigOutputReference
	_jsii_.Get(
		j,
		"summarizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) SummarizationConfigInput() *GoogleCesToolDataStoreToolModalityConfigsSummarizationConfig {
	var returns *GoogleCesToolDataStoreToolModalityConfigsSummarizationConfig
	_jsii_.Get(
		j,
		"summarizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesToolDataStoreToolModalityConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCesToolDataStoreToolModalityConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolDataStoreToolModalityConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolDataStoreToolModalityConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCesToolDataStoreToolModalityConfigsOutputReference_Override(g GoogleCesToolDataStoreToolModalityConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolDataStoreToolModalityConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference)SetModalityType(val *string) {
	if err := j.validateSetModalityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modalityType",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) PutGroundingConfig(value *GoogleCesToolDataStoreToolModalityConfigsGroundingConfig) {
	if err := g.validatePutGroundingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGroundingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) PutRewriterConfig(value *GoogleCesToolDataStoreToolModalityConfigsRewriterConfig) {
	if err := g.validatePutRewriterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRewriterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) PutSummarizationConfig(value *GoogleCesToolDataStoreToolModalityConfigsSummarizationConfig) {
	if err := g.validatePutSummarizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSummarizationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ResetGroundingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGroundingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ResetRewriterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRewriterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ResetSummarizationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSummarizationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolModalityConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

