/*
Copyright 2020 The OpenYurt Authors.

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

const (
	//todo
	YurtAppManagerNodePool = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.9
  creationTimestamp: null
  name: nodepools.apps.openyurt.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.type
    description: The type of nodepool
    name: Type
    type: string
  - JSONPath: .status.readyNodeNum
    description: The number of ready nodes in the pool
    name: ReadyNodes
    type: integer
  - JSONPath: .status.unreadyNodeNum
    name: NotReadyNodes
    type: integer
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: apps.openyurt.io
  names:
    categories:
    - all
    kind: NodePool
    listKind: NodePoolList
    plural: nodepools
    shortNames:
    - np
    singular: nodepool
  scope: Cluster
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: NodePool is the Schema for the nodepools API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: NodePoolSpec defines the desired state of NodePool
          properties:
            annotations:
              additionalProperties:
                type: string
              description: 'If specified, the Annotations will be added to all nodes. NOTE: existing labels with samy keys on the nodes will be overwritten.'
              type: object
            labels:
              additionalProperties:
                type: string
              description: 'If specified, the Labels will be added to all nodes. NOTE: existing labels with samy keys on the nodes will be overwritten.'
              type: object
            selector:
              description: A label query over nodes to consider for adding to the pool
              properties:
                matchExpressions:
                  description: matchExpressions is a list of label selector requirements. The requirements are ANDed.
                  items:
                    description: A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
                    properties:
                      key:
                        description: key is the label key that the selector applies to.
                        type: string
                      operator:
                        description: operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
                        type: string
                      values:
                        description: values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
                        items:
                          type: string
                        type: array
                    required:
                    - key
                    - operator
                    type: object
                  type: array
                matchLabels:
                  additionalProperties:
                    type: string
                  description: matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
                  type: object
              type: object
            taints:
              description: If specified, the Taints will be added to all nodes.
              items:
                description: The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
                type: object
              type: array
            type:
              description: The type of the NodePool
              type: string
          type: object
        status:
          description: NodePoolStatus defines the observed state of NodePool
          properties:
            nodes:
              description: The list of nodes' names in the pool
              items:
                type: string
              type: array
            readyNodeNum:
              description: Total number of ready nodes in the pool.
              format: int32
              type: integer
            unreadyNodeNum:
              description: Total number of unready nodes in the pool.
              format: int32
              type: integer
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
	YurtAppManagerUnitedDeployment = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.9
  creationTimestamp: null
  name: uniteddeployments.apps.openyurt.io
spec:
  additionalPrinterColumns:
  - JSONPath: .status.readyReplicas
    description: The number of pods ready.
    name: READY
    type: integer
  - JSONPath: .status.templateType
    description: The WorkloadTemplate Type.
    name: WorkloadTemplate
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
    name: AGE
    type: date
  group: apps.openyurt.io
  names:
    kind: UnitedDeployment
    listKind: UnitedDeploymentList
    plural: uniteddeployments
    shortNames:
    - ud
    singular: uniteddeployment
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: UnitedDeployment is the Schema for the uniteddeployments API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: UnitedDeploymentSpec defines the desired state of UnitedDeployment.
          properties:
            revisionHistoryLimit:
              description: Indicates the number of histories to be conserved. If unspecified, defaults to 10.
              format: int32
              type: integer
            selector:
              description: Selector is a label query over pods that should match the replica count. It must match the pod template's labels.
              properties:
                matchExpressions:
                  description: matchExpressions is a list of label selector requirements. The requirements are ANDed.
                  items:
                    description: A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
                    properties:
                      key:
                        description: key is the label key that the selector applies to.
                        type: string
                      operator:
                        description: operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
                        type: string
                      values:
                        description: values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
                        items:
                          type: string
                        type: array
                    required:
                    - key
                    - operator
                    type: object
                  type: array
                matchLabels:
                  additionalProperties:
                    type: string
                  description: matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
                  type: object
              type: object
            topology:
              description: Topology describes the pods distribution detail between each of pools.
              properties:
                pools:
                  description: Contains the details of each pool. Each element in this array represents one pool which will be provisioned and managed by UnitedDeployment.
                  items:
                    description: Pool defines the detail of a pool.
                    properties:
                      name:
                        description: Indicates pool name as a DNS_LABEL, which will be used to generate pool workload name prefix in the format '<deployment-name>-<pool-name>-'. Name should be unique between all of the pools under one UnitedDeployment. Name is NodePool Name
                        type: string
                      nodeSelectorTerm:
                        description: Indicates the node selector to form the pool. Depending on the node selector, pods provisioned could be distributed across multiple groups of nodes. A pool's nodeSelectorTerm is not allowed to be updated.
                        type: object
                      patch:
                        description: Indicates the patch for the templateSpec Now support strategic merge path :https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/#notes-on-the-strategic-merge-patch Patch takes precedence over Replicas fields If the Patch also modifies the Replicas, use the Replicas value in the Patch
                        type: object
                      replicas:
                        description: Indicates the number of the pod to be created under this pool.
                        format: int32
                        type: integer
                      tolerations:
                        description: Indicates the tolerations the pods under this pool have. A pool's tolerations is not allowed to be updated.
                        items:
                          description: The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.
                          type: object
                        type: array
                    required:
                    - name
                    type: object
                  type: array
              type: object
            workloadTemplate:
              description: WorkloadTemplate describes the pool that will be created.
              properties:
                deploymentTemplate:
                  description: Deployment template
                  properties:
                    metadata:
                      type: object
                    spec:
                      description: DeploymentSpec is the specification of the desired behavior of the Deployment.
                      type: object
                  required:
                  - spec
                  type: object
                statefulSetTemplate:
                  description: StatefulSet template
                  properties:
                    metadata:
                      type: object
                    spec:
                      description: A StatefulSetSpec is the specification of a StatefulSet.
                      type: object
                  required:
                  - spec
                  type: object
              type: object
          required:
          - selector
          type: object
        status:
          description: UnitedDeploymentStatus defines the observed state of UnitedDeployment.
          properties:
            collisionCount:
              description: Count of hash collisions for the UnitedDeployment. The UnitedDeployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
              format: int32
              type: integer
            conditions:
              description: Represents the latest available observations of a UnitedDeployment's current state.
              items:
                description: UnitedDeploymentCondition describes current state of a UnitedDeployment.
                properties:
                  lastTransitionTime:
                    description: Last time the condition transitioned from one status to another.
                    format: date-time
                    type: string
                  message:
                    description: A human readable message indicating details about the transition.
                    type: string
                  reason:
                    description: The reason for the condition's last transition.
                    type: string
                  status:
                    description: Status of the condition, one of True, False, Unknown.
                    type: string
                  type:
                    description: Type of in place set condition.
                    type: string
                type: object
              type: array
            currentRevision:
              description: CurrentRevision, if not empty, indicates the current version of the UnitedDeployment.
              type: string
            observedGeneration:
              description: ObservedGeneration is the most recent generation observed for this UnitedDeployment. It corresponds to the UnitedDeployment's generation, which is updated on mutation by the API Server.
              format: int64
              type: integer
            poolReplicas:
              additionalProperties:
                format: int32
                type: integer
              description: Records the topology detail information of the replicas of each pool.
              type: object
            readyReplicas:
              description: The number of ready replicas.
              format: int32
              type: integer
            replicas:
              description: Replicas is the most recently observed number of replicas.
              format: int32
              type: integer
            templateType:
              description: TemplateType indicates the type of PoolTemplate
              type: string
          required:
          - currentRevision
          - replicas
          - templateType
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
	YurtAppManagerYurtAppDaemon = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.9
  creationTimestamp: null
  name: yurtappdaemons.apps.openyurt.io
spec:
  additionalPrinterColumns:
  - JSONPath: .status.templateType
    description: The WorkloadTemplate Type.
    name: WorkloadTemplate
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
    name: AGE
    type: date
  group: apps.openyurt.io
  names:
    kind: YurtAppDaemon
    listKind: YurtAppDaemonList
    plural: yurtappdaemons
    shortNames:
    - yad
    singular: yurtappdaemon
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: YurtAppDaemon is the Schema for the YurtAppDaemon API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: YurtAppDaemonSpec defines the desired state of YurtAppDaemon.
          properties:
            nodepoolSelector:
              description: NodePoolSelector is a label query over nodepool that should match the replica count. It must match the nodepool's labels.
              properties:
                matchExpressions:
                  description: matchExpressions is a list of label selector requirements. The requirements are ANDed.
                  items:
                    description: A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
                    properties:
                      key:
                        description: key is the label key that the selector applies to.
                        type: string
                      operator:
                        description: operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
                        type: string
                      values:
                        description: values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
                        items:
                          type: string
                        type: array
                    required:
                    - key
                    - operator
                    type: object
                  type: array
                matchLabels:
                  additionalProperties:
                    type: string
                  description: matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
                  type: object
              type: object
            revisionHistoryLimit:
              description: Indicates the number of histories to be conserved. If unspecified, defaults to 10.
              format: int32
              type: integer
            selector:
              description: Selector is a label query over pods that should match the replica count. It must match the pod template's labels.
              properties:
                matchExpressions:
                  description: matchExpressions is a list of label selector requirements. The requirements are ANDed.
                  items:
                    description: A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
                    properties:
                      key:
                        description: key is the label key that the selector applies to.
                        type: string
                      operator:
                        description: operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
                        type: string
                      values:
                        description: values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
                        items:
                          type: string
                        type: array
                    required:
                    - key
                    - operator
                    type: object
                  type: array
                matchLabels:
                  additionalProperties:
                    type: string
                  description: matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
                  type: object
              type: object
            workloadTemplate:
              description: WorkloadTemplate describes the pool that will be created.
              properties:
                deploymentTemplate:
                  description: Deployment template
                  properties:
                    metadata:
                      type: object
                    spec:
                      description: DeploymentSpec is the specification of the desired behavior of the Deployment.
                      type: object
                  required:
                  - spec
                  type: object
                statefulSetTemplate:
                  description: StatefulSet template
                  properties:
                    metadata:
                      type: object
                    spec:
                      description: A StatefulSetSpec is the specification of a StatefulSet.
                      type: object
                  required:
                  - spec
                  type: object
              type: object
          required:
          - nodepoolSelector
          - selector
          type: object
        status:
          description: YurtAppDaemonStatus defines the observed state of YurtAppDaemon.
          properties:
            collisionCount:
              description: Count of hash collisions for the YurtAppDaemon. The YurtAppDaemon controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
              format: int32
              type: integer
            conditions:
              description: Represents the latest available observations of a YurtAppDaemon's current state.
              items:
                description: YurtAppDaemonCondition describes current state of a YurtAppDaemon.
                properties:
                  lastTransitionTime:
                    description: Last time the condition transitioned from one status to another.
                    format: date-time
                    type: string
                  message:
                    description: A human readable message indicating details about the transition.
                    type: string
                  reason:
                    description: The reason for the condition's last transition.
                    type: string
                  status:
                    description: Status of the condition, one of True, False, Unknown.
                    type: string
                  type:
                    description: Type of in place set condition.
                    type: string
                type: object
              type: array
            currentRevision:
              description: CurrentRevision, if not empty, indicates the current version of the YurtAppDaemon.
              type: string
            nodepools:
              description: NodePools indicates the list of node pools selected by YurtAppDaemon
              items:
                type: string
              type: array
            observedGeneration:
              description: ObservedGeneration is the most recent generation observed for this YurtAppDaemon. It corresponds to the YurtAppDaemon's generation, which is updated on mutation by the API Server.
              format: int64
              type: integer
            templateType:
              description: TemplateType indicates the type of PoolTemplate
              type: string
          required:
          - currentRevision
          - templateType
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
	YurtAppManagerYurtIngress = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.9
  creationTimestamp: null
  name: yurtingresses.apps.openyurt.io
spec:
  additionalPrinterColumns:
  - JSONPath: .status.nginx_ingress_controller_version
    description: The nginx ingress controller version
    name: Nginx-Ingress-Version
    type: string
  - JSONPath: .status.ingress_controller_replicas_per_pool
    description: The nginx ingress controller replicas per pool
    name: Replicas-Per-Pool
    type: integer
  - JSONPath: .status.readyNum
    description: The number of pools on which ingress is enabled
    name: ReadyNum
    type: integer
  - JSONPath: .status.unreadyNum
    description: The number of pools on which ingress is enabling or enable failed
    name: NotReadyNum
    type: integer
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: apps.openyurt.io
  names:
    categories:
    - all
    kind: YurtIngress
    listKind: YurtIngressList
    plural: yurtingresses
    shortNames:
    - ying
    singular: yurtingress
  scope: Cluster
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: YurtIngress is the Schema for the yurtingresses API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: YurtIngressSpec defines the desired state of YurtIngress
          properties:
            ingress_controller_replicas_per_pool:
              description: Indicates the number of the ingress controllers to be deployed under all the specified nodepools.
              format: int32
              type: integer
            pools:
              description: Indicates all the nodepools on which to enable ingress.
              items:
                description: IngressPool defines the details of a Pool for ingress
                properties:
                  name:
                    description: Indicates the pool name.
                    type: string
                required:
                - name
                type: object
              type: array
          type: object
        status:
          description: YurtIngressStatus defines the observed state of YurtIngress
          properties:
            conditions:
              description: Indicates all the nodepools on which to enable ingress.
              properties:
                ingressreadypools:
                  description: Indicates the pools that ingress controller is deployed successfully.
                  items:
                    type: string
                  type: array
                ingressunreadypools:
                  description: Indicates the pools that ingress controller is being deployed or deployed failed.
                  items:
                    description: IngressNotReadyPool defines the condition details of an ingress not ready Pool
                    properties:
                      name:
                        description: Indicates the pool name.
                        type: string
                      poolinfo:
                        description: Info of ingress not ready condition.
                        properties:
                          lastTransitionTime:
                            description: Last time the condition transitioned from one status to another.
                            format: date-time
                            type: string
                          message:
                            description: A human readable message indicating details about the transition.
                            type: string
                          reason:
                            description: The reason for the condition's last transition.
                            type: string
                          type:
                            description: Type of ingress not ready condition.
                            type: string
                        type: object
                    required:
                    - name
                    type: object
                  type: array
              type: object
            ingress_controller_replicas_per_pool:
              description: Indicates the number of the ingress controllers deployed under all the specified nodepools.
              format: int32
              type: integer
            nginx_ingress_controller_version:
              description: Indicates the nginx ingress controller version deployed under all the specified nodepools.
              type: string
            readyNum:
              description: Total number of ready pools on which ingress is enabled.
              format: int32
              type: integer
            unreadyNum:
              description: Total number of unready pools on which ingress is enabling or enable failed.
              format: int32
              type: integer
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
	YurtAppManagerRole = `
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: yurt-app-leader-election-role
  namespace: kube-system
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - configmaps/status
  verbs:
  - get
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
    `
	YurtAppManagerClusterRole = `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: yurt-app-manager-role
rules:
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - validatingwebhookconfigurations
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps
  resources:
  - controllerrevisions
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps
  resources:
  - deployments/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - apps
  resources:
  - statefulsets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps
  resources:
  - statefulsets/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - apps.openyurt.io
  resources:
  - nodepools
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps.openyurt.io
  resources:
  - nodepools/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - apps.openyurt.io
  resources:
  - uniteddeployments
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps.openyurt.io
  resources:
  - uniteddeployments/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - apps.openyurt.io
  resources:
  - yurtappdaemons
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps.openyurt.io
  resources:
  - yurtappdaemons/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - apps.openyurt.io
  resources:
  - yurtingresses
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps.openyurt.io
  resources:
  - yurtingresses/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - batch
  resources:
  - jobs
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - coordination.k8s.io
  resources:
  - leases
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - namespaces
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - persistentvolumeclaims
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterrolebindings
  verbs:
  - '*'
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterroles
  verbs:
  - '*'
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - rolebindings
  verbs:
  - '*'
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - roles
  verbs:
  - '*'
`
	YurtAppManagerRolebinding = `
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: yurt-app-leader-election-rolebinding
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: yurt-app-leader-election-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: kube-system
`
	YurtAppManagerClusterRolebinding = `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: yurt-app-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: yurt-app-manager-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: kube-system
`
	//todo
	YurtAppManagerSecret = `
apiVersion: v1
kind: Secret
metadata:
  name: yurt-app-webhook-certs
  namespace: kube-system
`
	YurtAppManagerService = `
apiVersion: v1
kind: Service
metadata:
  name: yurt-app-webhook-service
  namespace: kube-system
spec:
  ports:
  - port: 443
    targetPort: 9876
  selector:
    control-plane: yurt-app-manager
`
	YurtAppManagerDeployment = `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: yurt-app-manager
  name: yurt-app-manager
  namespace: kube-system
spec:
  replicas: 2
  selector:
    matchLabels:
      control-plane: yurt-app-manager
  template:
    metadata:
      labels:
        control-plane: yurt-app-manager
    spec:
      containers:
      - args:
        - --enable-leader-election
        - --v=4
        command:
        - /usr/local/bin/yurt-app-manager
        image: openyurt/yurt-app-manager:v0.5.0
        imagePullPolicy: Always
        name: manager
        ports:
        - containerPort: 9443
          name: webhook-server
          protocol: TCP
        volumeMounts:
        - mountPath: /tmp/k8s-webhook-server/serving-certs
          name: cert
          readOnly: true
      nodeSelector:
        beta.kubernetes.io/arch: amd64
        beta.kubernetes.io/os: linux
        openyurt.io/is-edge-worker: "false"
      priorityClassName: system-node-critical
      terminationGracePeriodSeconds: 10
      tolerations:
      - effect: NoSchedule
        key: node-role.alibabacloud.com/addon
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
        operator: Exists
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: yurt-app-webhook-certs
`
	//todo
	YurtAppManagerMutatingWebhookConfiguration = `
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  annotations:
    template: ""
  name: yurt-app-mutating-webhook-configuration
webhooks:
- clientConfig:
    caBundle: Cg==
    service:
      name: yurt-app-webhook-service
      namespace: kube-system
      path: /mutate-apps-openyurt-io-v1alpha1-nodepool
  failurePolicy: Fail
  name: mnodepool.kb.io
  rules:
  - apiGroups:
    - apps.openyurt.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - nodepools
- clientConfig:
    caBundle: Cg==
    service:
      name: yurt-app-webhook-service
      namespace: kube-system
      path: /mutate-apps-openyurt-io-v1alpha1-uniteddeployment
  failurePolicy: Fail
  name: muniteddeployment.kb.io
  rules:
  - apiGroups:
    - apps.openyurt.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - uniteddeployments
- clientConfig:
    caBundle: Cg==
    service:
      name: yurt-app-webhook-service
      namespace: kube-system
      path: /mutate-apps-openyurt-io-v1alpha1-yurtappdaemon
  failurePolicy: Fail
  name: myurtappdaemon.kb.io
  rules:
  - apiGroups:
    - apps.openyurt.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - yurtappdaemons
`
	//todo
	YurtAppManagerValidatingWebhookConfiguration = `
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  annotations:
    template: ""
  name: yurt-app-validating-webhook-configuration
webhooks:
- clientConfig:
    caBundle: Cg==
    service:
      name: yurt-app-webhook-service
      namespace: kube-system
      path: /validate-apps-openyurt-io-v1alpha1-nodepool
  failurePolicy: Fail
  name: vnodepool.kb.io
  rules:
  - apiGroups:
    - apps.openyurt.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    - DELETE
    resources:
    - nodepools
- clientConfig:
    caBundle: Cg==
    service:
      name: yurt-app-webhook-service
      namespace: kube-system
      path: /validate-apps-openyurt-io-v1alpha1-uniteddeployment
  failurePolicy: Fail
  name: vuniteddeployment.kb.io
  rules:
  - apiGroups:
    - apps.openyurt.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - uniteddeployments
- clientConfig:
    caBundle: Cg==
    service:
      name: yurt-app-webhook-service
      namespace: kube-system
      path: /validate-apps-openyurt-io-v1alpha1-yurtappdaemon
  failurePolicy: Fail
  name: vyurtappdaemon.kb.io
  rules:
  - apiGroups:
    - apps.openyurt.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - yurtappdaemons
- clientConfig:
    caBundle: Cg==
    service:
      name: yurt-app-webhook-service
      namespace: kube-system
      path: /validate-apps-openyurt-io-v1alpha1-yurtingress
  failurePolicy: Fail
  name: vyurtingress.kb.io
  rules:
  - apiGroups:
    - apps.openyurt.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    - DELETE
    resources:
    - yurtingresses
`
)
