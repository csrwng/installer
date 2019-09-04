package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"

	"github.com/openshift/installer/pkg/asset"
	hostedns "github.com/openshift/installer/pkg/asset/hosted/ns"
	assettls "github.com/openshift/installer/pkg/asset/tls"
)

// EtcdServerKey is the asset that generates the etcd client key/cert pair.
type EtcdServerKey struct {
	assettls.SignedCertKey
}

var _ asset.Asset = (*EtcdServerKey)(nil)

// Dependencies returns the dependency of the the cert/key pair, which includes
// the parent CA, and namespace.
func (a *EtcdServerKey) Dependencies() []asset.Asset {
	return []asset.Asset{
		&hostedns.Namespace{},
		&assettls.EtcdSignerCertKey{},
	}
}

// Generate generates the cert/key pair based on its dependencies.
func (a *EtcdServerKey) Generate(dependencies asset.Parents) error {
	ca := &assettls.EtcdSignerCertKey{}
	ns := &hostedns.Namespace{}
	dependencies.Get(ca, ns)

	dnsNames := []string{
		fmt.Sprintf("*.etcd.%s.svc", ns.ClusterNamespace),
		fmt.Sprintf("etcd-client.%s.svc", ns.ClusterNamespace),
		"localhost",
	}

	cfg := &assettls.CertCfg{
		Subject:   pkix.Name{CommonName: "etcd-server", OrganizationalUnit: []string{"etcd"}},
		DNSNames:  dnsNames,
		KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsages: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		Validity: assettls.ValidityTenYears,
	}

	return a.SignedCertKey.Generate(cfg, ca, "etcd-server", assettls.DoNotAppendParent)
}

// Name returns the human-friendly name of the asset.
func (a *EtcdServerKey) Name() string {
	return "Certificate (etcd-server)"
}
