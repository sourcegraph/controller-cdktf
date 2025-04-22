package googlevmwareengineprivatecloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevmwareengineprivatecloud/internal"
)

type GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference interface {
	cdktf.ComplexObject
	AutoscalingPolicies() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesList
	AutoscalingPoliciesInput() interface{}
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
	CoolDownPeriod() *string
	SetCoolDownPeriod(val *string)
	CoolDownPeriodInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings
	SetInternalValue(val *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings)
	MaxClusterNodeCount() *float64
	SetMaxClusterNodeCount(val *float64)
	MaxClusterNodeCountInput() *float64
	MinClusterNodeCount() *float64
	SetMinClusterNodeCount(val *float64)
	MinClusterNodeCountInput() *float64
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
	PutAutoscalingPolicies(value interface{})
	ResetCoolDownPeriod()
	ResetMaxClusterNodeCount()
	ResetMinClusterNodeCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference
type jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) AutoscalingPolicies() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesList {
	var returns GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesList
	_jsii_.Get(
		j,
		"autoscalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) AutoscalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) CoolDownPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) CoolDownPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) InternalValue() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings {
	var returns *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MaxClusterNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MaxClusterNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MinClusterNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MinClusterNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareenginePrivateCloud.GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference_Override(g GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareenginePrivateCloud.GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetCoolDownPeriod(val *string) {
	if err := j.validateSetCoolDownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coolDownPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetInternalValue(val *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetMaxClusterNodeCount(val *float64) {
	if err := j.validateSetMaxClusterNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClusterNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetMinClusterNodeCount(val *float64) {
	if err := j.validateSetMinClusterNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minClusterNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) PutAutoscalingPolicies(value interface{}) {
	if err := g.validatePutAutoscalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ResetCoolDownPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCoolDownPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ResetMaxClusterNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxClusterNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ResetMinClusterNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMinClusterNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

