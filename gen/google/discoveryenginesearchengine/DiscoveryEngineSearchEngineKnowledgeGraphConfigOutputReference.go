package discoveryenginesearchengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginesearchengine/internal"
)

type DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference interface {
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
	FeatureConfig() DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference
	FeatureConfigInput() *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig
	// Experimental.
	Fqn() *string
	InternalValue() *DiscoveryEngineSearchEngineKnowledgeGraphConfig
	SetInternalValue(val *DiscoveryEngineSearchEngineKnowledgeGraphConfig)
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
	PutFeatureConfig(value *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig)
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

// The jsii proxy struct for DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference
type jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) CloudKnowledgeGraphTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudKnowledgeGraphTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) CloudKnowledgeGraphTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudKnowledgeGraphTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnableCloudKnowledgeGraph() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudKnowledgeGraph",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnableCloudKnowledgeGraphInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudKnowledgeGraphInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnablePrivateKnowledgeGraph() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateKnowledgeGraph",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnablePrivateKnowledgeGraphInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateKnowledgeGraphInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) FeatureConfig() DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference {
	var returns DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfigOutputReference
	_jsii_.Get(
		j,
		"featureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) FeatureConfigInput() *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig {
	var returns *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig
	_jsii_.Get(
		j,
		"featureConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InternalValue() *DiscoveryEngineSearchEngineKnowledgeGraphConfig {
	var returns *DiscoveryEngineSearchEngineKnowledgeGraphConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.discoveryEngineSearchEngine.DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference_Override(d DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.discoveryEngineSearchEngine.DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetCloudKnowledgeGraphTypes(val *[]*string) {
	if err := j.validateSetCloudKnowledgeGraphTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudKnowledgeGraphTypes",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetEnableCloudKnowledgeGraph(val interface{}) {
	if err := j.validateSetEnableCloudKnowledgeGraphParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudKnowledgeGraph",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetEnablePrivateKnowledgeGraph(val interface{}) {
	if err := j.validateSetEnablePrivateKnowledgeGraphParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateKnowledgeGraph",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetInternalValue(val *DiscoveryEngineSearchEngineKnowledgeGraphConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) PutFeatureConfig(value *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig) {
	if err := d.validatePutFeatureConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFeatureConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetCloudKnowledgeGraphTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudKnowledgeGraphTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetEnableCloudKnowledgeGraph() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableCloudKnowledgeGraph",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetEnablePrivateKnowledgeGraph() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePrivateKnowledgeGraph",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ResetFeatureConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetFeatureConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

