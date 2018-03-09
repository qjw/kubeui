package service

type Namespace struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Items      []struct {
		Kind     string `json:"kind,omitempty"`
		Metadata struct {
			Name              string            `json:"name"`
			CreationTimestamp string            `json:"creationTimestamp"`
			SelfLink          string            `json:"selfLink"`
			Labels            map[string]string `json:"labels"`
			Annotations       map[string]string `json:"annotations"`
		} `json:"metadata"`
	} `json:"items"`
}

type Env struct {
	Name      string `json:"name"`
	Value     string `json:"value,omitempty"`
	ValueFrom *struct {
		FieldRef *struct {
			ApiVersion string `json:"apiVersion,omitempty"`
			FieldPath  string `json:"fieldPath,omitempty"`
		} `json:"fieldRef,omitempty"`
		SecretKeyRef *struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"secretKeyRef,omitempty"`
	} `json:"valueFrom,omitempty"`
}

type Volume struct {
	Name     string `json:"name"`
	EmptyDir *struct {
		SizeLimit string `json:"sizeLimit,omitempty"`
	} `json:"emptyDir,omitempty"`
	ConfigMap *struct {
		DefaultMode int    `json:"defaultMode,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"configMap,omitempty"`
	Secret *struct {
		DefaultMode int    `json:"defaultMode,omitempty"`
		SecretName  string `json:"secretName,omitempty"`
	} `json:"secret,omitempty"`
}

type Container struct {
	Name            string `json:"name"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
	NodeName        string `json:"nodeName,omitempty"`
	ServiceAccount  string `json:"serviceAccount,omitempty"`
	RestartPolicy   string `json:"restartPolicy,omitempty"`
	Env             []Env  `json:"env,omitempty"`
	Ports           []struct {
		Name          string `json:"name,omitempty"`
		ContainerPort int    `json:"containerPort,omitempty"`
		Protocol      string `json:"protocol,omitempty"`
	} `json:"ports,omitempty"`
	VolumeMounts []struct {
		Name      string `json:"name,omitempty"`
		MountPath string `json:"mountPath,omitempty"`
		ReadOnly  bool   `json:"readOnly,omitempty"`
	} `json:"volumeMounts,omitempty"`
}

type ContainerSpec struct {
	Containers     []Container `json:"containers"`
	Volumes        []Volume    `json:"volumes,omitempty"`
	ServiceAccount string      `json:"serviceAccount,omitempty"`
	SchedulerName  string      `json:"schedulerName,omitempty"`
	RestartPolicy  string      `json:"restartPolicy,omitempty"`
	DnsPolicy      string      `json:"dnsPolicy,omitempty"`
}

type Pod struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Items      []struct {
		Kind     string `json:"kind,omitempty"`
		Metadata struct {
			Name      string            `json:"name"`
			NameSpace string            `json:"namespace"`
			Labels    map[string]string `json:"labels"`
		} `json:"metadata"`
		Spec ContainerSpec `json:"spec"`
	} `json:"items"`
}

type Deployment struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Items      []struct {
		Kind     string `json:"kind,omitempty"`
		Metadata struct {
			Name      string            `json:"name"`
			NameSpace string            `json:"namespace"`
			Labels    map[string]string `json:"labels"`
		} `json:"metadata"`
		Spec struct {
			Replicas int `json:"replicas"`
			Template struct {
				Spec ContainerSpec `json:"spec"`
			} `json:"template"`
		} `json:"spec"`
	} `json:"items"`
}

type SecretList struct {
	ApiVersion string       `json:"apiVersion,omitempty"`
	Items      []SecretLite `json:"items"`
}

type SecretLite struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Type       string `json:"type,omitempty"`
	Metadata   struct {
		Name        string            `json:"name"`
		NameSpace   string            `json:"namespace"`
		Labels      map[string]string `json:"labels,omitempty"`
		Annotations map[string]string `json:"annotations,omitempty"`
	} `json:"metadata"`
}

type Secret struct {
	SecretLite
	Data map[string]string `json:"data,omitempty"`
}
