package googlecomposerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomposerenvironment/internal"
)

type GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig
	SetInternalValue(val *GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig)
	SnapshotCreationSchedule() *string
	SetSnapshotCreationSchedule(val *string)
	SnapshotCreationScheduleInput() *string
	SnapshotLocation() *string
	SetSnapshotLocation(val *string)
	SnapshotLocationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	ResetSnapshotCreationSchedule()
	ResetSnapshotLocation()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference
type jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) InternalValue() *GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig {
	var returns *GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotCreationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCreationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotCreationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCreationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewGoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference_Override(g GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetInternalValue(val *GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetSnapshotCreationSchedule(val *string) {
	if err := j.validateSetSnapshotCreationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotCreationSchedule",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetSnapshotLocation(val *string) {
	if err := j.validateSetSnapshotLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ResetSnapshotCreationSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetSnapshotCreationSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ResetSnapshotLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetSnapshotLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

