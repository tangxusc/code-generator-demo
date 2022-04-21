set -o errexit
set -o nounset
set -o pipefail

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
../vendor/k8s.io/code-generator/generate-groups.sh \
  all \
  github.com/stream-stack/common/crd/generated \
  ../crd/ \
  storeset:v1 \
  --go-header-file $(pwd)/boilerplate.go.txt \
  --output-base $(pwd)/../