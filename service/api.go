package service

type KubeApi interface {
	GetNamespaces() (*Namespace, error)
	GetPods(namespace string) (*Pod, error)
	GetDeployments(namespace string) (*Deployment, error)
	GetSecrets(namespace string) (*SecretList, error)
	GetSecret(namespace, id string) (*Secret, error)
	DeleteSecret(namespace, id string) (string, error)
	UpdateSecret(namespace, id string, data map[string]string) (string, error)
}

func CreateCliApi() KubeApi {
	return &KubeCliApi{}
}
