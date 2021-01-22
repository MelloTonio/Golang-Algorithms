package Interview

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Get() string {
	return (*s)[len(*s)-1]
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isValidParenthesis(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var myStack Stack

	for _, v := range s {
		if v == '[' || v == '{' || v == '(' {
			myStack.Push(string(v))
		} else if v == ']' && !myStack.IsEmpty() && myStack.Get() == "[" {
			myStack.Pop()
		} else if v == '}' && !myStack.IsEmpty() && myStack.Get() == "{" {
			myStack.Pop()
		} else if v == ')' && !myStack.IsEmpty() && myStack.Get() == "(" {
			myStack.Pop()
		} else {
			return false
		}

	}

	return myStack.IsEmpty()

}
