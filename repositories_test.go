package main

import (
  "testing"
  "io/ioutil"
	"os"
	"os/exec"
)

func TestCheckRepo(t *testing.T) {
  tempDir, err := ioutil.TempDir("", "get")
  
  if err != nil {
    t.Fatal("Could not create tempDir")
  }

	cmd := exec.Command("git", "init", "test")
	cmd.Dir = tempDir
  if cmd.Run() != nil {
    t.Fatal("Could not create originRepo")
	}

  configuration := Configuration{
    Username: "alice",
    Token: "secret",
    Path: tempDir,
  }
  env := Env{
    Config: configuration,
  }
  repo := Repo{
    FullName: "alice/example",
    SSHUrl: tempDir + "/test",
    HTTPSUrl: tempDir + "/test",
  }

  // First call to checkRepo() should clone the given valid repo
  out := "clone"
  res := checkRepo(repo, env)
  if res != out {
    t.Errorf("checkRepo() = '%s', want '%s'", res, out)
  }

  // Subsequent calls to checkRepo() should fetch the given valid repo
  out = "fetch"
  res = checkRepo(repo, env)
  if res != out {
    t.Errorf("checkRepo() = '%s', want '%s'", res, out)
  }

  os.RemoveAll(tempDir)
}
