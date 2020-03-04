# instructional designer senior policy
# These are senior rules, so the instructional designer policies plus these
path "v1/data/lms/capture/programs" {
  capabilities = ["delete"]
}

path "v1/data/lms/capture/courses" {
  capabilities = ["delete"]
}

path "v1/data/lms/capture/modules" {
  capabilities = ["delete"]
}

path "v1/data/lms/capture/sessions" {
  capabilities = ["delete"]
}

path "v1/data/lms/capture/quizz" {
  capabilities = ["delete"]
}

# Publishdata/ capabilities
path "v1/data/lms/capture/programs/publish" {
  capabilities = ["read"]
}

path "v1/data/lms/capture/courses/publish" {
  capabilities = ["read"]
}

path "v1/data/lms/capture/modules/publish" {
  capabilities = ["read"]
}

path "v1/data/lms/capture/sessions/publish" {
  capabilities = ["read"]
}
