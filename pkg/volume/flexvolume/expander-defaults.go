/*
Copyright 2018 The Kubernetes Authors.

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

package flexvolume

import (
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
)

type expanderDefaults struct {
	plugin *flexVolumePlugin
}

func newExpanderDefaults(plugin *flexVolumePlugin) *expanderDefaults {
	return &expanderDefaults{plugin}
}

func (e *expanderDefaults) ExpandVolumeDevice(spec *volume.Spec, newSize resource.Quantity, oldSize resource.Quantity) (resource.Quantity, error) {
	glog.Warning(logPrefix(e.plugin), "using default expand for volume ", spec.Name(), ", to size ", newSize, " from ", oldSize)
	return newSize, nil
}

// the defaults for ExpandFS return a generic resize indicator that will trigger the operation executor to go ahead with
// generic filesystem resize
func (e *expanderDefaults) ExpandFS(spec *volume.Spec, devicePath, deviceMountPath string, _, _ resource.Quantity) error {
	glog.Warning(logPrefix(e.plugin), "using default filesystem resize for volume ", spec.Name(), ", at ", devicePath)
	_, err := util.GenericResizeFS(e.plugin.host, e.plugin.GetPluginName(), devicePath, deviceMountPath)
	return err
}
