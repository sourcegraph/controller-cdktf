package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference interface {
	cdktn.ComplexObject
	AutoscaledRolloutPolicy() GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicyOutputReference
	AutoscaledRolloutPolicyInput() *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy
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
	InternalValue() *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettings
	SetInternalValue(val *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettings)
	NodePoolSoakDuration() *string
	SetNodePoolSoakDuration(val *string)
	NodePoolSoakDurationInput() *string
	StandardRolloutPolicy() GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference
	StandardRolloutPolicyInput() *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy
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
	PutAutoscaledRolloutPolicy(value *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy)
	PutStandardRolloutPolicy(value *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy)
	ResetAutoscaledRolloutPolicy()
	ResetNodePoolSoakDuration()
	ResetStandardRolloutPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference
type jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) AutoscaledRolloutPolicy() GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicyOutputReference {
	var returns GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicyOutputReference
	_jsii_.Get(
		j,
		"autoscaledRolloutPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) AutoscaledRolloutPolicyInput() *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy {
	var returns *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy
	_jsii_.Get(
		j,
		"autoscaledRolloutPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) InternalValue() *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettings {
	var returns *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) NodePoolSoakDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePoolSoakDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) NodePoolSoakDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePoolSoakDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) StandardRolloutPolicy() GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference {
	var returns GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference
	_jsii_.Get(
		j,
		"standardRolloutPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) StandardRolloutPolicyInput() *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy {
	var returns *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy
	_jsii_.Get(
		j,
		"standardRolloutPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference_Override(g GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference)SetInternalValue(val *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference)SetNodePoolSoakDuration(val *string) {
	if err := j.validateSetNodePoolSoakDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePoolSoakDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) PutAutoscaledRolloutPolicy(value *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy) {
	if err := g.validatePutAutoscaledRolloutPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscaledRolloutPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) PutStandardRolloutPolicy(value *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy) {
	if err := g.validatePutStandardRolloutPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStandardRolloutPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) ResetAutoscaledRolloutPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscaledRolloutPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) ResetNodePoolSoakDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetNodePoolSoakDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) ResetStandardRolloutPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetStandardRolloutPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

