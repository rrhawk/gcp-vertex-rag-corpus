variable "project_id" {
  description = "The project ID to deploy Vertex AI RAG resources"
  type        = string
}

variable "region" {
  description = "The region where resources are being created"
  type        = string
  default     = "us-central1"
}

variable "rag_corpuses" {
  description = "Map of RAG corpuses to create"
  type = map(object({
    description = optional(string)
    embedding_model_config = optional(object({
      model = optional(string, "text-embedding-004")
    }))
  }))
  default = {}
}
