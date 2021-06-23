#!/bin/bash
function init() {
    echo "Initializing.."
    readonly orig_cwd="$PWD"
    readonly script_path="${BASH_SOURCE[0]}"
    readonly script_dir="$(dirname $(readlink -f ${script_path}))"
    readonly deployment_yaml=${script_dir}/k8s/deployment.yaml
    readonly service_yaml=${script_dir}/k8s/service*   
    readonly hpa_yaml=${script_dir}/k8s/hpa.yaml
}

function sed_config() {
    echo "Updating config from envars..."
    readonly mongo_db_pass=$(echo ${{MONGO_PASS}} | base64 -d)
    readonly jwt_secret_keyd=$(echo ${{JWT_SECRET_KEY}} | base64 -d)
    readonly jwt_issd=$(echo ${{JWT_ISS}} | base64 -d)
    readonly jwt_member_secretd=$(echo ${{JWT_MEMBER_SECRET_KEY}} | base64 -d)
    readonly jwt_member_issd=$(echo ${{JWT_MEMBER_ISS}} | base64 -d)

    sed -i "s/{{log-param-file}}/${{LOG_PARAM_FILE}}/g" ${deployment_yaml}
    sed -i "s/{{timeout}}/${{TIMEOUT}}/g" ${deployment_yaml}

    sed -i "s/{{mongo-host}}/${{MONGO_HOST}}/g" ${deployment_yaml}
    sed -i "s/{{mongo-port}}/${{MONGO_PORT}}/g" ${deployment_yaml}
    sed -i "s/{{mongo-user}}/${{MONGO_USER}}/g" ${deployment_yaml}
    sed -i "s/{{mongo-pass}}/${mongo_db_pass}/g" ${deployment_yaml}
    sed -i "s/{{mongo-db-name}}/${{MONGO_DB_NAME}}/g" ${deployment_yaml}
    sed -i "s/{{mongo-db-auth}}/${{MONGO_DB_AUTH}}/g" ${deployment_yaml}
    sed -i "s/{{mongo-atlas}}/${{MONGO_ATLAS}}/g" ${deployment_yaml}

    sed -i "s/{{mongo-collection-user}}/${{MONGO_COLLECTION_USER}}/g" ${deployment_yaml}

    sed -i "s/{{jwt-secret-key}}/${jwt_secret_keyd}/g" ${deployment_yaml}
    sed -i "s/{{jwt-iss}}/${jwt_issd}/g" ${deployment_yaml}
    sed -i "s/{{jwt-member-secret-key}}/${jwt_member_secretd}/g" ${deployment_yaml}
    sed -i "s/{{jwt-member-iss}}/${jwt_member_issd}/g" ${deployment_yaml}

    sed -i "s/{{k8s-ns}}/${{K8S_NAMESPACE}}/g" ${deployment_yaml}
    sed -i "s/{{port}}/${{PORT}}/g" ${deployment_yaml}
    sed -i "s/{{go-env}}/${{GO_ENV}}/g" ${deployment_yaml}

    sed -i "s/{{k8s-ns}}/${{K8S_NAMESPACE}}/g" ${service_yaml}
    sed -i "s/{{port}}/${{PORT}}/g" ${service_yaml}

    sed -i "s/{{k8s-ns}}/${{K8S_NAMESPACE}}/g" ${hpa_yaml}
    sed -i "s/{{min-replicas}}/${{MIN_REPLICAS}}/g" ${hpa_yaml}
    sed -i "s/{{max-replicas}}/${{MAX_REPLICAS}}/g" ${hpa_yaml}
    sed -i "s/{{cpu-limit}}/${{CPU_LIMIT}}/g" ${deployment_yaml}
    sed -i "s/{{cpu-request}}/${{CPU_REQUEST}}/g" ${deployment_yaml}
    sed -i "s/{{memory-limit}}/${{MEMORY_LIMIT}}/g" ${deployment_yaml}
    sed -i "s/{{memory-request}}/${{MEMORY_REQUEST}}/g" ${deployment_yaml}
}

function deploy_dev() {
    echo "Deploying..."
    cat ${deployment_yaml}
    cat ${service_yaml}
    rsync -avz ${script_dir}/k8s/ deployer@35.219.4.100:/home/deployer/tmp/
    cat ${script_dir}/scripts/k8s_apply_step | ssh deployer@35.219.4.100
}

function deploy_prod() {
    echo "Deploying.."
    cat ${script_dir}/k8s/service-prod.yaml
    kubectl apply -f ${deployment_yaml}
    kubectl apply -f ${script_dir}/k8s/service-prod.yaml
    kubectl apply -f ${hpa_yaml}
}

function main() {
    init
    sed_config
    if [[ $TARGET == "DEV" ]]; then
        deploy_dev
    elif [[ $TARGET == "PROD" ]]; then
        deploy_prod
    else
        echo "No TARGET defined!"
        exit 1
    fi
}

main "$@"