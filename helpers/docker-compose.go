package helpers

import (
  "io/ioutil"
  "path"

  "gopkg.in/yaml.v2"
)

type DockerComposeMap map[string]DockerComposeService

type DockerComposeService struct {
  Build string
  Image string
}

func DockerCompose(root string) DockerComposeMap {
  return parseYAML(readYAML(path.Join(root, "docker-compose.yml")))
}

func parseYAML(s []byte) (d DockerComposeMap) {
  if err := yaml.Unmarshal(s, &d); err != nil {
    panic(err)
  }
  return
}

func readYAML(path string) []byte {
  if s, err := ioutil.ReadFile(path); err != nil {
    panic(err)
  } else {
    return s
  }
}