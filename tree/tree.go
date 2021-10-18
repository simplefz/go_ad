package tree

/***
	ADT 树
    DATA
        树是有一个根节点和若干棵子树构成。树中节点具有相同数据类型及层次关系。
    Operation
        InitTree(t) 构造空树
        DestroyTree(t) 销毁数
        CreateTree(t,definition) 按definition 中给出树的定义来构造树
        ClearTree(t) 清空树
		TreeEmpty(t) 判断树是否为空
		TreeDepth(t) 返回树的深度
 		Root(t) 返回t的根节点
		Value(t,curE) curE 是t树中的一个节点，返回此节点的值
		Assign(t,curE,value) 给t树中的curE 赋值为value
        Parent(t,curE) 若curE是树t非根节点返回他的双亲，否则返回空
        LeftChild(t,curE) 若curE是树t非叶节点，则返回他的最左孩子，否则返回空
	endADT
*/
