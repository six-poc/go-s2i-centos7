# S2I Golang


## OpenShift
### create Image Stream for golang

```sh
$ oc creat -f openshift/is.yaml -n  openshift
```
 or
```sh 
$ oc import-image --from="bigg01/go-s2i-centos7" --scheduled=true golang --confirm
```
```sh
$ oc new-app https://github.com/bigg01/go-s2i-centos7.git --docker-image=golang:latest --context-dir=test-app
```


# helm install 

https://blog.openshift.com/getting-started-helm-openshift/

```sh

export TILLER_NAMESPACE=tiller
$ oc new-project golang-example
Now using project "golang-example" on server "https://...".
...
Step 5: Grant the Tiller server edit access to the current project. The Tiller server will probably need at least “edit” access to each project where it will manage applications. In the case that Tiller will be handling Charts containing Role objects, “admin” access will be needed.

$ oc policy add-role-to-user edit "system:serviceaccount:${TILLER_NAMESPACE}:tiller" -n golang-example
Step 6: Install a Helm Chart. As an example, we’ll install the trusty OpenShift nodejs-ex sample application:

$ helm install ./helm/golang-ex/ --namespace golang-example

$ helm install --replace ./helm/golang-ex -n golang-example --namespace golang-example

$ helm install https://github.com/jim-minter/nodejs-ex/raw/helm/helm/nodejs-0.1.tgz  --namespace golang-example

```