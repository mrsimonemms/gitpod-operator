FROM gitpod/workspace-go
# Install the Operator SDK
RUN curl -L "https://github.com/operator-framework/operator-sdk/releases/download/$(curl -s https://api.github.com/repos/operator-framework/operator-sdk/releases/latest | jq -r '.tag_name')/operator-sdk_linux_amd64" -o /tmp/operator-sdk \
 && sudo install -o root -g root -m 0755 /tmp/operator-sdk /usr/local/bin \
 && sudo sh -c "operator-sdk completion bash > /etc/bash_completion.d/operator-sdk" \
 && operator-sdk version
# Install Kubectl
RUN curl -L "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" -o /tmp/kubectl \
 && sudo install -o root -g root -m 0755 /tmp/kubectl /usr/local/bin/kubectl \
 && sudo sh -c "kubectl completion bash > /etc/bash_completion.d/kubectl" \
 && kubectl version --client
