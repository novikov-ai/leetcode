package valid_parentheses

/*
runtime: 0-2 ms (0 ms - without const, using hard-coded symbols)
memory: 2.1 MB
*/

const (
	CircleOpen   = "("
	CircleClosed = ")"
	SquareOpen   = "["
	SquareClosed = "]"
	CurlyOpen    = "{"
	CurlyClosed  = "}"
)

func IsValid(s string) bool {
	stack := make([]string, 0, len(s))

	for _, b := range s {
		str := string(b)

		switch str {
		case CircleClosed, SquareClosed, CurlyClosed:
			if len(stack) == 0 {
				return false
			}

		case CircleOpen, SquareOpen, CurlyOpen:
			stack = append(stack, str)
			continue
		}

		last := stack[len(stack)-1]
		switch last {
		case CircleOpen:
			if str != CircleClosed {
				return false
			}
		case SquareOpen:
			if str != SquareClosed {
				return false
			}
		case CurlyOpen:
			if str != CurlyClosed {
				return false
			}
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
