package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference interface {
	cdktf.ComplexObject
	BlueGreenSettings() GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsOutputReference
	BlueGreenSettingsInput() *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings
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
	InternalValue() *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings
	SetInternalValue(val *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings)
	MaxSurge() *float64
	SetMaxSurge(val *float64)
	MaxSurgeInput() *float64
	MaxUnavailable() *float64
	SetMaxUnavailable(val *float64)
	MaxUnavailableInput() *float64
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
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
	PutBlueGreenSettings(value *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings)
	ResetBlueGreenSettings()
	ResetMaxSurge()
	ResetMaxUnavailable()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference
type jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) BlueGreenSettings() GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsOutputReference {
	var returns GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsOutputReference
	_jsii_.Get(
		j,
		"blueGreenSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) BlueGreenSettingsInput() *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings {
	var returns *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings
	_jsii_.Get(
		j,
		"blueGreenSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) InternalValue() *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings {
	var returns *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxSurge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxSurgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxUnavailable() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxUnavailableInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference_Override(g GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetInternalValue(val *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetMaxSurge(val *float64) {
	if err := j.validateSetMaxSurgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurge",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetMaxUnavailable(val *float64) {
	if err := j.validateSetMaxUnavailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) PutBlueGreenSettings(value *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings) {
	if err := g.validatePutBlueGreenSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBlueGreenSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetBlueGreenSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetBlueGreenSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetMaxSurge() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxSurge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetMaxUnavailable() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxUnavailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

