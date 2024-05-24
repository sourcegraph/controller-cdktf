package googlegkehubfeaturemembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkehubfeaturemembership/internal"
)

type GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference interface {
	cdktf.ComplexObject
	AuditIntervalSeconds() *float64
	SetAuditIntervalSeconds(val *float64)
	AuditIntervalSecondsInput() *float64
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
	ConstraintViolationLimit() *float64
	SetConstraintViolationLimit(val *float64)
	ConstraintViolationLimitInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentConfigs() GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList
	DeploymentConfigsInput() interface{}
	ExemptableNamespaces() *[]*string
	SetExemptableNamespaces(val *[]*string)
	ExemptableNamespacesInput() *[]*string
	// Experimental.
	Fqn() *string
	InstallSpec() *string
	SetInstallSpec(val *string)
	InstallSpecInput() *string
	InternalValue() *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig
	SetInternalValue(val *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig)
	LogDeniesEnabled() interface{}
	SetLogDeniesEnabled(val interface{})
	LogDeniesEnabledInput() interface{}
	Monitoring() GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference
	MonitoringInput() *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	MutationEnabled() interface{}
	SetMutationEnabled(val interface{})
	MutationEnabledInput() interface{}
	PolicyContent() GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference
	PolicyContentInput() *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	ReferentialRulesEnabled() interface{}
	SetReferentialRulesEnabled(val interface{})
	ReferentialRulesEnabledInput() interface{}
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
	PutDeploymentConfigs(value interface{})
	PutMonitoring(value *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring)
	PutPolicyContent(value *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent)
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference
type jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) AuditIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"auditIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) AuditIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"auditIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ConstraintViolationLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constraintViolationLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ConstraintViolationLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"constraintViolationLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) DeploymentConfigs() GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList {
	var returns GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsList
	_jsii_.Get(
		j,
		"deploymentConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) DeploymentConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ExemptableNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ExemptableNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InstallSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InstallSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InternalValue() *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	var returns *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) LogDeniesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeniesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) LogDeniesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeniesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) Monitoring() GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference {
	var returns GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) MonitoringInput() *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	var returns *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) MutationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) MutationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PolicyContent() GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference {
	var returns GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentOutputReference
	_jsii_.Get(
		j,
		"policyContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PolicyContentInput() *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	var returns *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	_jsii_.Get(
		j,
		"policyContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ReferentialRulesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referentialRulesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ReferentialRulesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referentialRulesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference_Override(g GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetAuditIntervalSeconds(val *float64) {
	if err := j.validateSetAuditIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetConstraintViolationLimit(val *float64) {
	if err := j.validateSetConstraintViolationLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constraintViolationLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetExemptableNamespaces(val *[]*string) {
	if err := j.validateSetExemptableNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exemptableNamespaces",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetInstallSpec(val *string) {
	if err := j.validateSetInstallSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installSpec",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetInternalValue(val *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetLogDeniesEnabled(val interface{}) {
	if err := j.validateSetLogDeniesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDeniesEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetMutationEnabled(val interface{}) {
	if err := j.validateSetMutationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutationEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetReferentialRulesEnabled(val interface{}) {
	if err := j.validateSetReferentialRulesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referentialRulesEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PutDeploymentConfigs(value interface{}) {
	if err := g.validatePutDeploymentConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PutMonitoring(value *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) {
	if err := g.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) PutPolicyContent(value *GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) {
	if err := g.validatePutPolicyContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPolicyContent",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetAuditIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetConstraintViolationLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetConstraintViolationLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetDeploymentConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetExemptableNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetExemptableNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetInstallSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetInstallSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetLogDeniesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLogDeniesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetMutationEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMutationEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetPolicyContent() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyContent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ResetReferentialRulesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetReferentialRulesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

