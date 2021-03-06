module github.com/opdev/podconfig-operator

go 1.13

require (
	github.com/containernetworking/plugins v0.8.7
	github.com/docker/libcontainer v2.2.1+incompatible
	github.com/go-logr/logr v0.1.0
	github.com/milosgajdos/tenus v0.0.3
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/vishvananda/netlink v1.1.0
	google.golang.org/grpc v1.27.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/cri-api v0.19.3
	sigs.k8s.io/controller-runtime v0.6.2
)
