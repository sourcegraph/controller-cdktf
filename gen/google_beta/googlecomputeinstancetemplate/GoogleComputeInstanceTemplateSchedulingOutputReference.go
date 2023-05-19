package googlecomputeinstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancetemplate/internal"
)

type GoogleComputeInstanceTemplateSchedulingOutputReference interface {
	cdktf.ComplexObject
	AutomaticRestart() interface{}
	SetAutomaticRestart(val interface{})
	AutomaticRestartInput() interface{}
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
	InstanceTerminationAction() *string
	SetInstanceTerminationAction(val *string)
	InstanceTerminationActionInput() *string
	InternalValue() *GoogleComputeInstanceTemplateScheduling
	SetInternalValue(val *GoogleComputeInstanceTemplateScheduling)
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
	Preemptible() interface{}
	SetPreemptible(val interface{})
	PreemptibleInput() interface{}
	ProvisioningModel() *string
	SetProvisioningModel(val *string)
	ProvisioningModelInput() *string
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
	PutMaxRunDuration(value *GoogleComputeInstanceTemplateSchedulingMaxRunDuration)
	PutNodeAffinities(value interface{})
	ResetAutomaticRestart()
	ResetInstanceTerminationAction()
	ResetMaintenanceInterval()
	ResetMaxRunDuration()
	ResetMinNodeCpus()
	ResetNodeAffinities()
	ResetOnHostMaintenance()
	ResetPreemptible()
	ResetProvisioningModel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceTemplateSchedulingOutputReference
type jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceTemplateSchedulingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeInstanceTemplateSchedulingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceTemplateSchedulingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplateSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceTemplateSchedulingOutputReference_Override(g GoogleComputeInstanceTemplateSchedulingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplateSchedulingOutputReference",
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

func (j *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetAutomaticRestart() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomaticRestart",
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		g,
		"resetPreemptible",
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

func (g *jsiiProxy_GoogleComputeInstanceTemplateSchedulingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

