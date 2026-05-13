package googlenetappbackupvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetappbackupvault/internal"
)

type GoogleNetappBackupVaultBackupRetentionPolicyOutputReference interface {
	cdktf.ComplexObject
	BackupMinimumEnforcedRetentionDays() *float64
	SetBackupMinimumEnforcedRetentionDays(val *float64)
	BackupMinimumEnforcedRetentionDaysInput() *float64
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
	DailyBackupImmutable() interface{}
	SetDailyBackupImmutable(val interface{})
	DailyBackupImmutableInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetappBackupVaultBackupRetentionPolicy
	SetInternalValue(val *GoogleNetappBackupVaultBackupRetentionPolicy)
	ManualBackupImmutable() interface{}
	SetManualBackupImmutable(val interface{})
	ManualBackupImmutableInput() interface{}
	MonthlyBackupImmutable() interface{}
	SetMonthlyBackupImmutable(val interface{})
	MonthlyBackupImmutableInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklyBackupImmutable() interface{}
	SetWeeklyBackupImmutable(val interface{})
	WeeklyBackupImmutableInput() interface{}
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
	ResetDailyBackupImmutable()
	ResetManualBackupImmutable()
	ResetMonthlyBackupImmutable()
	ResetWeeklyBackupImmutable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetappBackupVaultBackupRetentionPolicyOutputReference
type jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) BackupMinimumEnforcedRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinimumEnforcedRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) BackupMinimumEnforcedRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinimumEnforcedRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) DailyBackupImmutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) DailyBackupImmutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyBackupImmutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) InternalValue() *GoogleNetappBackupVaultBackupRetentionPolicy {
	var returns *GoogleNetappBackupVaultBackupRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ManualBackupImmutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ManualBackupImmutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualBackupImmutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) MonthlyBackupImmutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) MonthlyBackupImmutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyBackupImmutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) WeeklyBackupImmutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) WeeklyBackupImmutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklyBackupImmutableInput",
		&returns,
	)
	return returns
}


func NewGoogleNetappBackupVaultBackupRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetappBackupVaultBackupRetentionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetappBackupVaultBackupRetentionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappBackupVault.GoogleNetappBackupVaultBackupRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetappBackupVaultBackupRetentionPolicyOutputReference_Override(g GoogleNetappBackupVaultBackupRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappBackupVault.GoogleNetappBackupVaultBackupRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetBackupMinimumEnforcedRetentionDays(val *float64) {
	if err := j.validateSetBackupMinimumEnforcedRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinimumEnforcedRetentionDays",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetDailyBackupImmutable(val interface{}) {
	if err := j.validateSetDailyBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyBackupImmutable",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetInternalValue(val *GoogleNetappBackupVaultBackupRetentionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetManualBackupImmutable(val interface{}) {
	if err := j.validateSetManualBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualBackupImmutable",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetMonthlyBackupImmutable(val interface{}) {
	if err := j.validateSetMonthlyBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthlyBackupImmutable",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference)SetWeeklyBackupImmutable(val interface{}) {
	if err := j.validateSetWeeklyBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyBackupImmutable",
		val,
	)
}

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ResetDailyBackupImmutable() {
	_jsii_.InvokeVoid(
		g,
		"resetDailyBackupImmutable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ResetManualBackupImmutable() {
	_jsii_.InvokeVoid(
		g,
		"resetManualBackupImmutable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ResetMonthlyBackupImmutable() {
	_jsii_.InvokeVoid(
		g,
		"resetMonthlyBackupImmutable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ResetWeeklyBackupImmutable() {
	_jsii_.InvokeVoid(
		g,
		"resetWeeklyBackupImmutable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetappBackupVaultBackupRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

