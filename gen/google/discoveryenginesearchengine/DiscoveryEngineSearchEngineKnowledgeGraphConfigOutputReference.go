package discoveryenginesearchengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginesearchengine/internal"
)

type DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference interface {
	cdktf.ComplexObject
	CloudKnowledgeGraphTypes() *[]*string
	SetCloudKnowledgeGraphTypes(val *[]*string)
	CloudKnowledgeGraphTypesInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	EnableCloudKnowledgeGraph() any
	SetEnableCloudKnowledgeGraph(val any)
	EnableCloudKnowledgeGraphInput() any
	EnablePrivateKnowledgeGraph() any
	SetEnablePrivateKnowledgeGraph(val any)
	EnablePrivateKnowledgeGraphInput() any
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
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutFeatureConfig(value *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig)
	ResetCloudKnowledgeGraphTypes()
	ResetEnableCloudKnowledgeGraph()
	ResetEnablePrivateKnowledgeGraph()
	ResetFeatureConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference
type jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) ComplexObjectIndex() any {
	var returns any
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

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnableCloudKnowledgeGraph() any {
	var returns any
	_jsii_.Get(
		j,
		"enableCloudKnowledgeGraph",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnableCloudKnowledgeGraphInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableCloudKnowledgeGraphInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnablePrivateKnowledgeGraph() any {
	var returns any
	_jsii_.Get(
		j,
		"enablePrivateKnowledgeGraph",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) EnablePrivateKnowledgeGraphInput() any {
	var returns any
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

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineSearchEngine.DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference_Override(d DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineSearchEngine.DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetCloudKnowledgeGraphTypes(val *[]*string) {
	if err := j.validateSetCloudKnowledgeGraphTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudKnowledgeGraphTypes",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetEnableCloudKnowledgeGraph(val any) {
	if err := j.validateSetEnableCloudKnowledgeGraphParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudKnowledgeGraph",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetEnablePrivateKnowledgeGraph(val any) {
	if err := j.validateSetEnablePrivateKnowledgeGraphParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateKnowledgeGraph",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetInternalValue(val *DiscoveryEngineSearchEngineKnowledgeGraphConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
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
		[]any{value},
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

func (d *jsiiProxy_DiscoveryEngineSearchEngineKnowledgeGraphConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
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
