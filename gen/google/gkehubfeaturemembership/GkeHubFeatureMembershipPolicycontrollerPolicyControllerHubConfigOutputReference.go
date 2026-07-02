package gkehubfeaturemembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkehubfeaturemembership/internal"
)

type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference interface {
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
	DeploymentConfigs() GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList
	DeploymentConfigsInput() any
	ExemptableNamespaces() *[]*string
	SetExemptableNamespaces(val *[]*string)
	ExemptableNamespacesInput() *[]*string
	// Experimental.
	Fqn() *string
	InstallSpec() *string
	SetInstallSpec(val *string)
	InstallSpecInput() *string
	InternalValue() *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig
	SetInternalValue(val *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig)
	LogDeniesEnabled() any
	SetLogDeniesEnabled(val any)
	LogDeniesEnabledInput() any
	Monitoring() GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference
	MonitoringInput() *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	MutationEnabled() any
	SetMutationEnabled(val any)
	MutationEnabledInput() any
	PolicyContent() GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference
	PolicyContentInput() *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
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
	PutMonitoring(value *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring)
	PutPolicyContent(value *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent)
	ResetAuditIntervalSeconds()
	ResetConstraintViolationLimit()
	ResetDeploymentConfigs()
	ResetExemptableNamespaces()
	ResetInstallSpec()
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

// The jsii proxy struct for GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference
type jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) AuditIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"auditIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) AuditIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"auditIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ConstraintViolationLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constraintViolationLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ConstraintViolationLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constraintViolationLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) DeploymentConfigs() GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList {
	var returns GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList
	_jsii_.Get(
		j,
		"deploymentConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) DeploymentConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deploymentConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ExemptableNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ExemptableNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InstallSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InstallSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InternalValue() *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	var returns *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) LogDeniesEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"logDeniesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) LogDeniesEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"logDeniesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) Monitoring() GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference {
	var returns GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) MonitoringInput() *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	var returns *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) MutationEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"mutationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) MutationEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"mutationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PolicyContent() GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference {
	var returns GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference
	_jsii_.Get(
		j,
		"policyContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PolicyContentInput() *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	var returns *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	_jsii_.Get(
		j,
		"policyContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ReferentialRulesEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"referentialRulesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ReferentialRulesEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"referentialRulesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeatureMembership.GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference_Override(g GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeatureMembership.GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetAuditIntervalSeconds(val *float64) {
	if err := j.validateSetAuditIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetConstraintViolationLimit(val *float64) {
	if err := j.validateSetConstraintViolationLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constraintViolationLimit",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetExemptableNamespaces(val *[]*string) {
	if err := j.validateSetExemptableNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exemptableNamespaces",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetInstallSpec(val *string) {
	if err := j.validateSetInstallSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installSpec",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetInternalValue(val *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetLogDeniesEnabled(val any) {
	if err := j.validateSetLogDeniesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDeniesEnabled",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetMutationEnabled(val any) {
	if err := j.validateSetMutationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutationEnabled",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetReferentialRulesEnabled(val any) {
	if err := j.validateSetReferentialRulesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referentialRulesEnabled",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PutDeploymentConfigs(value any) {
	if err := g.validatePutDeploymentConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentConfigs",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PutMonitoring(value *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) {
	if err := g.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoring",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PutPolicyContent(value *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) {
	if err := g.validatePutPolicyContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPolicyContent",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetAuditIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetConstraintViolationLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetConstraintViolationLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetDeploymentConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetExemptableNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetExemptableNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetInstallSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetInstallSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetLogDeniesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLogDeniesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetMutationEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMutationEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetPolicyContent() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyContent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetReferentialRulesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetReferentialRulesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
