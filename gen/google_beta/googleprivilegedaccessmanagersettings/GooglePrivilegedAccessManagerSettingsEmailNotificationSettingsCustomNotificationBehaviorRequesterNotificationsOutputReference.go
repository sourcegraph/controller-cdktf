package googleprivilegedaccessmanagersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivilegedaccessmanagersettings/internal"
)

type GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference interface {
	cdktn.ComplexObject
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
	EntitlementAssigned() *string
	SetEntitlementAssigned(val *string)
	EntitlementAssignedInput() *string
	// Experimental.
	Fqn() *string
	GrantActivated() *string
	SetGrantActivated(val *string)
	GrantActivatedInput() *string
	GrantActivationFailed() *string
	SetGrantActivationFailed(val *string)
	GrantActivationFailedInput() *string
	GrantDenied() *string
	SetGrantDenied(val *string)
	GrantDeniedInput() *string
	GrantEnded() *string
	SetGrantEnded(val *string)
	GrantEndedInput() *string
	GrantExpired() *string
	SetGrantExpired(val *string)
	GrantExpiredInput() *string
	GrantExternallyModified() *string
	SetGrantExternallyModified(val *string)
	GrantExternallyModifiedInput() *string
	GrantRevoked() *string
	SetGrantRevoked(val *string)
	GrantRevokedInput() *string
	InternalValue() *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotifications
	SetInternalValue(val *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotifications)
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
	ResetEntitlementAssigned()
	ResetGrantActivated()
	ResetGrantActivationFailed()
	ResetGrantDenied()
	ResetGrantEnded()
	ResetGrantExpired()
	ResetGrantExternallyModified()
	ResetGrantRevoked()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference
type jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) EntitlementAssigned() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementAssigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) EntitlementAssignedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementAssignedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantActivated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantActivated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantActivatedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantActivatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantActivationFailed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantActivationFailed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantActivationFailedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantActivationFailedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantDenied() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantDenied",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantDeniedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantDeniedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantEnded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantEnded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantEndedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantEndedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantExpired() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantExpired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantExpiredInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantExpiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantExternallyModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantExternallyModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantExternallyModifiedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantExternallyModifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantRevoked() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantRevoked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GrantRevokedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantRevokedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) InternalValue() *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotifications {
	var returns *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerSettings.GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference_Override(g GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerSettings.GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetEntitlementAssigned(val *string) {
	if err := j.validateSetEntitlementAssignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entitlementAssigned",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetGrantActivated(val *string) {
	if err := j.validateSetGrantActivatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantActivated",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetGrantActivationFailed(val *string) {
	if err := j.validateSetGrantActivationFailedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantActivationFailed",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetGrantDenied(val *string) {
	if err := j.validateSetGrantDeniedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantDenied",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetGrantEnded(val *string) {
	if err := j.validateSetGrantEndedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantEnded",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetGrantExpired(val *string) {
	if err := j.validateSetGrantExpiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantExpired",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetGrantExternallyModified(val *string) {
	if err := j.validateSetGrantExternallyModifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantExternallyModified",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetGrantRevoked(val *string) {
	if err := j.validateSetGrantRevokedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantRevoked",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetInternalValue(val *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetEntitlementAssigned() {
	_jsii_.InvokeVoid(
		g,
		"resetEntitlementAssigned",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetGrantActivated() {
	_jsii_.InvokeVoid(
		g,
		"resetGrantActivated",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetGrantActivationFailed() {
	_jsii_.InvokeVoid(
		g,
		"resetGrantActivationFailed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetGrantDenied() {
	_jsii_.InvokeVoid(
		g,
		"resetGrantDenied",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetGrantEnded() {
	_jsii_.InvokeVoid(
		g,
		"resetGrantEnded",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetGrantExpired() {
	_jsii_.InvokeVoid(
		g,
		"resetGrantExpired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetGrantExternallyModified() {
	_jsii_.InvokeVoid(
		g,
		"resetGrantExternallyModified",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ResetGrantRevoked() {
	_jsii_.InvokeVoid(
		g,
		"resetGrantRevoked",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

