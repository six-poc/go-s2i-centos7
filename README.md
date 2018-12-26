# S2I Golang


## OpenShift
### create Image Stream for golang
$ oc creat -f openshift/is.yaml -n  openshift
 or 
$ oc import-image --from="bigg01/go-s2i-centos7" --scheduled=true golang --confirm


$ oc new-app https://github.com/bigg01/go-s2i-centos7.git --docker-image=golang:latest --context-dir=test-app