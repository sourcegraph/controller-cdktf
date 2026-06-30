package discoveryenginerecommendationengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginerecommendationengine/internal"
)

type DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig
	SetInternalValue(val *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig)
	MostPopularConfig() DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfigOutputReference
	MostPopularConfigInput() *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig
	RecommendedForYouConfig() DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfigOutputReference
	RecommendedForYouConfigInput() *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfig
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
	PutMostPopularConfig(value *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig)
	PutRecommendedForYouConfig(value *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfig)
	ResetMostPopularConfig()
	ResetRecommendedForYouConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference
type jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) InternalValue() *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig {
	var returns *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) MostPopularConfig() DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfigOutputReference {
	var returns DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfigOutputReference
	_jsii_.Get(
		j,
		"mostPopularConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) MostPopularConfigInput() *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig {
	var returns *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig
	_jsii_.Get(
		j,
		"mostPopularConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) RecommendedForYouConfig() DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfigOutputReference {
	var returns DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfigOutputReference
	_jsii_.Get(
		j,
		"recommendedForYouConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) RecommendedForYouConfigInput() *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfig {
	var returns *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfig
	_jsii_.Get(
		j,
		"recommendedForYouConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineRecommendationEngine.DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference_Override(d DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineRecommendationEngine.DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) SetInternalValue(val *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) PutMostPopularConfig(value *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig) {
	if err := d.validatePutMostPopularConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMostPopularConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) PutRecommendedForYouConfig(value *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfig) {
	if err := d.validatePutRecommendedForYouConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRecommendedForYouConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) ResetMostPopularConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetMostPopularConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) ResetRecommendedForYouConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetRecommendedForYouConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
