package computeinstancefromtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeinstancefromtemplate/internal"
)

type ComputeInstanceFromTemplateSchedulingOutputReference interface {
	cdktf.ComplexObject
	AutomaticRestart() any
	SetAutomaticRestart(val any)
	AutomaticRestartInput() any
	AvailabilityDomain() *float64
	SetAvailabilityDomain(val *float64)
	AvailabilityDomainInput() *float64
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InstanceTerminationAction() *string
	SetInstanceTerminationAction(val *string)
	InstanceTerminationActionInput() *string
	InternalValue() *ComputeInstanceFromTemplateScheduling
	SetInternalValue(val *ComputeInstanceFromTemplateScheduling)
	LocalSsdRecoveryTimeout() ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeoutOutputReference
	LocalSsdRecoveryTimeoutInput() *ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeout
	MaxRunDuration() ComputeInstanceFromTemplateSchedulingMaxRunDurationOutputReference
	MaxRunDurationInput() *ComputeInstanceFromTemplateSchedulingMaxRunDuration
	MinNodeCpus() *float64
	SetMinNodeCpus(val *float64)
	MinNodeCpusInput() *float64
	NodeAffinities() ComputeInstanceFromTemplateSchedulingNodeAffinitiesList
	NodeAffinitiesInput() any
	OnHostMaintenance() *string
	SetOnHostMaintenance(val *string)
	OnHostMaintenanceInput() *string
	OnInstanceStopAction() ComputeInstanceFromTemplateSchedulingOnInstanceStopActionOutputReference
	OnInstanceStopActionInput() *ComputeInstanceFromTemplateSchedulingOnInstanceStopAction
	Preemptible() any
	SetPreemptible(val any)
	PreemptibleInput() any
	ProvisioningModel() *string
	SetProvisioningModel(val *string)
	ProvisioningModelInput() *string
	TerminationTime() *string
	SetTerminationTime(val *string)
	TerminationTimeInput() *string
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
	PutLocalSsdRecoveryTimeout(value *ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeout)
	PutMaxRunDuration(value *ComputeInstanceFromTemplateSchedulingMaxRunDuration)
	PutNodeAffinities(value any)
	PutOnInstanceStopAction(value *ComputeInstanceFromTemplateSchedulingOnInstanceStopAction)
	ResetAutomaticRestart()
	ResetAvailabilityDomain()
	ResetInstanceTerminationAction()
	ResetLocalSsdRecoveryTimeout()
	ResetMaxRunDuration()
	ResetMinNodeCpus()
	ResetNodeAffinities()
	ResetOnHostMaintenance()
	ResetOnInstanceStopAction()
	ResetPreemptible()
	ResetProvisioningModel()
	ResetTerminationTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceFromTemplateSchedulingOutputReference
type jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) AutomaticRestart() any {
	var returns any
	_jsii_.Get(
		j,
		"automaticRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) AutomaticRestartInput() any {
	var returns any
	_jsii_.Get(
		j,
		"automaticRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) AvailabilityDomain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) AvailabilityDomainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) InstanceTerminationAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) InstanceTerminationActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) InternalValue() *ComputeInstanceFromTemplateScheduling {
	var returns *ComputeInstanceFromTemplateScheduling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) LocalSsdRecoveryTimeout() ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeoutOutputReference {
	var returns ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeoutOutputReference
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) LocalSsdRecoveryTimeoutInput() *ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeout {
	var returns *ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeout
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) MaxRunDuration() ComputeInstanceFromTemplateSchedulingMaxRunDurationOutputReference {
	var returns ComputeInstanceFromTemplateSchedulingMaxRunDurationOutputReference
	_jsii_.Get(
		j,
		"maxRunDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) MaxRunDurationInput() *ComputeInstanceFromTemplateSchedulingMaxRunDuration {
	var returns *ComputeInstanceFromTemplateSchedulingMaxRunDuration
	_jsii_.Get(
		j,
		"maxRunDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) MinNodeCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) MinNodeCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) NodeAffinities() ComputeInstanceFromTemplateSchedulingNodeAffinitiesList {
	var returns ComputeInstanceFromTemplateSchedulingNodeAffinitiesList
	_jsii_.Get(
		j,
		"nodeAffinities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) NodeAffinitiesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nodeAffinitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) OnHostMaintenance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) OnHostMaintenanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) OnInstanceStopAction() ComputeInstanceFromTemplateSchedulingOnInstanceStopActionOutputReference {
	var returns ComputeInstanceFromTemplateSchedulingOnInstanceStopActionOutputReference
	_jsii_.Get(
		j,
		"onInstanceStopAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) OnInstanceStopActionInput() *ComputeInstanceFromTemplateSchedulingOnInstanceStopAction {
	var returns *ComputeInstanceFromTemplateSchedulingOnInstanceStopAction
	_jsii_.Get(
		j,
		"onInstanceStopActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) Preemptible() any {
	var returns any
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) PreemptibleInput() any {
	var returns any
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ProvisioningModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ProvisioningModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) TerminationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) TerminationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeInstanceFromTemplateSchedulingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceFromTemplateSchedulingOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceFromTemplateSchedulingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceFromTemplate.ComputeInstanceFromTemplateSchedulingOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceFromTemplateSchedulingOutputReference_Override(c ComputeInstanceFromTemplateSchedulingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceFromTemplate.ComputeInstanceFromTemplateSchedulingOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetAutomaticRestart(val any) {
	if err := j.validateSetAutomaticRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticRestart",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetAvailabilityDomain(val *float64) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetInstanceTerminationAction(val *string) {
	if err := j.validateSetInstanceTerminationActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTerminationAction",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetInternalValue(val *ComputeInstanceFromTemplateScheduling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetMinNodeCpus(val *float64) {
	if err := j.validateSetMinNodeCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCpus",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetOnHostMaintenance(val *string) {
	if err := j.validateSetOnHostMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onHostMaintenance",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetPreemptible(val any) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetProvisioningModel(val *string) {
	if err := j.validateSetProvisioningModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningModel",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetTerminationTime(val *string) {
	if err := j.validateSetTerminationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationTime",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) PutLocalSsdRecoveryTimeout(value *ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeout) {
	if err := c.validatePutLocalSsdRecoveryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocalSsdRecoveryTimeout",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) PutMaxRunDuration(value *ComputeInstanceFromTemplateSchedulingMaxRunDuration) {
	if err := c.validatePutMaxRunDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxRunDuration",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) PutNodeAffinities(value any) {
	if err := c.validatePutNodeAffinitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeAffinities",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) PutOnInstanceStopAction(value *ComputeInstanceFromTemplateSchedulingOnInstanceStopAction) {
	if err := c.validatePutOnInstanceStopActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnInstanceStopAction",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetAutomaticRestart() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomaticRestart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetInstanceTerminationAction() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceTerminationAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetLocalSsdRecoveryTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalSsdRecoveryTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetMaxRunDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRunDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetMinNodeCpus() {
	_jsii_.InvokeVoid(
		c,
		"resetMinNodeCpus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetNodeAffinities() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeAffinities",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetOnHostMaintenance() {
	_jsii_.InvokeVoid(
		c,
		"resetOnHostMaintenance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetOnInstanceStopAction() {
	_jsii_.InvokeVoid(
		c,
		"resetOnInstanceStopAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		c,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetProvisioningModel() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisioningModel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ResetTerminationTime() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateSchedulingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
