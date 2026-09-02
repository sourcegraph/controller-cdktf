package googlecomputeinstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancetemplate/internal"
)

type GoogleComputeInstanceTemplateSchedulingOutputReference interface {
	cdktn.ComplexObject
	AutomaticRestart() interface{}
	SetAutomaticRestart(val interface{})
	AutomaticRestartInput() interface{}
	AvailabilityDomain() *float64
	SetAvailabilityDomain(val *float64)
	AvailabilityDomainInput() *float64
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
	GracefulShutdown() GoogleComputeInstanceTemplateSchedulingGracefulShutdownOutputReference
	GracefulShutdownInput() *GoogleComputeInstanceTemplateSchedulingGracefulShutdown
	HostErrorTimeoutSeconds() *float64
	SetHostErrorTimeoutSeconds(val *float64)
	HostErrorTimeoutSecondsInput() *float64
	InstanceTerminationAction() *string
	SetInstanceTerminationAction(val *string)
	InstanceTerminationActionInput() *string
	InternalValue() *GoogleComputeInstanceTemplateScheduling
	SetInternalValue(val *GoogleComputeInstanceTemplateScheduling)
	LocalSsdRecoveryTimeout() GoogleComputeInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList
	LocalSsdRecoveryTimeoutInput() interface{}
	MaintenanceInterval() *string
	SetMaintenanceInterval(val *string)
	MaintenanceIntervalInput() *string
	MaxRunDuration() GoogleComputeInstanceTemplateSchedulingMaxRunDurationOutputReference
	MaxRunDurationInput() *GoogleComputeInstanceTemplateSchedulingMaxRunDuration
	MinNodeCpus() *float64
	SetMinNodeCpus(val *float64)
	MinNodeCpusInput() *float64
	NodeAffinities() GoogleComputeInstanceTemplateSchedulingNodeAffinitiesList
	NodeAffinitiesInput() interface{}
	OnHostMaintenance() *string
	SetOnHostMaintenance(val *string)
	OnHostMaintenanceInput() *string
	OnInstanceStopAction() GoogleComputeInstanceTemplateSchedulingOnInstanceStopActionOutputReference
	OnInstanceStopActionInput() *GoogleComputeInstanceTemplateSchedulingOnInstanceStopAction
	Preemptible() interface{}
	SetPreemptible(val interface{})
	PreemptibleInput() interface{}
	PreemptionNoticeDuration() GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDurationOutputReference
	PreemptionNoticeDurationInput() *GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDuration
	ProvisioningModel() *string
	SetProvisioningModel(val *string)
	ProvisioningModelInput() *string
	SkipGuestOsShutdown() interface{}
	SetSkipGuestOsShutdown(val interface{})
	SkipGuestOsShutdownInput() interface{}
	TerminationTime() *string
	SetTerminationTime(val *string)
	TerminationTimeInput() *string
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
	PutGracefulShutdown(value *GoogleComputeInstanceTemplateSchedulingGracefulShutdown)
	PutLocalSsdRecoveryTimeout(value interface{})
	PutMaxRunDuration(value *GoogleComputeInstanceTemplateSchedulingMaxRunDuration)
	PutNodeAffinities(value interface{})
	PutOnInstanceStopAction(value *GoogleComputeInstanceTemplateSchedulingOnInstanceStopAction)
	PutPreemptionNoticeDuration(value *GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDuration)
	ResetAutomaticRestart()
	ResetAvailabilityDomain()
	ResetGracefulShutdown()
	ResetHostErrorTimeoutSeconds()
	ResetInstanceTerminationAction()
	ResetLocalSsdRecoveryTimeout()
	ResetMaintenanceInterval()
	ResetMaxRunDuration()
	ResetMinNodeCpus()
	ResetNodeAffinities()
	ResetOnHostMaintenance()
	ResetOnInstanceStopAction()
	ResetPreemptible()
	ResetPreemptionNoticeDuration()
	ResetProvisioningModel()
	ResetSkipGuestOsShutdown()
	ResetTerminationTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceTemplateSchedulingOutputReference
type jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) AutomaticRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) AutomaticRestartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) AvailabilityDomain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) AvailabilityDomainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GracefulShutdown() GoogleComputeInstanceTemplateSchedulingGracefulShutdownOutputReference {
	var returns GoogleComputeInstanceTemplateSchedulingGracefulShutdownOutputReference
	_jsii_.Get(
		j,
		"gracefulShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GracefulShutdownInput() *GoogleComputeInstanceTemplateSchedulingGracefulShutdown {
	var returns *GoogleComputeInstanceTemplateSchedulingGracefulShutdown
	_jsii_.Get(
		j,
		"gracefulShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) HostErrorTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hostErrorTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) HostErrorTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hostErrorTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) InstanceTerminationAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) InstanceTerminationActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) InternalValue() *GoogleComputeInstanceTemplateScheduling {
	var returns *GoogleComputeInstanceTemplateScheduling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) LocalSsdRecoveryTimeout() GoogleComputeInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList {
	var returns GoogleComputeInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) LocalSsdRecoveryTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) MaintenanceInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) MaintenanceIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) MaxRunDuration() GoogleComputeInstanceTemplateSchedulingMaxRunDurationOutputReference {
	var returns GoogleComputeInstanceTemplateSchedulingMaxRunDurationOutputReference
	_jsii_.Get(
		j,
		"maxRunDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) MaxRunDurationInput() *GoogleComputeInstanceTemplateSchedulingMaxRunDuration {
	var returns *GoogleComputeInstanceTemplateSchedulingMaxRunDuration
	_jsii_.Get(
		j,
		"maxRunDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) MinNodeCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) MinNodeCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) NodeAffinities() GoogleComputeInstanceTemplateSchedulingNodeAffinitiesList {
	var returns GoogleComputeInstanceTemplateSchedulingNodeAffinitiesList
	_jsii_.Get(
		j,
		"nodeAffinities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) NodeAffinitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeAffinitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) OnHostMaintenance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) OnHostMaintenanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) OnInstanceStopAction() GoogleComputeInstanceTemplateSchedulingOnInstanceStopActionOutputReference {
	var returns GoogleComputeInstanceTemplateSchedulingOnInstanceStopActionOutputReference
	_jsii_.Get(
		j,
		"onInstanceStopAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) OnInstanceStopActionInput() *GoogleComputeInstanceTemplateSchedulingOnInstanceStopAction {
	var returns *GoogleComputeInstanceTemplateSchedulingOnInstanceStopAction
	_jsii_.Get(
		j,
		"onInstanceStopActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PreemptionNoticeDuration() GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDurationOutputReference {
	var returns GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDurationOutputReference
	_jsii_.Get(
		j,
		"preemptionNoticeDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PreemptionNoticeDurationInput() *GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDuration {
	var returns *GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDuration
	_jsii_.Get(
		j,
		"preemptionNoticeDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ProvisioningModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ProvisioningModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) SkipGuestOsShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGuestOsShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) SkipGuestOsShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGuestOsShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) TerminationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) TerminationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceTemplateSchedulingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeInstanceTemplateSchedulingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceTemplateSchedulingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplateSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceTemplateSchedulingOutputReference_Override(g GoogleComputeInstanceTemplateSchedulingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplateSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetAutomaticRestart(val interface{}) {
	if err := j.validateSetAutomaticRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticRestart",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetAvailabilityDomain(val *float64) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetHostErrorTimeoutSeconds(val *float64) {
	if err := j.validateSetHostErrorTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostErrorTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetInstanceTerminationAction(val *string) {
	if err := j.validateSetInstanceTerminationActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTerminationAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetInternalValue(val *GoogleComputeInstanceTemplateScheduling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetMaintenanceInterval(val *string) {
	if err := j.validateSetMaintenanceIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetMinNodeCpus(val *float64) {
	if err := j.validateSetMinNodeCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCpus",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetOnHostMaintenance(val *string) {
	if err := j.validateSetOnHostMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onHostMaintenance",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetProvisioningModel(val *string) {
	if err := j.validateSetProvisioningModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningModel",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetSkipGuestOsShutdown(val interface{}) {
	if err := j.validateSetSkipGuestOsShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGuestOsShutdown",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetTerminationTime(val *string) {
	if err := j.validateSetTerminationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationTime",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PutGracefulShutdown(value *GoogleComputeInstanceTemplateSchedulingGracefulShutdown) {
	if err := g.validatePutGracefulShutdownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGracefulShutdown",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PutLocalSsdRecoveryTimeout(value interface{}) {
	if err := g.validatePutLocalSsdRecoveryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocalSsdRecoveryTimeout",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PutMaxRunDuration(value *GoogleComputeInstanceTemplateSchedulingMaxRunDuration) {
	if err := g.validatePutMaxRunDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaxRunDuration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PutNodeAffinities(value interface{}) {
	if err := g.validatePutNodeAffinitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeAffinities",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PutOnInstanceStopAction(value *GoogleComputeInstanceTemplateSchedulingOnInstanceStopAction) {
	if err := g.validatePutOnInstanceStopActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnInstanceStopAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) PutPreemptionNoticeDuration(value *GoogleComputeInstanceTemplateSchedulingPreemptionNoticeDuration) {
	if err := g.validatePutPreemptionNoticeDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPreemptionNoticeDuration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetAutomaticRestart() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomaticRestart",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetGracefulShutdown() {
	_jsii_.InvokeVoid(
		g,
		"resetGracefulShutdown",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetHostErrorTimeoutSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetHostErrorTimeoutSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetInstanceTerminationAction() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceTerminationAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetLocalSsdRecoveryTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsdRecoveryTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetMaintenanceInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetMaxRunDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRunDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetMinNodeCpus() {
	_jsii_.InvokeVoid(
		g,
		"resetMinNodeCpus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetNodeAffinities() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeAffinities",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetOnHostMaintenance() {
	_jsii_.InvokeVoid(
		g,
		"resetOnHostMaintenance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetOnInstanceStopAction() {
	_jsii_.InvokeVoid(
		g,
		"resetOnInstanceStopAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetPreemptionNoticeDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptionNoticeDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetProvisioningModel() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisioningModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetSkipGuestOsShutdown() {
	_jsii_.InvokeVoid(
		g,
		"resetSkipGuestOsShutdown",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetTerminationTime() {
	_jsii_.InvokeVoid(
		g,
		"resetTerminationTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

