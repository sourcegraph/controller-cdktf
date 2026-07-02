package gkehubfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkehubfeature/internal"
)

type GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference interface {
	cdktf.ComplexObject
	AuditIntervalSeconds() *float64
	SetAuditIntervalSeconds(val *float64)
	AuditIntervalSecondsInput() *float64
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
	ConstraintViolationLimit() *float64
	SetConstraintViolationLimit(val *float64)
	ConstraintViolationLimitInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentConfigs() GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList
	DeploymentConfigsInput() any
	ExemptableNamespaces() *[]*string
	SetExemptableNamespaces(val *[]*string)
	ExemptableNamespacesInput() *[]*string
	// Experimental.
	Fqn() *string
	InstallSpec() *string
	SetInstallSpec(val *string)
	InstallSpecInput() *string
	InternalValue() *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig
	SetInternalValue(val *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig)
	LogDeniesEnabled() any
	SetLogDeniesEnabled(val any)
	LogDeniesEnabledInput() any
	Monitoring() GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference
	MonitoringInput() *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoring
	MutationEnabled() any
	SetMutationEnabled(val any)
	MutationEnabledInput() any
	PolicyContent() GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference
	PolicyContentInput() *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent
	ReferentialRulesEnabled() any
	SetReferentialRulesEnabled(val any)
	ReferentialRulesEnabledInput() any
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
	PutDeploymentConfigs(value any)
	PutMonitoring(value *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoring)
	PutPolicyContent(value *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent)
	ResetAuditIntervalSeconds()
	ResetConstraintViolationLimit()
	ResetDeploymentConfigs()
	ResetExemptableNamespaces()
	ResetLogDeniesEnabled()
	ResetMonitoring()
	ResetMutationEnabled()
	ResetPolicyContent()
	ResetReferentialRulesEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference
type jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) AuditIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"auditIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) AuditIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"auditIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ConstraintViolationLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constraintViolationLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ConstraintViolationLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constraintViolationLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) DeploymentConfigs() GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList {
	var returns GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList
	_jsii_.Get(
		j,
		"deploymentConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) DeploymentConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deploymentConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ExemptableNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ExemptableNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) InstallSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) InstallSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) InternalValue() *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig {
	var returns *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) LogDeniesEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"logDeniesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) LogDeniesEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"logDeniesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) Monitoring() GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference {
	var returns GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) MonitoringInput() *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoring {
	var returns *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) MutationEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"mutationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) MutationEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"mutationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) PolicyContent() GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference {
	var returns GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference
	_jsii_.Get(
		j,
		"policyContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) PolicyContentInput() *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent {
	var returns *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent
	_jsii_.Get(
		j,
		"policyContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ReferentialRulesEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"referentialRulesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ReferentialRulesEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"referentialRulesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference_Override(g GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetAuditIntervalSeconds(val *float64) {
	if err := j.validateSetAuditIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetConstraintViolationLimit(val *float64) {
	if err := j.validateSetConstraintViolationLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constraintViolationLimit",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetExemptableNamespaces(val *[]*string) {
	if err := j.validateSetExemptableNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exemptableNamespaces",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetInstallSpec(val *string) {
	if err := j.validateSetInstallSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installSpec",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetInternalValue(val *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetLogDeniesEnabled(val any) {
	if err := j.validateSetLogDeniesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDeniesEnabled",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetMutationEnabled(val any) {
	if err := j.validateSetMutationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutationEnabled",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetReferentialRulesEnabled(val any) {
	if err := j.validateSetReferentialRulesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referentialRulesEnabled",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) PutDeploymentConfigs(value any) {
	if err := g.validatePutDeploymentConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentConfigs",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) PutMonitoring(value *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoring) {
	if err := g.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoring",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) PutPolicyContent(value *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent) {
	if err := g.validatePutPolicyContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPolicyContent",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetAuditIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetConstraintViolationLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetConstraintViolationLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetDeploymentConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetExemptableNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetExemptableNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetLogDeniesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLogDeniesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetMutationEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMutationEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetPolicyContent() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyContent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ResetReferentialRulesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetReferentialRulesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
