package notificationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/notificationpolicy/internal"
)

type NotificationPolicyFiltersOutputReference interface {
	cdktf.ComplexObject
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
	Enabled() *[]*string
	SetEnabled(val *[]*string)
	EnabledInput() *[]*string
	EventSource() *[]*string
	SetEventSource(val *[]*string)
	EventSourceInput() *[]*string
	EventType() *[]*string
	SetEventType(val *[]*string)
	EventTypeInput() *[]*string
	// Experimental.
	Fqn() *string
	HealthCheckId() *[]*string
	SetHealthCheckId(val *[]*string)
	HealthCheckIdInput() *[]*string
	InputId() *[]*string
	SetInputId(val *[]*string)
	InputIdInput() *[]*string
	InternalValue() *NotificationPolicyFilters
	SetInternalValue(val *NotificationPolicyFilters)
	Limit() *[]*string
	SetLimit(val *[]*string)
	LimitInput() *[]*string
	NewHealth() *[]*string
	SetNewHealth(val *[]*string)
	NewHealthInput() *[]*string
	PacketsPerSecond() *[]*string
	SetPacketsPerSecond(val *[]*string)
	PacketsPerSecondInput() *[]*string
	PoolId() *[]*string
	SetPoolId(val *[]*string)
	PoolIdInput() *[]*string
	Product() *[]*string
	SetProduct(val *[]*string)
	ProductInput() *[]*string
	Protocol() *[]*string
	SetProtocol(val *[]*string)
	ProtocolInput() *[]*string
	RequestsPerSecond() *[]*string
	SetRequestsPerSecond(val *[]*string)
	RequestsPerSecondInput() *[]*string
	Services() *[]*string
	SetServices(val *[]*string)
	ServicesInput() *[]*string
	Slo() *[]*string
	SetSlo(val *[]*string)
	SloInput() *[]*string
	Status() *[]*string
	SetStatus(val *[]*string)
	StatusInput() *[]*string
	TargetHost() *[]*string
	SetTargetHost(val *[]*string)
	TargetHostInput() *[]*string
	TargetZoneName() *[]*string
	SetTargetZoneName(val *[]*string)
	TargetZoneNameInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	ResetEnabled()
	ResetEventSource()
	ResetEventType()
	ResetHealthCheckId()
	ResetInputId()
	ResetLimit()
	ResetNewHealth()
	ResetPacketsPerSecond()
	ResetPoolId()
	ResetProduct()
	ResetProtocol()
	ResetRequestsPerSecond()
	ResetServices()
	ResetSlo()
	ResetStatus()
	ResetTargetHost()
	ResetTargetZoneName()
	ResetZones()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationPolicyFiltersOutputReference
type jsiiProxy_NotificationPolicyFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Enabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EnabledInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventSource() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventSourceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) EventTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) HealthCheckId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) HealthCheckIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InputId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InputIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) InternalValue() *NotificationPolicyFilters {
	var returns *NotificationPolicyFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Limit() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) LimitInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) NewHealth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) NewHealthInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"newHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PacketsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packetsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PacketsPerSecondInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packetsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PoolId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) PoolIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"poolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Product() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ProductInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Protocol() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ProtocolInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) RequestsPerSecond() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) RequestsPerSecondInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Slo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"slo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) SloInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sloInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Status() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) StatusInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetZoneName() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TargetZoneNameInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetZoneNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


func NewNotificationPolicyFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotificationPolicyFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationPolicyFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationPolicyFiltersOutputReference{}

	_jsii_.Create(
		"cloudflare.notificationPolicy.NotificationPolicyFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationPolicyFiltersOutputReference_Override(n NotificationPolicyFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.notificationPolicy.NotificationPolicyFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEnabled(val *[]*string) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEventSource(val *[]*string) {
	if err := j.validateSetEventSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSource",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetEventType(val *[]*string) {
	if err := j.validateSetEventTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventType",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetHealthCheckId(val *[]*string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetInputId(val *[]*string) {
	if err := j.validateSetInputIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetInternalValue(val *NotificationPolicyFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetLimit(val *[]*string) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetNewHealth(val *[]*string) {
	if err := j.validateSetNewHealthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newHealth",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetPacketsPerSecond(val *[]*string) {
	if err := j.validateSetPacketsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packetsPerSecond",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetPoolId(val *[]*string) {
	if err := j.validateSetPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolId",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetProduct(val *[]*string) {
	if err := j.validateSetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetProtocol(val *[]*string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetRequestsPerSecond(val *[]*string) {
	if err := j.validateSetRequestsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestsPerSecond",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetServices(val *[]*string) {
	if err := j.validateSetServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"services",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetSlo(val *[]*string) {
	if err := j.validateSetSloParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slo",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetStatus(val *[]*string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTargetHost(val *[]*string) {
	if err := j.validateSetTargetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetHost",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTargetZoneName(val *[]*string) {
	if err := j.validateSetTargetZoneNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetZoneName",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicyFiltersOutputReference)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEventSource() {
	_jsii_.InvokeVoid(
		n,
		"resetEventSource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetEventType() {
	_jsii_.InvokeVoid(
		n,
		"resetEventType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		n,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetInputId() {
	_jsii_.InvokeVoid(
		n,
		"resetInputId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		n,
		"resetLimit",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetNewHealth() {
	_jsii_.InvokeVoid(
		n,
		"resetNewHealth",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetPacketsPerSecond() {
	_jsii_.InvokeVoid(
		n,
		"resetPacketsPerSecond",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetPoolId() {
	_jsii_.InvokeVoid(
		n,
		"resetPoolId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetProduct() {
	_jsii_.InvokeVoid(
		n,
		"resetProduct",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetRequestsPerSecond() {
	_jsii_.InvokeVoid(
		n,
		"resetRequestsPerSecond",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		n,
		"resetServices",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetSlo() {
	_jsii_.InvokeVoid(
		n,
		"resetSlo",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		n,
		"resetStatus",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTargetHost() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetHost",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetTargetZoneName() {
	_jsii_.InvokeVoid(
		n,
		"resetTargetZoneName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ResetZones() {
	_jsii_.InvokeVoid(
		n,
		"resetZones",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicyFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

