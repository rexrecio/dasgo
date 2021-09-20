package dasgo

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(newdata interface{}) {
	s.list.Add(newdata)
}
func (s *Stack) Pop() interface{} {
	return s.list.DeleteHead()
}
func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}
func (s *Stack) Count() int {
	return s.list.Count()
}
func (s *Stack) Peek() interface{} {
	if s.list.IsEmpty() {
		return nil
	}
	return s.list.tail.next.data //Data at the head of the list
}
