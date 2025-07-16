#!/bin/bash

# ==============================================================================
#      UNIFIED DEPLOYMENT SCRIPT WITH SELECTABLE BUILD MODE
# ==============================================================================
# This script deploys an application to an existing GKE environment and allows
# the user to choose the build method.
#
# Usage:
#   ./run_deployment.sh --local   (Builds the Docker image on your machine)
#   ./run_deployment.sh --cloud   (Builds the Docker image using GCP Cloud Build)
#
# It ASSUMES:
# - A 'config.env' file exists.
# - The GKE cluster and Artifact Registry repo are already created.
# - For --cloud mode, a 'cloudbuild.yaml' file exists.
# ==============================================================================

# Stop script on any error
set -e

# --- HELPER FUNCTIONS ---
print_header() {
  echo ""
  echo "======================================================================"
  echo "  $1"
  echo "======================================================================"
}

# --- STEP 1: PARSE BUILD MODE & LOAD CONFIG ---
print_header "Step 1: Parsing Build Mode and Loading Configuration"

# Check for build mode argument
if [ -z "$1" ]; then
    echo "Error: Build mode not specified."
    echo "Usage: $0 --local OR $0 --cloud"
    exit 1
fi

BUILD_MODE=""
case "$1" in
    --local)
        BUILD_MODE="local"
        echo "Selected build mode: Local Docker"
        ;;
    --cloud)
        BUILD_MODE="cloud"
        echo "Selected build mode: GCP Cloud Build"
        ;;
    *)
        echo "Error: Invalid build mode '$1'."
        echo "Usage: $0 --local OR $0 --cloud"
        exit 1
        ;;
esac

# Load configuration from file
if [ -f config.env ]; then
  export $(grep -v '^#' config.env | xargs)
else
  echo "Error: Configuration file 'config.env' not found."
  exit 1
fi
FULL_IMAGE_NAME="${GCP_REGION}-docker.pkg.dev/${GCP_PROJECT}/${AR_REPO}/${IMAGE_NAME}:${IMAGE_TAG}"
echo "Configuration loaded for image: $FULL_IMAGE_NAME"


# --- STEP 2: BUILD AND PUSH IMAGE (CONDITIONAL) ---
if [ "$BUILD_MODE" == "local" ]; then
    print_header "Step 2: Building Image Locally with Docker"
    echo "Authenticating local Docker with Artifact Registry..."
    gcloud auth configure-docker ${GCP_REGION}-docker.pkg.dev

    echo "Building Docker image from local Dockerfile..."
    docker build -t $FULL_IMAGE_NAME .

    echo "Pushing Docker image to Artifact Registry..."
    docker push $FULL_IMAGE_NAME

elif [ "$BUILD_MODE" == "cloud" ]; then
    print_header "Step 2: Building Image with GCP Cloud Build"
    echo "Enabling the Cloud Build API (if not already enabled)..."
    gcloud services enable cloudbuild.googleapis.com

    echo "Submitting build to Cloud Build..."
    gcloud builds submit . \
        --config=cloudbuild.yaml \
        --substitutions=_FULL_IMAGE_NAME="$FULL_IMAGE_NAME"
fi
echo "Image build and push completed successfully."


# --- STEP 3: DEPLOY TO GKE (COMMON STEPS) ---
print_header "Step 3: Deploying to GKE Cluster"

echo "Connecting kubectl to cluster '$GKE_CLUSTER_NAME'..."
gcloud container clusters get-credentials $GKE_CLUSTER_NAME --region=$GCP_REGION

echo "Creating/updating Hugging Face token secret..."
kubectl delete secret generic hf-secret --ignore-not-found=true
kubectl create secret generic hf-secret --from-literal=token=$HF_READ_TOKEN

echo "Applying Kubernetes manifests..."
TEMP_DEPLOYMENT_FILE="kubernetes/02-deployment-final.yaml"
sed -e "s|IMAGE_PLACEHOLDER|${FULL_IMAGE_NAME}|g" \
    -e "s|HF_MODEL_ID_PLACEHOLDER|${HF_MODEL_ID}|g" \
    kubernetes/02-deployment.yaml > $TEMP_DEPLOYMENT_FILE

kubectl apply -f $TEMP_DEPLOYMENT_FILE
kubectl apply -f kubernetes/03-service.yaml

rm $TEMP_DEPLOYMENT_FILE
echo "Deployment manifests applied. Kubernetes will now perform a rolling update."


# --- STEP 4: MONITORING INSTRUCTIONS ---
print_header "Step 4: Monitoring the Rolling Update"
echo "Kubernetes is now updating your deployment to use the new image."
echo "Use these commands to monitor the status:"
echo ""
echo "1. Check the status of the rolling update:"
echo "   kubectl rollout status deployment/gemma-2b-deployment"
echo ""
echo "2. Watch the pods (a new one will be created, the old one will terminate):"
echo "   watch kubectl get pods -l app=gemma-2b-vllm"
echo ""
echo "3. Once the new pod is running, check its logs:"
echo '   kubectl logs -f $(kubectl get pods -l app=gemma-2b-vllm -o jsonpath="{.items[0].metadata.name}")'
echo ""
echo "======================= DEPLOYMENT SCRIPT COMPLETE ======================="