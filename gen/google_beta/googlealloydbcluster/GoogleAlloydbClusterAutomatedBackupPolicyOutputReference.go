package googlealloydbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlealloydbcluster/internal"
)

type GoogleAlloydbClusterAutomatedBackupPolicyOutputReference interface {
	cdktf.ComplexObject
	BackupWindow() *string
	SetBackupWindow(val *string)
	BackupWindowInput() *string
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
	EncryptionConfig() GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAlloydbClusterAutomatedBackupPolicy
	SetInternalValue(val *GoogleAlloydbClusterAutomatedBackupPolicy)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	QuantityBasedRetention() GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionOutputReference
	QuantityBasedRetentionInput() *GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeBasedRetention() GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionOutputReference
	TimeBasedRetentionInput() *GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention
	WeeklySchedule() GoogleAlloydbClusterAutomatedBackupPolicyWeeklyScheduleOutputReference
	WeeklyScheduleInput() *GoogleAlloydbClusterAutomatedBackupPolicyWeeklySchedule
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
	PutEncryptionConfig(value *GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfig)
	PutQuantityBasedRetention(value *GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention)
	PutTimeBasedRetention(value *GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention)
	PutWeeklySchedule(value *GoogleAlloydbClusterAutomatedBackupPolicyWeeklySchedule)
	ResetBackupWindow()
	ResetEnabled()
	ResetEncryptionConfig()
	ResetLabels()
	ResetLocation()
	ResetQuantityBasedRetention()
	ResetTimeBasedRetention()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAlloydbClusterAutomatedBackupPolicyOutputReference
type jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) BackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) BackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) EncryptionConfig() GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfigOutputReference {
	var returns GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) EncryptionConfigInput() *GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfig {
	var returns *GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) InternalValue() *GoogleAlloydbClusterAutomatedBackupPolicy {
	var returns *GoogleAlloydbClusterAutomatedBackupPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) QuantityBasedRetention() GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionOutputReference {
	var returns GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionOutputReference
	_jsii_.Get(
		j,
		"quantityBasedRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) QuantityBasedRetentionInput() *GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention {
	var returns *GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention
	_jsii_.Get(
		j,
		"quantityBasedRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) TimeBasedRetention() GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionOutputReference {
	var returns GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionOutputReference
	_jsii_.Get(
		j,
		"timeBasedRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) TimeBasedRetentionInput() *GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention {
	var returns *GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention
	_jsii_.Get(
		j,
		"timeBasedRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) WeeklySchedule() GoogleAlloydbClusterAutomatedBackupPolicyWeeklyScheduleOutputReference {
	var returns GoogleAlloydbClusterAutomatedBackupPolicyWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) WeeklyScheduleInput() *GoogleAlloydbClusterAutomatedBackupPolicyWeeklySchedule {
	var returns *GoogleAlloydbClusterAutomatedBackupPolicyWeeklySchedule
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewGoogleAlloydbClusterAutomatedBackupPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAlloydbClusterAutomatedBackupPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAlloydbClusterAutomatedBackupPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleAlloydbCluster.GoogleAlloydbClusterAutomatedBackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAlloydbClusterAutomatedBackupPolicyOutputReference_Override(g GoogleAlloydbClusterAutomatedBackupPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleAlloydbCluster.GoogleAlloydbClusterAutomatedBackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetBackupWindow(val *string) {
	if err := j.validateSetBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupWindow",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetInternalValue(val *GoogleAlloydbClusterAutomatedBackupPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) PutEncryptionConfig(value *GoogleAlloydbClusterAutomatedBackupPolicyEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) PutQuantityBasedRetention(value *GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention) {
	if err := g.validatePutQuantityBasedRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQuantityBasedRetention",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) PutTimeBasedRetention(value *GoogleAlloydbClusterAutomatedBackupPolicyTimeBasedRetention) {
	if err := g.validatePutTimeBasedRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeBasedRetention",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) PutWeeklySchedule(value *GoogleAlloydbClusterAutomatedBackupPolicyWeeklySchedule) {
	if err := g.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetBackupWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetQuantityBasedRetention() {
	_jsii_.InvokeVoid(
		g,
		"resetQuantityBasedRetention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetTimeBasedRetention() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeBasedRetention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAlloydbClusterAutomatedBackupPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

