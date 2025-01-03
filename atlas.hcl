variable "postgres_user" {
  type    = string
  default = getenv("DB_USER")
}

variable "postgres_password" {
  type    = string
  default = getenv("DB_PASSWORD")
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

env "dev" {
  src = data.external_schema.gorm.url
  url = "postgres://${var.postgres_user}:${var.postgres_password}@127.0.0.1:5432/vision?search_path=public&sslmode=disable"
  dev = "docker://postgres/latest/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
  exclude = [ "atlas_schema_revisions" ]
}
