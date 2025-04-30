package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterUserManagedKeysConfigOutputReference interface {
	cdktf.ComplexObject
	AggregationCa() *string
	SetAggregationCa(val *string)
	AggregationCaInput() *string
	ClusterCa() *string
	SetClusterCa(val *string)
	ClusterCaInput() *string
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
	ControlPlaneDiskEncryptionKey() *string
	SetControlPlaneDiskEncryptionKey(val *string)
	ControlPlaneDiskEncryptionKeyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EtcdApiCa() *string
	SetEtcdApiCa(val *string)
	EtcdApiCaInput() *string
	EtcdPeerCa() *string
	SetEtcdPeerCa(val *string)
	EtcdPeerCaInput() *string
	// Experimental.
	Fqn() *string
	GkeopsEtcdBackupEncryptionKey() *string
	SetGkeopsEtcdBackupEncryptionKey(val *string)
	GkeopsEtcdBackupEncryptionKeyInput() *string
	InternalValue() *GoogleContainerClusterUserManagedKeysConfig
	SetInternalValue(val *GoogleContainerClusterUserManagedKeysConfig)
	ServiceAccountSigningKeys() *[]*string
	SetServiceAccountSigningKeys(val *[]*string)
	ServiceAccountSigningKeysInput() *[]*string
	ServiceAccountVerificationKeys() *[]*string
	SetServiceAccountVerificationKeys(val *[]*string)
	ServiceAccountVerificationKeysInput() *[]*string
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
	ResetAggregationCa()
	ResetClusterCa()
	ResetControlPlaneDiskEncryptionKey()
	ResetEtcdApiCa()
	ResetEtcdPeerCa()
	ResetGkeopsEtcdBackupEncryptionKey()
	ResetServiceAccountSigningKeys()
	ResetServiceAccountVerificationKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterUserManagedKeysConfigOutputReference
type jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) AggregationCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) AggregationCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ClusterCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ClusterCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ControlPlaneDiskEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneDiskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ControlPlaneDiskEncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneDiskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) EtcdApiCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdApiCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) EtcdApiCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdApiCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) EtcdPeerCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdPeerCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) EtcdPeerCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdPeerCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GkeopsEtcdBackupEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeopsEtcdBackupEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GkeopsEtcdBackupEncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeopsEtcdBackupEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) InternalValue() *GoogleContainerClusterUserManagedKeysConfig {
	var returns *GoogleContainerClusterUserManagedKeysConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountSigningKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountSigningKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountSigningKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountSigningKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountVerificationKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountVerificationKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountVerificationKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountVerificationKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterUserManagedKeysConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterUserManagedKeysConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterUserManagedKeysConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterUserManagedKeysConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterUserManagedKeysConfigOutputReference_Override(g GoogleContainerClusterUserManagedKeysConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterUserManagedKeysConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetAggregationCa(val *string) {
	if err := j.validateSetAggregationCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregationCa",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetClusterCa(val *string) {
	if err := j.validateSetClusterCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterCa",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetControlPlaneDiskEncryptionKey(val *string) {
	if err := j.validateSetControlPlaneDiskEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneDiskEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetEtcdApiCa(val *string) {
	if err := j.validateSetEtcdApiCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etcdApiCa",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetEtcdPeerCa(val *string) {
	if err := j.validateSetEtcdPeerCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etcdPeerCa",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetGkeopsEtcdBackupEncryptionKey(val *string) {
	if err := j.validateSetGkeopsEtcdBackupEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gkeopsEtcdBackupEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetInternalValue(val *GoogleContainerClusterUserManagedKeysConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetServiceAccountSigningKeys(val *[]*string) {
	if err := j.validateSetServiceAccountSigningKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountSigningKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetServiceAccountVerificationKeys(val *[]*string) {
	if err := j.validateSetServiceAccountVerificationKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountVerificationKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetAggregationCa() {
	_jsii_.InvokeVoid(
		g,
		"resetAggregationCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetClusterCa() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetControlPlaneDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneDiskEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetEtcdApiCa() {
	_jsii_.InvokeVoid(
		g,
		"resetEtcdApiCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetEtcdPeerCa() {
	_jsii_.InvokeVoid(
		g,
		"resetEtcdPeerCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetGkeopsEtcdBackupEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeopsEtcdBackupEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetServiceAccountSigningKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountSigningKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ResetServiceAccountVerificationKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountVerificationKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterUserManagedKeysConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

