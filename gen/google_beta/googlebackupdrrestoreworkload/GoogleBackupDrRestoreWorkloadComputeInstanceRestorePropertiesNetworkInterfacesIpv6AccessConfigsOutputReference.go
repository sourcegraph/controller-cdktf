package googlebackupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrrestoreworkload/internal"
)

type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference interface {
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
	ExternalIp() *string
	SetExternalIp(val *string)
	ExternalIpInput() *string
	ExternalIpv6() *string
	SetExternalIpv6(val *string)
	ExternalIpv6Input() *string
	ExternalIpv6PrefixLength() *float64
	SetExternalIpv6PrefixLength(val *float64)
	ExternalIpv6PrefixLengthInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkTier() *string
	SetNetworkTier(val *string)
	NetworkTierInput() *string
	PublicPtrDomainName() *string
	SetPublicPtrDomainName(val *string)
	PublicPtrDomainNameInput() *string
	SetPublicPtr() interface{}
	SetSetPublicPtr(val interface{})
	SetPublicPtrInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetExternalIp()
	ResetExternalIpv6()
	ResetExternalIpv6PrefixLength()
	ResetName()
	ResetNetworkTier()
	ResetPublicPtrDomainName()
	ResetSetPublicPtr()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference
type jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalIpv6PrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ExternalIpv6PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalIpv6PrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) NetworkTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) NetworkTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) PublicPtrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicPtrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) PublicPtrDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicPtrDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) SetPublicPtr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setPublicPtr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) SetPublicPtrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setPublicPtrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference_Override(g GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetExternalIp(val *string) {
	if err := j.validateSetExternalIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIp",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetExternalIpv6(val *string) {
	if err := j.validateSetExternalIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIpv6",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetExternalIpv6PrefixLength(val *float64) {
	if err := j.validateSetExternalIpv6PrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIpv6PrefixLength",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetNetworkTier(val *string) {
	if err := j.validateSetNetworkTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTier",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetPublicPtrDomainName(val *string) {
	if err := j.validateSetPublicPtrDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicPtrDomainName",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetSetPublicPtr(val interface{}) {
	if err := j.validateSetSetPublicPtrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setPublicPtr",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetExternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetExternalIpv6() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalIpv6",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetExternalIpv6PrefixLength() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalIpv6PrefixLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetNetworkTier() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkTier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetPublicPtrDomainName() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicPtrDomainName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetSetPublicPtr() {
	_jsii_.InvokeVoid(
		g,
		"resetSetPublicPtr",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

