package googlednsrecordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlednsrecordset/internal"
)

type GoogleDnsRecordSetRoutingPolicyOutputReference interface {
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
	EnableGeoFencing() interface{}
	SetEnableGeoFencing(val interface{})
	EnableGeoFencingInput() interface{}
	// Experimental.
	Fqn() *string
	Geo() GoogleDnsRecordSetRoutingPolicyGeoList
	GeoInput() interface{}
	InternalValue() *GoogleDnsRecordSetRoutingPolicy
	SetInternalValue(val *GoogleDnsRecordSetRoutingPolicy)
	PrimaryBackup() GoogleDnsRecordSetRoutingPolicyPrimaryBackupOutputReference
	PrimaryBackupInput() *GoogleDnsRecordSetRoutingPolicyPrimaryBackup
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wrr() GoogleDnsRecordSetRoutingPolicyWrrList
	WrrInput() interface{}
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
	PutGeo(value interface{})
	PutPrimaryBackup(value *GoogleDnsRecordSetRoutingPolicyPrimaryBackup)
	PutWrr(value interface{})
	ResetEnableGeoFencing()
	ResetGeo()
	ResetPrimaryBackup()
	ResetWrr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDnsRecordSetRoutingPolicyOutputReference
type jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) EnableGeoFencing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGeoFencing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) EnableGeoFencingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGeoFencingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) Geo() GoogleDnsRecordSetRoutingPolicyGeoList {
	var returns GoogleDnsRecordSetRoutingPolicyGeoList
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GeoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) InternalValue() *GoogleDnsRecordSetRoutingPolicy {
	var returns *GoogleDnsRecordSetRoutingPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) PrimaryBackup() GoogleDnsRecordSetRoutingPolicyPrimaryBackupOutputReference {
	var returns GoogleDnsRecordSetRoutingPolicyPrimaryBackupOutputReference
	_jsii_.Get(
		j,
		"primaryBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) PrimaryBackupInput() *GoogleDnsRecordSetRoutingPolicyPrimaryBackup {
	var returns *GoogleDnsRecordSetRoutingPolicyPrimaryBackup
	_jsii_.Get(
		j,
		"primaryBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) Wrr() GoogleDnsRecordSetRoutingPolicyWrrList {
	var returns GoogleDnsRecordSetRoutingPolicyWrrList
	_jsii_.Get(
		j,
		"wrr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) WrrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wrrInput",
		&returns,
	)
	return returns
}


func NewGoogleDnsRecordSetRoutingPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDnsRecordSetRoutingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDnsRecordSetRoutingPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDnsRecordSetRoutingPolicyOutputReference_Override(g GoogleDnsRecordSetRoutingPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference)SetEnableGeoFencing(val interface{}) {
	if err := j.validateSetEnableGeoFencingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGeoFencing",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference)SetInternalValue(val *GoogleDnsRecordSetRoutingPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) PutGeo(value interface{}) {
	if err := g.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGeo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) PutPrimaryBackup(value *GoogleDnsRecordSetRoutingPolicyPrimaryBackup) {
	if err := g.validatePutPrimaryBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrimaryBackup",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) PutWrr(value interface{}) {
	if err := g.validatePutWrrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWrr",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ResetEnableGeoFencing() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableGeoFencing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		g,
		"resetGeo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ResetPrimaryBackup() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryBackup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ResetWrr() {
	_jsii_.InvokeVoid(
		g,
		"resetWrr",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

