/*
 * Copyright 2022 Holoinsight Project Authors. Licensed under Apache-2.0.
 */

package k8sutils

import (
	"github.com/traas-stack/holoinsight-agent/pkg/k8s/k8slabels"
	v1 "k8s.io/api/core/v1"
	"strings"
)

// IsSandbox checks whether the container described by the input parameter is a pod's sandbox.
func IsSandbox(containerName string, k8sContainerName string, labels map[string]string) bool {
	if strings.HasPrefix(containerName, "/k8s_POD") {
		return true
	}
	if k8sContainerName == "POD" {
		return true
	}
	if labels["io.kubernetes.docker.type"] == "podsandbox" {
		return true
	}
	if labels["io.kubernetes.pouch.type"] == "sandbox" {
		return true
	}
	return false
}

// IsInitContainer checks whether the container is the init container of the pod
func IsInitContainer(pod *v1.Pod, containerLabels map[string]string) bool {
	containerName := k8slabels.GetContainerName(containerLabels)
	for i := range pod.Spec.InitContainers {
		if pod.Spec.InitContainers[i].Name == containerName {
			return true
		}
	}
	return false
}
