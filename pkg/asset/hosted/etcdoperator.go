package hosted

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/openshift/installer/data"
	"github.com/openshift/installer/pkg/asset"
	hostedns "github.com/openshift/installer/pkg/asset/hosted/ns"
)

const (
	hostedDir = "hosted"
)

type EtcdOperator struct {
	files []*asset.File
}

// Name returns a human friendly name for the hosted cluster asset
func (o *EtcdOperator) Name() string {
	return "Etcd Operator Manifests"
}

// Dependencies returns all of the dependencies directly needed by the
// hosted cluster asset
func (o *EtcdOperator) Dependencies() []asset.Asset {
	return []asset.Asset{
		&hostedns.Namespace{},
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (o *EtcdOperator) Generate(dependencies asset.Parents) error {

	ns := &hostedns.Namespace{}
	dependencies.Get(ns)

	// Process static files
	for _, f := range []string{
		"etcd-cluster-role.yaml",
		"etcd-cluster.yaml",
		"etcd-operator.yaml",
	} {
		data, err := getFileContents(path.Join("manifests/hosted", f))
		if err != nil {
			return err
		}
		o.files = append(o.files, &asset.File{
			Filename: path.Join(hostedDir, f),
			Data:     data,
		})
	}

	//  Process templates
	templateData := map[string]string{"Namespace": ns.ClusterNamespace}
	for _, f := range []string{
		"etcd-cluster-role-binding.yaml",
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
func (o *EtcdOperator) Files() []*asset.File {
	return o.files
}

// Load returns the openshift asset from disk.
func (c *EtcdOperator) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}

// getFileContents the content of the given URI, assuming that it's a file
func getFileContents(uri string) ([]byte, error) {
	file, err := data.Assets.Open(uri)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
