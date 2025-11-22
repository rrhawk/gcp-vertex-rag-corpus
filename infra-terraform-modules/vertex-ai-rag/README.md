# Vertex AI RAG Corpus Terraform Module

This module manages Vertex AI RAG Corpuses using a custom Terraform provider.

## Prerequisites

### Custom Provider Installation

Since this module relies on a custom provider (`rrhawk/vertexairag`) that is not published to the public Terraform Registry, you must build and install it on the machine running Terraform/Terragrunt.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/rrhawk/gcp-vertex-rag-corpus.git
    cd gcp-vertex-rag-corpus
    ```

2.  **Build and Install the Provider:**
    ```bash
    make install
    ```
    This will build the provider from source and install it to `~/.terraform.d/plugins/registry.terraform.io/rrhawk/vertexairag/0.1.0/linux_amd64/`.

## Usage with Terragrunt (External Git Source)

To use this module in your `terragrunt.hcl` from another repository:

```hcl
terraform {
  # Point to the module in the git repository
  source = "git::https://github.com/rrhawk/gcp-vertex-rag-corpus.git//infra-terraform-modules/vertex-ai-rag?ref=main"
}

inputs = {
  project_id = "your-project-id"
  region     = "us-central1"

  rag_corpuses = {
    "my-corpus" = {
      description = "My RAG Corpus"
      embedding_model_config = {
        model = "text-embedding-004"
      }
    }
  }
}
```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| `project_id` | The project ID to deploy Vertex AI RAG resources | `string` | n/a | yes |
| `region` | The region where resources are being created | `string` | `us-central1` | no |
| `rag_corpuses` | Map of RAG corpuses to create | `map(object)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| `rag_corpuses` | Map of created RAG corpuses |
