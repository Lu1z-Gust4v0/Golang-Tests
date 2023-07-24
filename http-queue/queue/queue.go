package queue

import "errors"

type Queue struct {
  MaxSize int
  front int
  rear int
  values []int  
}

func NewQueue(maxSize int) *Queue {
  var values = make([]int, maxSize)

  return &Queue{
    MaxSize: maxSize,
    front: -1,
    rear: -1,
    values: values,
  }
}

func (queue *Queue) Enqueue(value int) error {
  if queue.IsFull() {
    return errors.New("Queue is full, cannot enqueue")
  } 
 
  if queue.IsEmpty() {
    queue.front++
  }

  queue.rear = (queue.rear + 1) % queue.MaxSize
  queue.values[queue.rear] = value

  return nil
}

func (queue *Queue) Dequeue() (int, error) {
  if queue.IsEmpty() {
    return 0, errors.New("Queue is empty, cannot dequeue")
  }
 
  var value int = queue.values[queue.front]
  
  if (queue.front == queue.rear) {
    queue.front = -1
    queue.rear = -1
    return value, nil
  }
  
  queue.front = (queue.front + 1) % queue.MaxSize

  return value, nil 
}

func (queue *Queue) IsFull() bool {
  if (queue.rear + 1) % queue.MaxSize == queue.front {
    return true
  }

  return false
}

func (queue *Queue) IsEmpty() bool {
  if queue.front == -1 {
    return true
  }
  return false
}

func (queue *Queue) GetFirst() (int, error) {
  if queue.IsEmpty() {
    return 0, errors.New("Queue is empty")
  }

  return queue.values[queue.front], nil
}

func (queue *Queue) GetLast() (int, error) {
  if queue.IsEmpty() {
    return 0, errors.New("Queue is empty")
  }

  return queue.values[queue.rear], nil
}
