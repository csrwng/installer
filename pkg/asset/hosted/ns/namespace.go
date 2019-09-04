package ns

import (
	"os"

	survey "gopkg.in/AlecAivazis/survey.v1"

	"github.com/openshift/installer/pkg/asset"
)

type Namespace struct {
	ClusterNamespace string
}

var _ asset.Asset = (*Namespace)(nil)

// Dependencies returns no dependencies.
func (n *Namespace) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Generate queries for the cluster name from the user.
func (n *Namespace) Generate(parents asset.Parents) error {
	namespace := os.Getenv("NAMESPACE")
	if len(namespace) > 0 {
		n.ClusterNamespace = namespace
		return nil
	}
	return survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Cluster Namespace",
				Help:    "The namespace where the cluster will be created.  This will be used to generate etcd certificates.",
			},
			Validate: survey.Required,
		},
	}, &n.ClusterNamespace)
}

// Name returns the human-friendly name of the asset.
func (n *Namespace) Name() string {
	return "Cluster Namespace"
}
