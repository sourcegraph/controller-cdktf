package googlegkehubfeaturemembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkehubfeaturemembership/internal"
)

type GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference interface {
	cdktf.ComplexObject
	AuditIntervalSeconds() *string
	SetAuditIntervalSeconds(val *string)
	AuditIntervalSecondsInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExemptableNamespaces() *[]*string
	SetExemptableNamespaces(val *[]*string)
	ExemptableNamespacesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGkeHubFeatureMembershipConfigmanagementPolicyController
	SetInternalValue(val *GoogleGkeHubFeatureMembershipConfigmanagementPolicyController)
	LogDeniesEnabled() interface{}
	SetLogDeniesEnabled(val interface{})
	LogDeniesEnabledInput() interface{}
	Monitoring() GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference
	MonitoringInput() *GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring
	MutationEnabled() interface{}
	SetMutationEnabled(val interface{})
	MutationEnabledInput() interface{}
	ReferentialRulesEnabled() interface{}
	SetReferentialRulesEnabled(val interface{})
	ReferentialRulesEnabledInput() interface{}
	TemplateLibraryInstalled() interface{}
	SetTemplateLibraryInstalled(val interface{})
	TemplateLibraryInstalledInput() interface{}
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
	PutMonitoring(value *GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring)
	ResetAuditIntervalSeconds()
	ResetEnabled()
	ResetExemptableNamespaces()
	ResetLogDeniesEnabled()
	ResetMonitoring()
	ResetMutationEnabled()
	ResetReferentialRulesEnabled()
	ResetTemplateLibraryInstalled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference
type jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) AuditIntervalSeconds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) AuditIntervalSecondsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ExemptableNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ExemptableNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exemptableNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) InternalValue() *GoogleGkeHubFeatureMembershipConfigmanagementPolicyController {
	var returns *GoogleGkeHubFeatureMembershipConfigmanagementPolicyController
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) LogDeniesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeniesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) LogDeniesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeniesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) Monitoring() GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference {
	var returns GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) MonitoringInput() *GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring {
	var returns *GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) MutationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) MutationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ReferentialRulesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referentialRulesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ReferentialRulesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referentialRulesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) TemplateLibraryInstalled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateLibraryInstalled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) TemplateLibraryInstalledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateLibraryInstalledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference_Override(g GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetAuditIntervalSeconds(val *string) {
	if err := j.validateSetAuditIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetExemptableNamespaces(val *[]*string) {
	if err := j.validateSetExemptableNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exemptableNamespaces",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetInternalValue(val *GoogleGkeHubFeatureMembershipConfigmanagementPolicyController) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetLogDeniesEnabled(val interface{}) {
	if err := j.validateSetLogDeniesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDeniesEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetMutationEnabled(val interface{}) {
	if err := j.validateSetMutationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutationEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetReferentialRulesEnabled(val interface{}) {
	if err := j.validateSetReferentialRulesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referentialRulesEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetTemplateLibraryInstalled(val interface{}) {
	if err := j.validateSetTemplateLibraryInstalledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateLibraryInstalled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) PutMonitoring(value *GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring) {
	if err := g.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetAuditIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetAuditIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetExemptableNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetExemptableNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetLogDeniesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLogDeniesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetMutationEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMutationEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetReferentialRulesEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetReferentialRulesEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ResetTemplateLibraryInstalled() {
	_jsii_.InvokeVoid(
		g,
		"resetTemplateLibraryInstalled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

