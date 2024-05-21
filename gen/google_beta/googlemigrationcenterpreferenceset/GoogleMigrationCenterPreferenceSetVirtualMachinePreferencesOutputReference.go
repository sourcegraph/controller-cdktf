package googlemigrationcenterpreferenceset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemigrationcenterpreferenceset/internal"
)

type GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference interface {
	cdktf.ComplexObject
	CommitmentPlan() *string
	SetCommitmentPlan(val *string)
	CommitmentPlanInput() *string
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
	ComputeEnginePreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference
	ComputeEnginePreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferences
	SetInternalValue(val *GoogleMigrationCenterPreferenceSetVirtualMachinePreferences)
	RegionPreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferencesOutputReference
	RegionPreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences
	SizingOptimizationStrategy() *string
	SetSizingOptimizationStrategy(val *string)
	SizingOptimizationStrategyInput() *string
	SoleTenancyPreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesOutputReference
	SoleTenancyPreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences
	TargetProduct() *string
	SetTargetProduct(val *string)
	TargetProductInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmwareEnginePreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference
	VmwareEnginePreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences
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
	PutComputeEnginePreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences)
	PutRegionPreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences)
	PutSoleTenancyPreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences)
	PutVmwareEnginePreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences)
	ResetCommitmentPlan()
	ResetComputeEnginePreferences()
	ResetRegionPreferences()
	ResetSizingOptimizationStrategy()
	ResetSoleTenancyPreferences()
	ResetTargetProduct()
	ResetVmwareEnginePreferences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference
type jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) CommitmentPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) CommitmentPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComputeEnginePreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference {
	var returns GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference
	_jsii_.Get(
		j,
		"computeEnginePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComputeEnginePreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences {
	var returns *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences
	_jsii_.Get(
		j,
		"computeEnginePreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) InternalValue() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferences {
	var returns *GoogleMigrationCenterPreferenceSetVirtualMachinePreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) RegionPreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferencesOutputReference {
	var returns GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferencesOutputReference
	_jsii_.Get(
		j,
		"regionPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) RegionPreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences {
	var returns *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences
	_jsii_.Get(
		j,
		"regionPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SizingOptimizationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingOptimizationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SizingOptimizationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingOptimizationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SoleTenancyPreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesOutputReference {
	var returns GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesOutputReference
	_jsii_.Get(
		j,
		"soleTenancyPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SoleTenancyPreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences {
	var returns *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences
	_jsii_.Get(
		j,
		"soleTenancyPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TargetProduct() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProduct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TargetProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProductInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) VmwareEnginePreferences() GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference {
	var returns GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference
	_jsii_.Get(
		j,
		"vmwareEnginePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) VmwareEnginePreferencesInput() *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences {
	var returns *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences
	_jsii_.Get(
		j,
		"vmwareEnginePreferencesInput",
		&returns,
	)
	return returns
}


func NewGoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMigrationCenterPreferenceSet.GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference_Override(g GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMigrationCenterPreferenceSet.GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetCommitmentPlan(val *string) {
	if err := j.validateSetCommitmentPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitmentPlan",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetInternalValue(val *GoogleMigrationCenterPreferenceSetVirtualMachinePreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetSizingOptimizationStrategy(val *string) {
	if err := j.validateSetSizingOptimizationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingOptimizationStrategy",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetTargetProduct(val *string) {
	if err := j.validateSetTargetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProduct",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutComputeEnginePreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences) {
	if err := g.validatePutComputeEnginePreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putComputeEnginePreferences",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutRegionPreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences) {
	if err := g.validatePutRegionPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegionPreferences",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutSoleTenancyPreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences) {
	if err := g.validatePutSoleTenancyPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoleTenancyPreferences",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutVmwareEnginePreferences(value *GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences) {
	if err := g.validatePutVmwareEnginePreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVmwareEnginePreferences",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetCommitmentPlan() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitmentPlan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetComputeEnginePreferences() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeEnginePreferences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetRegionPreferences() {
	_jsii_.InvokeVoid(
		g,
		"resetRegionPreferences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetSizingOptimizationStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetSizingOptimizationStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetSoleTenancyPreferences() {
	_jsii_.InvokeVoid(
		g,
		"resetSoleTenancyPreferences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetTargetProduct() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetProduct",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetVmwareEnginePreferences() {
	_jsii_.InvokeVoid(
		g,
		"resetVmwareEnginePreferences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

