package binary_trees

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
二叉树打印顺序：
1、先序：头、左、右
2、中序：左、头、右
3、后序：左、右、头
*/

/*
非递归遍历：
先序：
	1、从栈中弹出一个节点cur
	2、打印处理cur
	3、先右再左（如果有）
Stack<Node>stack 二 new Stack<Node>();
	stack.add(head);
while (Istack. isEmpty()） {
	head = stack.pop();
	System. out.print (head .value +""）;
	if (head. right != nu11) R
	stack.push (head. right)
}
stack• push(head.right);
if (head. left != null) { stack.push (head. left);


后序：
准备两个栈：stack1和stack2
以头左右的顺序写入stack1，逻辑为：
1、弹cur
2、cur放入stack2
3、先左再右
public static void posOrderUnRecur1(Node head)
System.out.print("pos-order: ");
if (head != nul1){
Stack<Node> s1 = new Stack<Node>();
Stack<Node> new Stack<Node>();
s1.push(head);
while (!s1. isEmpty()){
head = s1.pop();
s2.push (head);
if (head. left l= nul1) { s1.push(head. left);
if (head. right != nul1) {
s1.push(head.right);
}


System. out.print(s2.pop().value +"");
while (!s2. isEmpty()){
System.out. printin();

中序：
每棵子树左边界(顺着左节点查找)进栈，依次弹出的过程中对弹出的右节点重复
*/
