package googlediscoveryenginesearchengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginesearchengine/internal"
)

type GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference interface {
	cdktn.ComplexObject
	CloudKnowledgeGraphTypes() *[]*string
	SetCloudKnowledgeGraphTypes(val *[]*string)
	CloudKnowledgeGraphTypesInput() *[]*string
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
	EnableCloudKnowledgeGraph() interface{}
	SetEnableCloudKnowledgeGraph(val interface{})
	EnableCloudKnowledgeGraphInput() interface{}
	EnablePrivateKnowledgeGraph() interface{}
	SetEnablePrivateKnowledgeGraph(val interface{})
	EnablePrivateKnowledgeGraphInput() interface{}
	FeatureConfig() GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference
	FeatureConfigInput() *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfig
	SetInternalValue(val *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfig)
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
	PutFeatureConfig(value *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig)
	ResetCloudKnowledgeGraphTypes()
	ResetEnableCloudKnowledgeGraph()
	ResetEnablePrivateKnowledgeGraph()
	ResetFeatureConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference
type jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) CloudKnowledgeGraphTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudKnowledgeGraphTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) CloudKnowledgeGraphTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudKnowledgeGraphTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnableCloudKnowledgeGraph() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudKnowledgeGraph",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnableCloudKnowledgeGraphInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudKnowledgeGraphInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnablePrivateKnowledgeGraph() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateKnowledgeGraph",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnablePrivateKnowledgeGraphInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateKnowledgeGraphInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) FeatureConfig() GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference {
	var returns GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference
	_jsii_.Get(
		j,
		"featureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) FeatureConfigInput() *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig {
	var returns *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig
	_jsii_.Get(
		j,
		"featureConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InternalValue() *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfig {
	var returns *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineSearchEngine.GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference_Override(g GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineSearchEngine.GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetCloudKnowledgeGraphTypes(val *[]*string) {
	if err := j.validateSetCloudKnowledgeGraphTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudKnowledgeGraphTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetEnableCloudKnowledgeGraph(val interface{}) {
	if err := j.validateSetEnableCloudKnowledgeGraphParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudKnowledgeGraph",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetEnablePrivateKnowledgeGraph(val interface{}) {
	if err := j.validateSetEnablePrivateKnowledgeGraphParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateKnowledgeGraph",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetInternalValue(val *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) PutFeatureConfig(value *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig) {
	if err := g.validatePutFeatureConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFeatureConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetCloudKnowledgeGraphTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudKnowledgeGraphTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetEnableCloudKnowledgeGraph() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCloudKnowledgeGraph",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetEnablePrivateKnowledgeGraph() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePrivateKnowledgeGraph",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetFeatureConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetFeatureConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

