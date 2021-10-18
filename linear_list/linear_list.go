package linear_list

/****
ADT 线性表
Data
   线性表的数据对象集合为{a1,a2...an} 每个元素类型均为DataType. 其中出第一个元素a1外，每一个元素有且只有一个直接前驱
元素，除最后一个元素an外. 每一个元素有且只有一个直接后记元素.数据元素之间的关系是一对一的关系.
Operation
   InitList(*L);  		初始化操作建立一个空的线性表.
   ListEmpty(L);  		若线性表为空，返回true,否则返回false.
   ClearList(*L); 		将线性表清空.
   GetElem(L,i)(e);		讲线性表L中额第i个位置元素值返回给e.   时间复杂度 O(1)
   LocateElem(L,e);     在线性表L中查找给定值e相等的元素,如果查找成功返回该元素在表中序号标识成功，否则返回 0 标识失败. 时间复杂度 O(n)
   ListInsert(*L,i,e);  在线表L中的第i个位置插入新元素e 时间复杂度 O(n)
   ListDelete(*L,i,*e); 删除线性表L中第i个位置并用e返回其值 时间复杂度 O(n)
   ListLength(L);		返回线性表L的元素个数
endADT
****/

const (
	MaxSqlListArrCap = 2 << 15
)

type ElementType int

type SqList struct {
	Data   [MaxSqlListArrCap]ElementType // 数组存储值
	Length uint16                        // 线性表长度
}

// 初始化线性表
func InitList(length uint16) (m *SqList) {
	m = new(SqList)
	m.Length = length
	return
}

// 判断线性表是否为空
func (p *SqList) ListEmpty() (b bool) {
	if p.Length == 0 {
		b = true
		return
	}
	return
}

// 清空线性表 将所有值回复默认值
func (p *SqList) ClearList() (b bool) {
	if p.Length == 0 {
		return
	}
	b = true
	for i := 0; i < int(p.Length); i++ {
		p.Data[i] = 0
	}
	return
}

// 获取某个位置元素值
func (p *SqList) GetElem(i uint16) (b bool, v ElementType) {
	if i > p.Length {
		return
	}
	v = p.Data[i-1]
	b = true
	return
}

// 插入值
func (p *SqList) ListInsert(i uint16, v ElementType) (b bool, data []ElementType) {
	if i > p.Length || i == 0 || p.Data[p.Length-1] != 0 {
		return
	}
	for k := p.Length - 1; k > i; k-- {
		p.Data[k+1] = p.Data[k]
	}
	b = true
	p.Data[i-1] = v
	data = p.Data[:p.Length]
	return
}

// 删除值
func (p *SqList) ListDelete(i uint16) (b bool, e ElementType, data []ElementType) {
	if i > p.Length || i == 0 {
		return
	}
	e = p.Data[i-1]
	for k := i - 1; k < p.Length; k++ {
		p.Data[k] = p.Data[k+1]
	}
	b = true
	data = p.Data[:p.Length]
	return
}

func (p *SqList) ListLength() (b bool, l uint16) {
	if p.Length == 0 {
		return
	}
	var i uint16
	var j uint16
	j = 1
	for i = p.Length - 1; i > 0; i-- {
		if p.Data[i] != 0 {
			j += 1
		}
	}
	b = true
	l = j
	return
}
