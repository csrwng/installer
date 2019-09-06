package hosted

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/hosted/tls"
	assettls "github.com/openshift/installer/pkg/asset/tls"
)

type EtcdSecrets struct {
	files []*asset.File
}

// Name returns a human friendly name for the hosted cluster asset
func (o *EtcdSecrets) Name() string {
	return "Etcd Operator Secrets"
}

// Dependencies returns all of the dependencies directly needed by the
// hosted cluster asset
func (o *EtcdSecrets) Dependencies() []asset.Asset {
	return []asset.Asset{
		&assettls.EtcdSignerClientCertKey{},
		&assettls.EtcdSignerCertKey{},
		&tls.EtcdServerKey{},
		&tls.EtcdPeerCertKey{},
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (o *EtcdSecrets) Generate(dependencies asset.Parents) error {
	clientKey := &assettls.EtcdSignerClientCertKey{}
	signerCert := &assettls.EtcdSignerCertKey{}
	serverKey := &tls.EtcdServerKey{}
	peerKey := &tls.EtcdPeerCertKey{}
	dependencies.Get(clientKey, signerCert, serverKey, peerKey)

	templateData := map[string]string{
		"PeerCert":   base64.StdEncoding.EncodeToString(peerKey.Cert()),
		"PeerKey":    base64.StdEncoding.EncodeToString(peerKey.Key()),
		"ServerCert": base64.StdEncoding.EncodeToString(serverKey.Cert()),
		"ServerKey":  base64.StdEncoding.EncodeToString(serverKey.Key()),
		"ClientCert": base64.StdEncoding.EncodeToString(clientKey.Cert()),
		"ClientKey":  base64.StdEncoding.EncodeToString(clientKey.Key()),
		"EtcdCA":     base64.StdEncoding.EncodeToString(signerCert.Cert()),
	}

	for _, f := range []string{
		"etcd-peer-tls-secret.yaml",
		"etcd-server-tls-secret.yaml",
		"etcd-client-tls-secret.yaml",
	} {
		data, err := getFileContents(path.Join("manifests/hosted", fmt.Sprintf("%s.template", f)))
		if err != nil {
			return err
		}
		o.files = append(o.files, &asset.File{
			Filename: path.Join(hostedDir, f),
			Data:     applyTemplateData(data, templateData),
		})
	}
	return nil
}

// Files returns the manifests generated for the cluster
func (o *EtcdSecrets) Files() []*asset.File {
	return o.files
}

// Load returns the openshift asset from disk.
func (c *EtcdSecrets) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}

var templateFuncs = map[string]interface{}{
	"indent": indent,
}

func indent(indention int, v []byte) string {
	newline := "\n" + strings.Repeat(" ", indention)
	return strings.Replace(string(v), "\n", newline, -1)
}

func applyTemplateData(data []byte, templateData interface{}) []byte {
	template := template.Must(template.New("template").Funcs(templateFuncs).Parse(string(data)))
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, templateData); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
