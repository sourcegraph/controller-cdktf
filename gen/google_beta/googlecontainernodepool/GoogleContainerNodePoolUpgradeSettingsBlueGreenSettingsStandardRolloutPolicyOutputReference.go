package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainernodepool/internal"
)

type GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference interface {
	cdktf.ComplexObject
	BatchNodeCount() *float64
	SetBatchNodeCount(val *float64)
	BatchNodeCountInput() *float64
	BatchPercentage() *float64
	SetBatchPercentage(val *float64)
	BatchPercentageInput() *float64
	BatchSoakDuration() *string
	SetBatchSoakDuration(val *string)
	BatchSoakDurationInput() *string
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
	InternalValue() *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy
	SetInternalValue(val *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy)
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
	ResetBatchNodeCount()
	ResetBatchPercentage()
	ResetBatchSoakDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference
type jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchSoakDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSoakDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchSoakDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSoakDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) InternalValue() *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy {
	var returns *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerNodePool.GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference_Override(g GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerNodePool.GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetBatchNodeCount(val *float64) {
	if err := j.validateSetBatchNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetBatchPercentage(val *float64) {
	if err := j.validateSetBatchPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchPercentage",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetBatchSoakDuration(val *string) {
	if err := j.validateSetBatchSoakDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSoakDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetInternalValue(val *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ResetBatchNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetBatchNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ResetBatchPercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetBatchPercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ResetBatchSoakDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetBatchSoakDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

