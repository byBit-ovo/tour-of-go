type Node struct {
    Val int
    Next *Node
    Prev *Node
}

type MyLinkedList struct {
    Head *Node
    Size int
}


func Constructor() MyLinkedList {
    return MyLinkedList{Head:nil, Size:0}
}


func (this *MyLinkedList) Get(index int) int {
    if index >= this.Size {
        return -1
    }
    cur := this.Head
    for i:=0;i<index;i+=1 {
        cur = cur.Next
    }
    return cur.Val
    
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &Node{Val:val}
    if this.Head == nil{
        this.Head = newNode
        newNode.Next = newNode
        newNode.Prev = newNode
        this.Size += 1
        return
    }
    newNode.Next = this.Head
    newNode.Prev = this.Head.Prev
    this.Head.Prev.Next = newNode
    this.Head.Prev = newNode
    this.Head = newNode
    this.Size += 1
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &Node{Val:val}
    if this.Head == nil{
        this.Head = newNode
        newNode.Next = newNode
        newNode.Prev = newNode
        this.Size += 1
        return
    }
    newNode.Next = this.Head
    newNode.Prev = this.Head.Prev
    this.Head.Prev.Next = newNode
    this.Head.Prev = newNode
    this.Size += 1

}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.Size {
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }else if index == this.Size{
        this.AddAtTail(val)
        return
    }
    newNode := &Node{Val:val}
    prev := this.Head
    for i:=0;i<index-1;i+=1 {
        prev = prev.Next
    }
    newNode.Next = prev.Next
    prev.Next.Prev = newNode
    prev.Next = newNode
    newNode.Prev = prev
    this.Size += 1

}
func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.Size {
        return
    }
    if index == 0 {
        if this.Size == 1 {
            this.Head = nil
            this.Size--
            return
        }
        tail := this.Head.Prev
        this.Head = this.Head.Next
        this.Head.Prev = tail
        tail.Next = this.Head
        this.Size--
        return
    }
    prev := this.Head
    for i := 0; i < index-1; i++ {
        prev = prev.Next
    }
    target := prev.Next
    prev.Next = target.Next
    target.Next.Prev = prev
    this.Size--
}



/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */