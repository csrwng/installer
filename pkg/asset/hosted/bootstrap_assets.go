package hosted

import (
	"path"

	igntypes "github.com/coreos/ignition/config/v2_2/types"
	"github.com/vincent-petithory/dataurl"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/ignition/bootstrap"
)

type BootstrapFiles struct {
	files []*asset.File
}

// Name returns a human friendly name for the hosted cluster asset
func (b *BootstrapFiles) Name() string {
	return "Bootstrap Files"
}

// Dependencies returns all of the dependencies directly needed by the
// hosted cluster asset
func (b *BootstrapFiles) Dependencies() []asset.Asset {
	return []asset.Asset{
		&bootstrap.Bootstrap{},
	}
}

// Generate creates the manifests needed to launch a hosted cluster
func (b *BootstrapFiles) Generate(dependencies asset.Parents) error {

	bt := &bootstrap.Bootstrap{}
	dependencies.Get(bt)

	for _, file := range bt.Config.Storage.Files {
		b.files = append(b.files, assetFileFromIgnitionFile(file))
	}

	return nil
}

// Files returns the manifests generated for the cluster
func (b *BootstrapFiles) Files() []*asset.File {
	return b.files
}

// Load returns the openshift asset from disk.
func (b *BootstrapFiles) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}

func assetFileFromIgnitionFile(file igntypes.File) *asset.File {
	u, err := dataurl.DecodeString(file.Contents.Source)
	if err != nil {
		panic(err.Error()) // Should not happen
	}
	return &asset.File{
		Filename: path.Join("bootstrap", file.Path),
		Data:     u.Data,
	}
}
