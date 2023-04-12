package teamsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/teamsaccount/internal"
)

type TeamsAccountLoggingSettingsByRuleTypeOutputReference interface {
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
	Dns() TeamsAccountLoggingSettingsByRuleTypeDnsOutputReference
	DnsInput() *TeamsAccountLoggingSettingsByRuleTypeDns
	// Experimental.
	Fqn() *string
	Http() TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference
	HttpInput() *TeamsAccountLoggingSettingsByRuleTypeHttp
	InternalValue() *TeamsAccountLoggingSettingsByRuleType
	SetInternalValue(val *TeamsAccountLoggingSettingsByRuleType)
	L4() TeamsAccountLoggingSettingsByRuleTypeL4OutputReference
	L4Input() *TeamsAccountLoggingSettingsByRuleTypeL4
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
	PutDns(value *TeamsAccountLoggingSettingsByRuleTypeDns)
	PutHttp(value *TeamsAccountLoggingSettingsByRuleTypeHttp)
	PutL4(value *TeamsAccountLoggingSettingsByRuleTypeL4)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamsAccountLoggingSettingsByRuleTypeOutputReference
type jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) Dns() TeamsAccountLoggingSettingsByRuleTypeDnsOutputReference {
	var returns TeamsAccountLoggingSettingsByRuleTypeDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) DnsInput() *TeamsAccountLoggingSettingsByRuleTypeDns {
	var returns *TeamsAccountLoggingSettingsByRuleTypeDns
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) Http() TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference {
	var returns TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) HttpInput() *TeamsAccountLoggingSettingsByRuleTypeHttp {
	var returns *TeamsAccountLoggingSettingsByRuleTypeHttp
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) InternalValue() *TeamsAccountLoggingSettingsByRuleType {
	var returns *TeamsAccountLoggingSettingsByRuleType
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) L4() TeamsAccountLoggingSettingsByRuleTypeL4OutputReference {
	var returns TeamsAccountLoggingSettingsByRuleTypeL4OutputReference
	_jsii_.Get(
		j,
		"l4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) L4Input() *TeamsAccountLoggingSettingsByRuleTypeL4 {
	var returns *TeamsAccountLoggingSettingsByRuleTypeL4
	_jsii_.Get(
		j,
		"l4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTeamsAccountLoggingSettingsByRuleTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamsAccountLoggingSettingsByRuleTypeOutputReference {
	_init_.Initialize()

	if err := validateNewTeamsAccountLoggingSettingsByRuleTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccountLoggingSettingsByRuleTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamsAccountLoggingSettingsByRuleTypeOutputReference_Override(t TeamsAccountLoggingSettingsByRuleTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccountLoggingSettingsByRuleTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference)SetInternalValue(val *TeamsAccountLoggingSettingsByRuleType) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) PutDns(value *TeamsAccountLoggingSettingsByRuleTypeDns) {
	if err := t.validatePutDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) PutHttp(value *TeamsAccountLoggingSettingsByRuleTypeHttp) {
	if err := t.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) PutL4(value *TeamsAccountLoggingSettingsByRuleTypeL4) {
	if err := t.validatePutL4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putL4",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

