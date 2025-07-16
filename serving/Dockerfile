# ----------------------------------------------------------------------------------
# STAGE 1: The Base Image
# This is the starting point. It's like choosing the OS for a new computer.
# ----------------------------------------------------------------------------------
# We use an official image from the vLLM team. This saves us from a world of pain.
# It comes pre-loaded with:
#   - A specific version of Ubuntu Linux
#   - The correct NVIDIA CUDA Toolkit (e.g., 12.1)
#   - The correct version of PyTorch compiled against that CUDA version
#   - All of vLLM's Python dependencies
# 'vllm-openai' means it includes the OpenAI-compatible API server.
FROM vllm/vllm-openai:v0.4.2

# ----------------------------------------------------------------------------------
# STAGE 2: Set Environment Variables (Optional but good practice)
# These are variables that will exist inside the running container.
# ----------------------------------------------------------------------------------
# ENV HUGGING_FACE_HUB_TOKEN=""

# ----------------------------------------------------------------------------------
# STAGE 3: The Entrypoint / Command
# This is the most important part. It tells Docker what to run when the
# container is started from this image.
# ----------------------------------------------------------------------------------
# We use CMD, which stands for "command". It can be overridden from the
# Kubernetes manifest if needed.
# This is the equivalent of running the following in your terminal:
#
#   python -m vllm.entrypoints.openai.api_server \
#       --model your-username/my-llama3-finetuned \
#       --tensor-parallel-size 1 \
#       --host 0.0.0.0 \
#       --port 8000
#
CMD ["python", "-m", "vllm.entrypoints.openai.api_server", \
     # --model: This tells vLLM which model to download from Hugging Face Hub.
     #          This is why we don't need the model files in our image!
     #          You MUST replace this with your actual Hugging Face repo ID.
     "--model", "unsloth/gemma-2b-bnb-4bit", \
     \
     # --tensor-parallel-size: How many GPUs to use for ONE copy of the model.
     #                         Since our GKE node has one GPU, this must be 1.
     "--tensor-parallel-size", "1", \
     \
     # --host 0.0.0.0: Tells the server to listen for connections on all available
     #                 network interfaces inside the container, not just localhost.
     #                 This is essential for it to be reachable from outside.
     "--host", "0.0.0.0", \
     \
     # --port 8000: The network port inside the container that the server will listen on.
     #              Our Kubernetes Service will route traffic to this port.
     "--port", "8000"]