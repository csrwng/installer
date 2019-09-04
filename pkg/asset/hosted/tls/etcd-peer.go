package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"

	"github.com/openshift/installer/pkg/asset"
	hostedns "github.com/openshift/installer/pkg/asset/hosted/ns"
	assettls "github.com/openshift/installer/pkg/asset/tls"
)

// EtcdPeerCertKey is the asset that generates the etcd client key/cert pair.
type EtcdPeerCertKey struct {
	assettls.SignedCertKey
}

var _ asset.Asset = (*EtcdPeerCertKey)(nil)

// Dependencies returns the dependency of the the cert/key pair, which includes
// the parent CA, and install config if it depends on the install config for
// DNS names, etc.
func (a *EtcdPeerCertKey) Dependencies() []asset.Asset {
	return []asset.Asset{
		&hostedns.Namespace{},
		&assettls.EtcdSignerCertKey{},
	}
}

// Generate generates the cert/key pair based on its dependencies.
func (a *EtcdPeerCertKey) Generate(dependencies asset.Parents) error {
	ca := &assettls.EtcdSignerCertKey{}
	ns := &hostedns.Namespace{}
	dependencies.Get(ca, ns)

	dnsNames := []string{
		fmt.Sprintf("*.etcd.%s.svc", ns.ClusterNamespace),
		fmt.Sprintf("*.etcd.%s.svc.cluster.local", ns.ClusterNamespace),
	}

	cfg := &assettls.CertCfg{
		Subject:   pkix.Name{CommonName: "etcd-member", OrganizationalUnit: []string{"etcd"}},
		DNSNames:  dnsNames,
		KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsages: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		Validity: assettls.ValidityTenYears,
	}

	return a.SignedCertKey.Generate(cfg, ca, "etcd-client", assettls.DoNotAppendParent)
}

// Name returns the human-friendly name of the asset.
func (a *EtcdPeerCertKey) Name() string {
	return "Certificate (etcd-peer)"
}
