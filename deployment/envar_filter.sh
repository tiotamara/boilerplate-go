#!/bin/bash
readonly orig_cwd="$PWD"
readonly script_path="${BASH_SOURCE[0]}"
readonly script_dir="$(dirname $(readlink -f ${script_path}))"

sed -i "s/{K8S_NAMESPACE}/${TARGET}__K8S_NAMESPACE/g" ${script_dir}/deploy.sh
sed -i "s/{PORT}/${TARGET}__PORT/g" ${script_dir}/deploy.sh
sed -i "s/{GO_ENV}/${TARGET}__GO_ENV/g" ${script_dir}/deploy.sh
sed -i "s/{LOG_PARAM_FILE}/${TARGET}__LOG_PARAM_FILE/g" ${script_dir}/deploy.sh
sed -i "s/{TIMEOUT}/${TARGET}__TIMEOUT/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_HOST}/${TARGET}__MONGO_HOST/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_PORT}/${TARGET}__MONGO_PORT/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_USER}/${TARGET}__MONGO_USER/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_PASS}/${TARGET}__MONGO_PASS/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_DB_NAME}/${TARGET}__MONGO_DB_NAME/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_DB_AUTH}/${TARGET}__MONGO_DB_AUTH/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_ATLAS}/${TARGET}__MONGO_ATLAS/g" ${script_dir}/deploy.sh
sed -i "s/{MONGO_COLLECTION_USER}/${TARGET}__MONGO_COLLECTION_USER/g" ${script_dir}/deploy.sh
sed -i "s/{JWT_SECRET_KEY}/${TARGET}__JWT_SECRET_KEY/g" ${script_dir}/deploy.sh
sed -i "s/{JWT_ISS}/${TARGET}__JWT_ISS/g" ${script_dir}/deploy.sh
sed -i "s/{JWT_MEMBER_SECRET_KEY}/${TARGET}__JWT_MEMBER_SECRET_KEY/g" ${script_dir}/deploy.sh
sed -i "s/{JWT_MEMBER_ISS}/${TARGET}__JWT_MEMBER_ISS/g" ${script_dir}/deploy.sh
sed -i "s/{MIN_REPLICAS}/${TARGET}__MIN_REPLICAS/g" ${script_dir}/deploy.sh
sed -i "s/{MAX_REPLICAS}/${TARGET}__MAX_REPLICAS/g" ${script_dir}/deploy.sh
sed -i "s/{CPU_LIMIT}/${TARGET}__CPU_LIMIT/g" ${script_dir}/deploy.sh
sed -i "s/{CPU_REQUEST}/${TARGET}__CPU_REQUEST/g" ${script_dir}/deploy.sh
sed -i "s/{MEMORY_LIMIT}/${TARGET}__MEMORY_LIMIT/g" ${script_dir}/deploy.sh
sed -i "s/{MEMORY_REQUEST}/${TARGET}__MEMORY_REQUEST/g" ${script_dir}/deploy.sh

cat ${script_dir}/deploy.sh
