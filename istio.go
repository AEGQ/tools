// Copyright 2018 Naftis Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"errors"
	"fmt"

	"log"

	"github.com/ghodss/yaml"
	"istio.io/istio/pilot/pkg/config/kube/crd"
	istiomodel "istio.io/istio/pilot/pkg/model"
)

type istiocrdExecutor struct {
	client *crd.Client
}

// NewCrdExecutor returns a istiocrd executor.
func NewCrdExecutor() *istiocrdExecutor {
	c, e := crd.NewClient("/root/.kube/config", "", istiomodel.IstioConfigTypes, "")
	if e != nil {
		log.Panic("[executor] init istiocrd fail", "error", e)
	}
	return &istiocrdExecutor{
		client: c,
	}
}

func (i *istiocrdExecutor) create(varr []istiomodel.Config) (errs error) {
	for _, config := range varr {
		var err error
		if config.Namespace, err = handleNamespaces(config.Namespace); err != nil {
			return err
		}

		var rev string
		if rev, err = i.client.Create(config); err != nil {
			// if the config create fail, break loop and return error
			fmt.Println("Created config fail", "config", config.Key(), "error", err)
			return err
		}
		fmt.Println("Created config success", "config", config.Key(), "revision", rev)
	}
	return nil
}

func (i *istiocrdExecutor) replace(varr []istiomodel.Config) (errs error) {

	for _, config := range varr {
		var err error
		if config.Namespace, err = handleNamespaces(config.Namespace); err != nil {
			return err
		}

		// fill up revision
		if config.ResourceVersion == "" {
			current := i.client.Get(config.Type, config.Name, "default")
			config.ResourceVersion = current.ResourceVersion
			// clear resourceVersion for rollback
			current.ResourceVersion = ""
		}
		var newRev string
		if newRev, err = i.client.Update(config); err != nil {
			// if the config create fail, break loop and return error
			fmt.Println("Replace config fail", "config", config.Key(), "error", err)
			return err
		}
		fmt.Println("Replace config success", "config", config.Key(), "revision", newRev)
	}

	return nil
}

func (i *istiocrdExecutor) delete(varr []istiomodel.Config) (errs error) {
	for _, config := range varr {
		var err error
		if config.Namespace, err = handleNamespaces(config.Namespace); err != nil {
			return err
		}

		if err := i.client.Delete(config.Type, config.Name, config.Namespace); err != nil {
			fmt.Println("Delete config fail", "config", config.Key(), "error", err)
			// if the config delete fail, continue loop
			//errs = multierror.Append(errs, fmt.Errorf("cannot delete %s: %v", config.Key(), err))
		} else {
			fmt.Println("Delete config success", "config", config.Key())
		}
	}
	return nil
}

var (
	namespace        string
	defaultNamespace = "default"
)

func handleNamespaces(objectNamespace string) (string, error) {
	if objectNamespace != "" && namespace != "" && namespace != objectNamespace {
		return "", fmt.Errorf(`the namespace from the provided object "%s" does `+
			`not match the namespace "%s". You must pass '--namespace=%s' to perform `+
			`this operation`, objectNamespace, namespace, objectNamespace)
	}

	if namespace != "" {
		return namespace, nil
	}

	if objectNamespace != "" {
		return objectNamespace, nil
	}
	return defaultNamespace, nil
}

func (i *istiocrdExecutor) apply(task string) (errs error) {

	// ignore k8s configuration. TODO support k8s configuration
	varr, _, err := crd.ParseInputs(task)
	if err != nil {
		return err
	}
	if len(varr) == 0 {
		return errors.New("nothing to execute")
	}

	if err := i.create(varr); err != nil {
		if err = i.replace(varr); err != nil {
			return err
		}
	}

	return
}

func (i *istiocrdExecutor) yamlOutput(configList []istiomodel.Config) string {
	buf := bytes.NewBuffer([]byte{})
	descriptor := i.client.ConfigDescriptor()
	for _, config := range configList {
		schema, exists := descriptor.GetByType(config.Type)
		if !exists {
			fmt.Printf("Unknown kind %q for %v", crd.ResourceName(config.Type), config.Name)
			continue
		}
		obj, err := crd.ConvertConfig(schema, config)
		if err != nil {
			fmt.Printf("Could not decode %v: %v", config.Name, err)
			continue
		}
		bytes, err := yaml.Marshal(obj)
		if err != nil {
			fmt.Printf("Could not convert %v to YAML: %v", config, err)
			continue
		}

		buf.Write(bytes)
		buf.WriteString("---")
	}

	return buf.String()
}

const (
	virtualservice = `
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: myapp
  namespace: dev
spec:
  hosts:
  - "ingressgateway.istio.paas.shein.io"
  gateways:
  - app-gateway
  http:
    - route:
      - destination:
          host: myapp-v1
        weight: 50
      - destination:
          host: myapp-v2
        weight: 50       
      timeout: 1s`
)

func main() {
	istio := NewCrdExecutor()
	istio.apply(virtualservice)
}
