package types

type Stream string

func (s Stream) String() string {
	return string(s)
}

type Group string

func (g Group) String() string {
	return string(g)
}

type Consumer string

func (c Consumer) String() string {
	return string(c)
}

type Case string

func (c Case) String() string {
	return string(c)
}
