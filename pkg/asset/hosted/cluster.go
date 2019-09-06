package hosted

import (
	"github.com/openshift/installer/pkg/asset"
)

type Cluster struct {
	files []*asset.File
}

// Name returns a human friendly name for the hosted cluster asset
func (c *Cluster) Name() string {
	return "Hosted Cluster"
}

// Dependencies returns all of the dependencies directly needed by the
// hosted cluster asset
func (c *Cluster) Dependencies() []asset.Asset {
	return []asset.Asset{
		&BootstrapFiles{},
		&EtcdOperator{},
		&EtcdSecrets{},
		&ControlPlaneSecrets{},
		&RenderingScripts{},
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (c *Cluster) Generate(dependencies asset.Parents) error {
	etcdOperator := &EtcdOperator{}
	etcdSecrets := &EtcdSecrets{}
	bootstrap := &BootstrapFiles{}
	scripts := &RenderingScripts{}
	cpSecrets := &ControlPlaneSecrets{}

	dependencies.Get(etcdOperator, etcdSecrets, bootstrap, scripts, cpSecrets)

	files := []*asset.File{}
	files = append(files, etcdOperator.Files()...)
	files = append(files, etcdSecrets.Files()...)
	files = append(files, bootstrap.Files()...)
	files = append(files, scripts.Files()...)
	files = append(files, cpSecrets.Files()...)
	c.files = files

	return nil
}

// Files returns the manifests generated for the cluster
func (c *Cluster) Files() []*asset.File {
	return c.files
}

// Load returns the openshift asset from disk.
func (c *Cluster) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}
