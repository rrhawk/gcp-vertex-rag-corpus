data "google_client_config" "default" {}

provider "vertexairag" {
  project      = var.project_id
  region       = var.region
  access_token = data.google_client_config.default.access_token
}

resource "vertexairag_rag_corpus" "corpus" {
  for_each = var.rag_corpuses

  display_name = each.key
  description  = each.value.description

  dynamic "embedding_model_config" {
    for_each = each.value.embedding_model_config != null ? [each.value.embedding_model_config] : []
    content {
      model = embedding_model_config.value.model
    }
  }
}
