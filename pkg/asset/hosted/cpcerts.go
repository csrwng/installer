package hosted

import (
	"crypto/x509"
	"crypto/x509/pkix"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/tls"
)

// KubeAPIServerInternalServiceCertKey is the asset that generates the kube-apiserver serving key/cert pair for SNI external load balancer.
type KubeAPIServerInternalServiceCertKey struct {
	tls.SignedCertKey
}

var _ asset.Asset = (*KubeAPIServerInternalServiceCertKey)(nil)

// Dependencies returns the dependency of the the cert/key pair
func (a *KubeAPIServerInternalServiceCertKey) Dependencies() []asset.Asset {
	return []asset.Asset{
		&tls.KubeAPIServerLBSignerCertKey{},
	}
}

// Generate generates the cert/key pair based on its dependencies.
func (a *KubeAPIServerInternalServiceCertKey) Generate(dependencies asset.Parents) error {
	ca := &tls.KubeAPIServerLBSignerCertKey{}
	dependencies.Get(ca)

	cfg := &tls.CertCfg{
		Subject:      pkix.Name{CommonName: "system:kube-apiserver", Organization: []string{"kube-master"}},
		KeyUsages:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		Validity:     tls.ValidityOneDay,
		DNSNames: []string{
			"kube-apiserver",
			"kube-apiserver.%s.svc",
		},
	}
	return a.SignedCertKey.Generate(cfg, ca, "kube-apiserver-svc-server", tls.AppendParent)
}

// Name returns the human-friendly name of the asset.
func (a *KubeAPIServerInternalServiceCertKey) Name() string {
	return "Certificate (kube-apiserver-svc-server)"
}
