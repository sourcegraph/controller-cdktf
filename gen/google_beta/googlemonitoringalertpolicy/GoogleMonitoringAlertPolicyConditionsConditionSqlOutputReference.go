package googlemonitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringalertpolicy/internal"
)

type GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference interface {
	cdktf.ComplexObject
	BooleanTest() GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTestOutputReference
	BooleanTestInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTest
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
	Daily() GoogleMonitoringAlertPolicyConditionsConditionSqlDailyOutputReference
	DailyInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlDaily
	// Experimental.
	Fqn() *string
	Hourly() GoogleMonitoringAlertPolicyConditionsConditionSqlHourlyOutputReference
	HourlyInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlHourly
	InternalValue() *GoogleMonitoringAlertPolicyConditionsConditionSql
	SetInternalValue(val *GoogleMonitoringAlertPolicyConditionsConditionSql)
	Minutes() GoogleMonitoringAlertPolicyConditionsConditionSqlMinutesOutputReference
	MinutesInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlMinutes
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	RowCountTest() GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTestOutputReference
	RowCountTestInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest
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
	PutBooleanTest(value *GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTest)
	PutDaily(value *GoogleMonitoringAlertPolicyConditionsConditionSqlDaily)
	PutHourly(value *GoogleMonitoringAlertPolicyConditionsConditionSqlHourly)
	PutMinutes(value *GoogleMonitoringAlertPolicyConditionsConditionSqlMinutes)
	PutRowCountTest(value *GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest)
	ResetBooleanTest()
	ResetDaily()
	ResetHourly()
	ResetMinutes()
	ResetRowCountTest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference
type jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) BooleanTest() GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTestOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTestOutputReference
	_jsii_.Get(
		j,
		"booleanTest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) BooleanTestInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTest {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTest
	_jsii_.Get(
		j,
		"booleanTestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) Daily() GoogleMonitoringAlertPolicyConditionsConditionSqlDailyOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionSqlDailyOutputReference
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) DailyInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlDaily {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionSqlDaily
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) Hourly() GoogleMonitoringAlertPolicyConditionsConditionSqlHourlyOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionSqlHourlyOutputReference
	_jsii_.Get(
		j,
		"hourly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) HourlyInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlHourly {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionSqlHourly
	_jsii_.Get(
		j,
		"hourlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) InternalValue() *GoogleMonitoringAlertPolicyConditionsConditionSql {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionSql
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) Minutes() GoogleMonitoringAlertPolicyConditionsConditionSqlMinutesOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionSqlMinutesOutputReference
	_jsii_.Get(
		j,
		"minutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) MinutesInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlMinutes {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionSqlMinutes
	_jsii_.Get(
		j,
		"minutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) RowCountTest() GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTestOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTestOutputReference
	_jsii_.Get(
		j,
		"rowCountTest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) RowCountTestInput() *GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest
	_jsii_.Get(
		j,
		"rowCountTestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringAlertPolicyConditionsConditionSqlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference_Override(g GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference)SetInternalValue(val *GoogleMonitoringAlertPolicyConditionsConditionSql) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) PutBooleanTest(value *GoogleMonitoringAlertPolicyConditionsConditionSqlBooleanTest) {
	if err := g.validatePutBooleanTestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBooleanTest",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) PutDaily(value *GoogleMonitoringAlertPolicyConditionsConditionSqlDaily) {
	if err := g.validatePutDailyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDaily",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) PutHourly(value *GoogleMonitoringAlertPolicyConditionsConditionSqlHourly) {
	if err := g.validatePutHourlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHourly",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) PutMinutes(value *GoogleMonitoringAlertPolicyConditionsConditionSqlMinutes) {
	if err := g.validatePutMinutesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMinutes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) PutRowCountTest(value *GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest) {
	if err := g.validatePutRowCountTestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRowCountTest",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetBooleanTest() {
	_jsii_.InvokeVoid(
		g,
		"resetBooleanTest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		g,
		"resetDaily",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetHourly() {
	_jsii_.InvokeVoid(
		g,
		"resetHourly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetMinutes() {
	_jsii_.InvokeVoid(
		g,
		"resetMinutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetRowCountTest() {
	_jsii_.InvokeVoid(
		g,
		"resetRowCountTest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionSqlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

