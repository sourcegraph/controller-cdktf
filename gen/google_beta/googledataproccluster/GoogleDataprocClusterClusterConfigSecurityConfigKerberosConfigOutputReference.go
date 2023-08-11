package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataproccluster/internal"
)

type GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference interface {
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
	CrossRealmTrustAdminServer() *string
	SetCrossRealmTrustAdminServer(val *string)
	CrossRealmTrustAdminServerInput() *string
	CrossRealmTrustKdc() *string
	SetCrossRealmTrustKdc(val *string)
	CrossRealmTrustKdcInput() *string
	CrossRealmTrustRealm() *string
	SetCrossRealmTrustRealm(val *string)
	CrossRealmTrustRealmInput() *string
	CrossRealmTrustSharedPasswordUri() *string
	SetCrossRealmTrustSharedPasswordUri(val *string)
	CrossRealmTrustSharedPasswordUriInput() *string
	EnableKerberos() interface{}
	SetEnableKerberos(val interface{})
	EnableKerberosInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig
	SetInternalValue(val *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig)
	KdcDbKeyUri() *string
	SetKdcDbKeyUri(val *string)
	KdcDbKeyUriInput() *string
	KeyPasswordUri() *string
	SetKeyPasswordUri(val *string)
	KeyPasswordUriInput() *string
	KeystorePasswordUri() *string
	SetKeystorePasswordUri(val *string)
	KeystorePasswordUriInput() *string
	KeystoreUri() *string
	SetKeystoreUri(val *string)
	KeystoreUriInput() *string
	KmsKeyUri() *string
	SetKmsKeyUri(val *string)
	KmsKeyUriInput() *string
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
	RootPrincipalPasswordUri() *string
	SetRootPrincipalPasswordUri(val *string)
	RootPrincipalPasswordUriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TgtLifetimeHours() *float64
	SetTgtLifetimeHours(val *float64)
	TgtLifetimeHoursInput() *float64
	TruststorePasswordUri() *string
	SetTruststorePasswordUri(val *string)
	TruststorePasswordUriInput() *string
	TruststoreUri() *string
	SetTruststoreUri(val *string)
	TruststoreUriInput() *string
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
	ResetCrossRealmTrustAdminServer()
	ResetCrossRealmTrustKdc()
	ResetCrossRealmTrustRealm()
	ResetCrossRealmTrustSharedPasswordUri()
	ResetEnableKerberos()
	ResetKdcDbKeyUri()
	ResetKeyPasswordUri()
	ResetKeystorePasswordUri()
	ResetKeystoreUri()
	ResetRealm()
	ResetTgtLifetimeHours()
	ResetTruststorePasswordUri()
	ResetTruststoreUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustAdminServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustAdminServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustKdcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustKdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustRealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustRealmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) CrossRealmTrustSharedPasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustSharedPasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) EnableKerberosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKerberosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) InternalValue() *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig {
	var returns *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKeyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKeyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KdcDbKeyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcDbKeyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeyPasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystorePasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystorePasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystoreUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KeystoreUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKeyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) KmsKeyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) RootPrincipalPasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPrincipalPasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TgtLifetimeHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tgtLifetimeHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePasswordUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePasswordUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststorePasswordUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststorePasswordUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststoreUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) TruststoreUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"truststoreUriInput",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustAdminServer(val *string) {
	if err := j.validateSetCrossRealmTrustAdminServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustAdminServer",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustKdc(val *string) {
	if err := j.validateSetCrossRealmTrustKdcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustKdc",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustRealm(val *string) {
	if err := j.validateSetCrossRealmTrustRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustRealm",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetCrossRealmTrustSharedPasswordUri(val *string) {
	if err := j.validateSetCrossRealmTrustSharedPasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustSharedPasswordUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetEnableKerberos(val interface{}) {
	if err := j.validateSetEnableKerberosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKerberos",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKdcDbKeyUri(val *string) {
	if err := j.validateSetKdcDbKeyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcDbKeyUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeyPasswordUri(val *string) {
	if err := j.validateSetKeyPasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPasswordUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystorePasswordUri(val *string) {
	if err := j.validateSetKeystorePasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystorePasswordUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKeystoreUri(val *string) {
	if err := j.validateSetKeystoreUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystoreUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetKmsKeyUri(val *string) {
	if err := j.validateSetKmsKeyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetRootPrincipalPasswordUri(val *string) {
	if err := j.validateSetRootPrincipalPasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPrincipalPasswordUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTgtLifetimeHours(val *float64) {
	if err := j.validateSetTgtLifetimeHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tgtLifetimeHours",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststorePasswordUri(val *string) {
	if err := j.validateSetTruststorePasswordUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststorePasswordUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference)SetTruststoreUri(val *string) {
	if err := j.validateSetTruststoreUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truststoreUri",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustAdminServer() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustAdminServer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustKdc() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustKdc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustRealm() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustRealm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetCrossRealmTrustSharedPasswordUri() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossRealmTrustSharedPasswordUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetEnableKerberos() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableKerberos",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKdcDbKeyUri() {
	_jsii_.InvokeVoid(
		g,
		"resetKdcDbKeyUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeyPasswordUri() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyPasswordUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystorePasswordUri() {
	_jsii_.InvokeVoid(
		g,
		"resetKeystorePasswordUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetKeystoreUri() {
	_jsii_.InvokeVoid(
		g,
		"resetKeystoreUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetRealm() {
	_jsii_.InvokeVoid(
		g,
		"resetRealm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTgtLifetimeHours() {
	_jsii_.InvokeVoid(
		g,
		"resetTgtLifetimeHours",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststorePasswordUri() {
	_jsii_.InvokeVoid(
		g,
		"resetTruststorePasswordUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ResetTruststoreUri() {
	_jsii_.InvokeVoid(
		g,
		"resetTruststoreUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

