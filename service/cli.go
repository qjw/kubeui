package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	getNamespacesP  = "get namespaces --output=json"
	getPodsP        = "get pods -n %s --output=json"
	getDeploymentsP = "get deploy -n %s --output=json"
	getSecretsP     = "get secret -n %s --output=json"
	getSecretP      = "get secret -n %s %s --output=json"
	deleteSecretP   = "delete secret -n %s %s"
	updateResourceP = "apply -f -"
)

type KubeCliApi struct {
}

func (this KubeCliApi) kubectlInput(arg, input string) (string, error) {
	args := strings.Split(arg, " ")
	cmd := exec.Command("kubectl", args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	in := func(str string) func(io.WriteCloser) {
		return func(stdin io.WriteCloser) {
			defer stdin.Close()
			io.Copy(stdin, bytes.NewBufferString(str))
		}
	}

	populate_stdin_func := in(input)
	populate_stdin_func(stdin)
	out, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (this KubeCliApi) kubectlRaw(arg string) (string, error) {
	args := strings.Split(arg, " ")
	out, err := exec.Command("kubectl", args...).Output()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(out), nil
}

func (this KubeCliApi) kubectl(arg string, obj interface{}) error {
	args := strings.Split(arg, " ")
	out, err := exec.Command("kubectl", args...).Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := json.Unmarshal(out, obj); err != nil {
		fmt.Println(err.Error())
		return err
	}

	// for debug
	//	out, err = json.MarshalIndent(obj, "\t\t\t\t", "  ")
	//	fmt.Println(string(out))
	// end debug
	return nil
}

func (this KubeCliApi) GetNamespaces() (*Namespace, error) {
	var obj Namespace
	if err := this.kubectl(getNamespacesP, &obj); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return &obj, nil
}

func (this KubeCliApi) GetPods(namespace string) (*Pod, error) {
	var obj Pod
	if err := this.kubectl(fmt.Sprintf(getPodsP, namespace), &obj); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return &obj, nil
}

func (this KubeCliApi) GetDeployments(namespace string) (*Deployment, error) {
	var obj Deployment
	if err := this.kubectl(fmt.Sprintf(getDeploymentsP, namespace), &obj); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return &obj, nil
}

func (this KubeCliApi) GetSecrets(namespace string) (*SecretList, error) {
	var obj SecretList
	if err := this.kubectl(fmt.Sprintf(getSecretsP, namespace), &obj); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return &obj, nil
}

func (this KubeCliApi) GetSecret(namespace, id string) (*Secret, error) {
	var obj Secret
	if err := this.kubectl(fmt.Sprintf(getSecretP, namespace, id), &obj); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}

	return &obj, nil
}

func (this KubeCliApi) DeleteSecret(namespace, id string) (string, error) {
	if out, err := this.kubectlRaw(fmt.Sprintf(deleteSecretP, namespace, id)); err != nil {
		fmt.Printf(err.Error())
		return out, err
	} else {
		return out, nil
	}
}

func (this KubeCliApi) UpdateSecret(namespace, id string, data map[string]string) (string, error) {
	var obj Secret
	if err := this.kubectl(fmt.Sprintf(getSecretP, namespace, id), &obj); err != nil {
		fmt.Printf(err.Error())
		return err.Error(), err
	}
	if len(data) < 1 {
		err := fmt.Errorf("can not set empty secret with namespace %s id %s", namespace, id)
		return err.Error(), err
	}

	// 清空老的
	//	for k := range obj.Data {
	//		delete(obj.Data, k)
	//	}

	for k, v := range data {
		obj.Data[k] = base64.StdEncoding.EncodeToString([]byte(v))
	}

	if out, err := json.MarshalIndent(obj, "\t", "  "); err != nil {
		fmt.Printf(err.Error())
		return err.Error(), err
	} else {
		// fmt.Println(string(out))
		if output, err := this.kubectlInput(updateResourceP, string(out)); err != nil {
			fmt.Printf(err.Error())
			return output, err
		} else {
			return output, nil
		}
	}
}
