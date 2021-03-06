/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package constants

import (
	"path/filepath"

	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	"k8s.io/kubernetes/pkg/util/homedir"
)

// MachineName is the name to use for the VM.
const MachineName = "minikubeVM"

// APIServerPort is the port that the API server should listen on.
const APIServerPort = 8443

// Fix for windows
var Minipath = filepath.Join(homedir.HomeDir(), ".minikube")

// TODO: Fix for windows
// KubeconfigPath is the path to the Kubernetes client config
var KubeconfigPath = clientcmd.RecommendedHomeFile

// MinikubeContext is the kubeconfig context name used for minikube
const MinikubeContext = "minikube"

// MakeMiniPath is a utility to calculate a relative path to our directory.
func MakeMiniPath(fileName ...string) string {
	args := []string{Minipath}
	args = append(args, fileName...)
	return filepath.Join(args...)
}

// Only pass along these flags to localkube.
var LogFlags = [...]string{
	"v",
	"vmodule",
}

var SupportedVMDrivers = [...]string{
	"virtualbox",
	"vmwarefusion",
	"kvm",
	"xhyve",
}

const (
	DefaultIsoUrl   = "https://storage.googleapis.com/minikube/minikube-0.5.iso"
	DefaultMemory   = 1024
	DefaultCPUS     = 1
	DefaultDiskSize = "20g"
	DefaultVMDriver = "virtualbox"
)

const (
	RemoteLocalKubeErrPath = "/var/lib/localkube/localkube.err"
	RemoteLocalKubeOutPath = "/var/lib/localkube/localkube.out"
)
