package singly_linked_list

/****
ADT 单链表
Data
    单链表可以将每个元素视为一个node节点 每个node 节点包含 数据域 以及指针域例如
    |数据域|指针域|
    | 12  |0x12x|
    单链表数据集合为 {node1,node2,...node(n)} n 个节点链接成一个链表。
    1.单链表的第一个节点的存储位置叫做头节点
    2.单链表最后一个节点指针为空 null
    3.单链表的指针域存放的是下一个节点的指针地址
Operation
  LinkListInit()  		初始化方法
  IsEmpty()  			判断是不是一个空链表
  GetLength() 			获取量表长度 O(n)
  Add() 				添加头 O(1)
  Append() 				尾部添加 O(n)
  Insert()              指定位置添加 O(n)
  GetAllElement()       获取所有元素值O(n)
  DelElement()          删除指定元素O(n)
  RemoveElement()       删除指定元素O(n)
endADT
****/

type ElementType int

type Node struct {
	Data ElementType
	Next *Node
}

type LinkList struct {
	headNode *Node
}

//LinkListInit  初始化
func LinkListInit() (m *LinkList) {
	m = new(LinkList)
	return
}

//IsEmpty 判断是不是一个空链表
func (p *LinkList) IsEmpty() (b bool) {
	if p.headNode == nil {
		b = true
		return
	}
	return
}

//GetLength  获取长度
func (p *LinkList) GetLength() (n int) {
	cur := p.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	n = count
	return
}

//Add 添加头
func (p *LinkList) Add(data ElementType) {
	node := &Node{
		Data: data,
	}
	if p.IsEmpty() {
		p.headNode = node
		return
	}
	node.Next = p.headNode
	p.headNode = node
	return
}

//Append 尾添加
func (p *LinkList) Append(data ElementType) {
	node := &Node{
		Data: data,
		Next: nil,
	}
	if p.IsEmpty() {
		p.headNode = node
	} else {
		cur := p.headNode
		i := 0
		for cur.Next != nil {
			cur = cur.Next
			i++
		}
		cur.Next = node
	}
	return
}

//Insert 指定位置添加
func (p *LinkList) Insert(i int, data ElementType) {
	if i <= 0 { // 如果小于0 头添加
		p.Add(data)
	} else if i > p.GetLength() { // 如果大于长度尾部添加
		p.Append(data)
	} else {
		node := &Node{
			Data: data,
		}
		count := 0
		// 获取当前节点位置信息
		cur := p.headNode
		for cur.Next != nil && count < (i-1) {
			cur = cur.Next
			count++
		}
		node.Next = cur.Next
		cur.Next = node
	}
	return
}

//GetAllElement 获取所有元素
func (p *LinkList) GetAllElement() (data []ElementType) {
	if p.IsEmpty() {
		return
	}
	cur := p.headNode
	for cur.Next != nil {
		data = append(data, cur.Data)
		cur = cur.Next
	}
	data = append(data, cur.Data)
	return
}

//DelElement 删除指定元素
func (p *LinkList) DelElement(i int) {
	if p.IsEmpty() {
		return
	}
	if i == 0 {
		node := p.headNode.Next
		p.headNode = node
		return
	}
	if p.GetLength() < i {
		return
	}
	cur := p.headNode
	j := 0
	upNode := p.headNode
	for cur.Next != nil && j < i {
		if j == (i - 1) {
			upNode = cur
		}
		cur = cur.Next
		j++
	}
	if p.GetLength()-1 == i {
		upNode.Next = nil
	} else {
		upNode.Next = cur
	}
	return
}

//RemoveElement 删除指定元素
func (p *LinkList) RemoveElement(i int) {
	if p.IsEmpty() {
		return
	}
	if p.GetLength() < i {
		return
	}
	if i <= 0 {
		node := p.headNode.Next
		p.headNode = node
		return
	}
	cur := p.headNode
	j := 0
	for cur.Next != nil && j < (i-1) {
		cur = cur.Next
		j++
	}
	if p.GetLength()-1 == i {
		cur.Next = nil
	} else {
		cur.Next = cur.Next.Next
	}
	return
}
