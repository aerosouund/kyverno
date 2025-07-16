# Openreports demo

## **Environment creation**

- (OPTIONAL) Install reports-server: 
Install the latest reports server release. and replace the image with this  `aerosouund/reports-server:0.1.0`  . (no official release candidate has been created).  
And use the following role:  
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app.kubernetes.io/instance: reports-server
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: reports-server
    app.kubernetes.io/version: v0.1.4-rc.1
    helm.sh/chart: reports-server-0.1.4-rc.1
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
    rbac.authorization.k8s.io/aggregate-to-view: "true"
  name: reports-server
rules:
- apiGroups:
  - reports.kyverno.io
  resources:
  - ephemeralreports
  - clusterephemeralreports
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
  - deletecollection
- apiGroups:
  - ""
  resources:
  - namespaces
  verbs:
  - get
- apiGroups:
  - apiregistration.k8s.io
  resources:
  - apiservices
  verbs:
  - create
- apiGroups:
  - apiregistration.k8s.io
  resourceNames:
  - v1.reports.kyverno.io
  - v1alpha2.wgpolicyk8s.io
  - v1alpha1.openreports.io
  resources:
  - apiservices
  verbs:
  - get
  - delete
  - update
  - patch
- apiGroups:
  - wgpolicyk8s.io
  resources:
  - policyreports
  - policyreports/status
  - clusterpolicyreports
  - clusterpolicyreports/status
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
  - deletecollection
- apiGroups:
  - openreports.io
  resources:
  - reports
  - reports/status
  - clusterreports
  - clusterreports/status
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
  - deletecollection
- apiGroups:
  - ""
  - events.k8s.io
  resources:
  - events
  verbs:
  - create
  - patch
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
```

- Install the latest kyverno chart and ensure all the new resources are installed (generating, mutating, validating policies).

- Ensure that the role of the reports controller is the following:  
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: kyverno:reports-controller:core
  labels:
    app.kubernetes.io/component: reports-controller
    app.kubernetes.io/instance: kyverno
    app.kubernetes.io/part-of: kyverno
    app.kubernetes.io/version: latest
rules:
  - apiGroups:
      - apiextensions.k8s.io
    resources:
      - customresourcedefinitions
    verbs:
      - get
  - apiGroups:
      - ''
    resources:
      - configmaps
      - namespaces
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - kyverno.io
    resources:
      - globalcontextentries
      - globalcontextentries/status
      - policyexceptions
      - policies
      - clusterpolicies
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch
      - deletecollection
  - apiGroups:
      - policies.kyverno.io
    resources:
      - validatingpolicies
      - validatingpolicies/status
      - imagevalidatingpolicies
      - imagevalidatingpolicies/status
      - generatingpolicies
      - mutatingpolicies
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch
      - deletecollection
  - apiGroups:
      - policies.kyverno.io
    resources:
      - policyexceptions
      - policyexceptions/status
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - admissionregistration.k8s.io
    resources:
      - validatingadmissionpolicies
      - validatingadmissionpolicybindings
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - reports.kyverno.io
    resources:
      - ephemeralreports
      - clusterephemeralreports
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch
      - deletecollection
  - apiGroups:
      - wgpolicyk8s.io
    resources:
      - policyreports
      - policyreports/status
      - clusterpolicyreports
      - clusterpolicyreports/status
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch
      - deletecollection
  - apiGroups:
      - openreports.io
    resources:
      - reports
      - reports/status
      - clusterreports
      - clusterreports/status
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch
      - deletecollection
  - apiGroups:
      - ''
      - events.k8s.io
    resources:
      - events
    verbs:
      - create
      - patch
```


- Change the reports server deployment to use the image  `aerosouund/reports-controller:v0.1.1`  . This is needed because the 1.15 release candidate has not been built yet. And use the following flags, of course the other flags apart from the openreports one have no effect on functionality. But this is to ensure stability and that you get an environment that matches what i had in the demo  

```yaml
        - --disableMetrics=false
        - --openreportsEnabled=true
        - --otelConfig=prometheus
        - --metricsPort=8000
        - --resyncPeriod=15m
        - --admissionReports=true
        - --aggregateReports=true
        - --policyReports=true
        - --backgroundScan=true
        - --validatingAdmissionPolicyReports=false
        - --backgroundScanWorkers=2
        - --backgroundScanInterval=1h
        - --skipResourceFilters=true
        - --enableConfigMapCaching=true
        - --enableDeferredLoading=true
        - --maxAPICallResponseLength=2000000
        - --loggingFormat=text
        - --v=2
        - --omitEvents=PolicyApplied,PolicySkipped
        - --enablePolicyException=false
        - --allowInsecureRegistry=false
        - --registryCredentialHelpers=default,google,amazon,azure,github
        - --enableReporting=validate,mutate,mutateExisting,imageVerify,generate
```
- Install the latest policy reporter helm chart  `3.3.1`  . it should work out of the box. to ensure it doesn't filter any namespace or source. remove the config file flag from the deployment. or remove the source filters section in the file  
```yaml
        - --port=8080
        - --config=/app/config.yaml <<<< remove this
        - --dbfile=/sqlite/database.db
        - --metrics-enabled=false
        - --rest-enabled=false
        - --profile=false
        - --lease-name=policy-reporter
        - --template-dir=/app/templates
```

## Demo

- In a terminal, run this  `kubectl logs -f deploy/policy-reporter -n policy-reporter`

- Open another terminal, and create a targetconfig resource. e.x:  
```yaml
apiVersion: policyreporter.kyverno.io/v1alpha1
kind: TargetConfig
metadata:
  name: webhook-example
spec:
  webhook:
    webhook: "http://your-server-ip:3000"
    headers: 
      foo: bar
    skipTLS: true
```

- Create a policy, ideally a policy that will generate reports for existing resources. e.x:  
```yaml
apiVersion: kyverno.io/v1
kind: Policy
metadata:
  name: require-app-label
  namespace: default
spec:
  admission: true
  background: true
  rules:
  - match:
      resources:
        kinds:
        - Pod
    name: check-app-label
    skipBackgroundRequests: true
    validate:
      message: Pods must have an 'app' label.
      pattern:
        metadata:
          labels:
            appu: ?*
  validationFailureAction: enforce
```
- Wait for a few seconds, the reports controller should generate openreports  `Report`  resources, which should be picked up by the policy-reporter and sent to the designated data store configured in the targetconfig

- If you had the reports server installed, confirm that resources weren't stored in ETCD. run:  
`kubectl get --raw=/metrics | grep apiserver_storage_objects | grep openre`
- (OPTIONAL) Port forward your postgres container, and connect to it. if you used the default installation of the reports server the password should be  `reports`  

```bash
psql --host localhost --username postgres
\c reportsdb;
SELECT * FROM reports;
```