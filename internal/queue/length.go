package queue

type Length int

const (
	Short = iota
	Long
)

func ToLength(txt string) (Length, error) {
	if txt == "short" {
		return Short, nil
	} else if txt == "long" {
		return Long, nil
	}

	return 0, nil
}
