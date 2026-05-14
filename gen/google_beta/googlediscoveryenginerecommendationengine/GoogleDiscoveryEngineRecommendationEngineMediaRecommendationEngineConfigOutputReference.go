package googlediscoveryenginerecommendationengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginerecommendationengine/internal"
)

type GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference interface {
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
	EngineFeaturesConfig() GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference
	EngineFeaturesConfigInput() *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfig
	SetInternalValue(val *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfig)
	OptimizationObjective() *string
	SetOptimizationObjective(val *string)
	OptimizationObjectiveConfig() GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfigOutputReference
	OptimizationObjectiveConfigInput() *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfig
	OptimizationObjectiveInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrainingState() *string
	SetTrainingState(val *string)
	TrainingStateInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutEngineFeaturesConfig(value *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig)
	PutOptimizationObjectiveConfig(value *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfig)
	ResetEngineFeaturesConfig()
	ResetOptimizationObjective()
	ResetOptimizationObjectiveConfig()
	ResetTrainingState()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference
type jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) EngineFeaturesConfig() GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference {
	var returns GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigOutputReference
	_jsii_.Get(
		j,
		"engineFeaturesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) EngineFeaturesConfigInput() *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig {
	var returns *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig
	_jsii_.Get(
		j,
		"engineFeaturesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) InternalValue() *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfig {
	var returns *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) OptimizationObjective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optimizationObjective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) OptimizationObjectiveConfig() GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfigOutputReference {
	var returns GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfigOutputReference
	_jsii_.Get(
		j,
		"optimizationObjectiveConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) OptimizationObjectiveConfigInput() *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfig {
	var returns *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfig
	_jsii_.Get(
		j,
		"optimizationObjectiveConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) OptimizationObjectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optimizationObjectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) TrainingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) TrainingStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineRecommendationEngine.GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference_Override(g GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineRecommendationEngine.GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetInternalValue(val *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetOptimizationObjective(val *string) {
	if err := j.validateSetOptimizationObjectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizationObjective",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetTrainingState(val *string) {
	if err := j.validateSetTrainingStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingState",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) PutEngineFeaturesConfig(value *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig) {
	if err := g.validatePutEngineFeaturesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEngineFeaturesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) PutOptimizationObjectiveConfig(value *GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfig) {
	if err := g.validatePutOptimizationObjectiveConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOptimizationObjectiveConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ResetEngineFeaturesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEngineFeaturesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ResetOptimizationObjective() {
	_jsii_.InvokeVoid(
		g,
		"resetOptimizationObjective",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ResetOptimizationObjectiveConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOptimizationObjectiveConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ResetTrainingState() {
	_jsii_.InvokeVoid(
		g,
		"resetTrainingState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

