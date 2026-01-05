package repo

import (
  "fmt"
  "strings"
)

type Repo struct {
  Owner   Owner
  Name    Name
  Version Version
  State   State
}

func New(owner Owner, name Name, version Version) (Repo, error) {
  repo := Repo{
    Owner:   owner,
    Name:    name,
    Version: version,
    State:   States.Missing,
  }

  if err := repo.validate(); err != nil {
    return Repo{}, err
  }

  return repo, nil
}

func Parse(uid Uid) (Repo, error) {
  parts := strings.Split(uid.String(), "__")

  if len(parts) != 3 {
    return Repo{}, fmt.Errorf("folder '%s' format is incorrect", uid)
  }

  repo, err := New(Owner(parts[0]), Name(parts[1]), Version(parts[2]))
  if err != nil {
    return Repo{}, err
  }

  return repo, nil
}

func (repo Repo) Uid() Uid {
  elements := []string{repo.Owner.String(), repo.Name.String(), repo.Version.String()}
  uid := strings.Join(elements, "__")
  return Uid(uid)
}

func (repo Repo) SetState(state State) Repo {
  repo.State = state
  return repo
}

func (repo Repo) validate() error {
  fields := map[string]string{
    "Owner":   string(repo.Owner),
    "Name":    string(repo.Name),
    "Version": string(repo.Version),
  }

  for fieldName, value := range fields {
    if value == "" {
      return fmt.Errorf("%s is undefined", fieldName)
    }
  }

  return nil
}
