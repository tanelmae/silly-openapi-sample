#!/bin/bash
PACKAGE="gen"
GEN_DIR="pkg/${PACKAGE}"
mkdir -p ${GEN_DIR}

# Could be generated with single command
# but it is nice to have them in separate files
oapi-codegen -o ${GEN_DIR}/types.go -package ${PACKAGE} -generate types api.yaml
oapi-codegen -o ${GEN_DIR}/server.go -package ${PACKAGE} -generate server api.yaml
oapi-codegen -o ${GEN_DIR}/spec.go -package ${PACKAGE} -generate spec api.yaml
