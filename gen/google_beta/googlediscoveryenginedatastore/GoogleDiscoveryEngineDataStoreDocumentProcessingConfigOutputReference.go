package googlediscoveryenginedatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginedatastore/internal"
)

type GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference interface {
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
	DefaultParsingConfig() GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference
	DefaultParsingConfigInput() *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDiscoveryEngineDataStoreDocumentProcessingConfig
	SetInternalValue(val *GoogleDiscoveryEngineDataStoreDocumentProcessingConfig)
	Name() *string
	ParsingConfigOverrides() GoogleDiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesList
	ParsingConfigOverridesInput() interface{}
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
	PutDefaultParsingConfig(value *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig)
	PutParsingConfigOverrides(value interface{})
	ResetDefaultParsingConfig()
	ResetParsingConfigOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference
type jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) DefaultParsingConfig() GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference {
	var returns GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference
	_jsii_.Get(
		j,
		"defaultParsingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) DefaultParsingConfigInput() *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig {
	var returns *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig
	_jsii_.Get(
		j,
		"defaultParsingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) InternalValue() *GoogleDiscoveryEngineDataStoreDocumentProcessingConfig {
	var returns *GoogleDiscoveryEngineDataStoreDocumentProcessingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ParsingConfigOverrides() GoogleDiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesList {
	var returns GoogleDiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesList
	_jsii_.Get(
		j,
		"parsingConfigOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ParsingConfigOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parsingConfigOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineDataStore.GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference_Override(g GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineDataStore.GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetInternalValue(val *GoogleDiscoveryEngineDataStoreDocumentProcessingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) PutDefaultParsingConfig(value *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig) {
	if err := g.validatePutDefaultParsingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultParsingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) PutParsingConfigOverrides(value interface{}) {
	if err := g.validatePutParsingConfigOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParsingConfigOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ResetDefaultParsingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultParsingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ResetParsingConfigOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetParsingConfigOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

