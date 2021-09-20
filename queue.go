package dasgo

type Queue struct {
	list LinkedList
}

func (q *Queue) Enqueue(newdata interface{}) {
	if !q.list.IsAppend {
		q.list.SetAppend()
	}
	q.list.Add(newdata)
}
func (q *Queue) Dequeue() interface{} {
	return q.list.DeleteHead()
}
func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}
func (q *Queue) Count() int {
	return q.list.Count()
}
func (q *Queue) Peek() interface{} {
	if q.list.IsEmpty() {
		return nil
	}
	return q.list.tail.next.data //Data at the head of the list
}
