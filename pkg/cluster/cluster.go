package cluster

import (
	"context"

	"github.com/puppetlabs/horsehead/v2/workdir"
	"github.com/puppetlabs/relay/pkg/dialog"
	"k8s.io/apimachinery/pkg/runtime"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	ClusterName     = "relay-workflows"
	NetworkName     = "relay-workflows-net"
	WorkerCount     = 2
	HostStorageName = "local-storage"

	DefaultRegistryPort         = 5000
	DefaultLoadBalancerHostPort = 8080
	DefaultLoadBalancerNodePort = 80
)

type ClientOptions struct {
	Scheme *runtime.Scheme
}

type Config struct {
	WorkDir *workdir.WorkDir
	dialog.Dialog
}

type CreateOptions struct {
	// LoadBalancerHostPort is the port on the host to bind to when mapping
	// between the host machine and the service load balancer node.
	LoadBalancerHostPort int
	// ImageRegistryPort is the port to use on both the host AND the nodes for
	// the container image registry service. This isn't required, to be the
	// same, but to avoid confusion when tagging images on the host and
	// actually using them inside the cluster, we use the same port for both.
	// This allows someone to tag and push an image with
	// `localhost:5000/my-image:latest` and use the same repo/image/tag
	// combination from k8s resources.
	ImageRegistryPort int
}

// Manager provides methods to manage the lifecycle of a cluster.
type Manager interface {
	Exists(ctx context.Context) (bool, error)
	Create(ctx context.Context, opts CreateOptions) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Delete(ctx context.Context) error
	ImportImages(ctx context.Context, images ...string) error
	GetKubeconfig(ctx context.Context) (*clientcmdapi.Config, error)
	WriteKubeconfig(ctx context.Context, path string) error
	GetClient(ctx context.Context, opts ClientOptions) (*Client, error)
}

// NewManager returns a new selected Manager. Since k3d is
// the only one supported right now, we just return a manager
// that one. Using this function ensures the delegate manager
// always satisfies the Manager interface.
func NewManager(cfg Config) Manager {
	return NewK3dClusterManager(cfg)
}
