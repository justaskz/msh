package repos

type Uid string

func (uid Uid) String() string {
	return string(uid)
}
