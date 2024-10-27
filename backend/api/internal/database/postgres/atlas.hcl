env "local" {
  src = "file://schema.sql"
  dev = "docker://postgres/16/dev?search_path=public"
  migration {
    dir    = "file://migrations"
    format = "golang-migrate"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
