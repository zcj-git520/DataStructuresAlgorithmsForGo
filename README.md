# DataStructuresAlgorithmsForGo
## 数据结构的存储方式
* 数据结构的存储的方式只有两种：
1. 顺序存储的数组
2. 链式存储的链表
* 数组和链表是其他数据结构的基础
### 数据和链表的优缺点
* 数组的优缺点：
- 优点：数组是紧凑连续存储，可以随机访问，通过索引快速找到对应的元素，而且相对节约空间，即访问的是时间复杂度为O(1)
- 缺点：因数组是连续存储，当扩容时，需要重新分配一块更大的空间，在将之前的数组进行复制，导致时间复制度为O(n),且当数组进行插入或者进行删除时冒也需要搬移，时间复制度也为O(n)
* 链表的优缺点：
- 优点：因元素不连续，且通过指针进行连接，对于插入和删除时间较快，时间复杂度为O(1)
- 缺点：因元素不连续，故不能连续索引访问，对于查询时间复杂度为O(1)-O(n)
## 数据结构的基本操作
* 数据结构的基本操作为：对数据的增删改查，且尽可能的高效的增删改查
* 查：各种数据结构的查询方式：线性查询和非线性查询
- 线性查询是以for/while 迭代为代表方式
- 非线性查询是为递归为代表的方式
## DataStructures
### linked_list
链表的数据结构
#### singly_linked_list
* 单向链表实现路径:Linked_list/Singly_linked_list
#### double_linked_list
* 双向链表实现路径：Linked_list/double_linked_list
#### circle_double_linked_list
* 循环双向链表实现路径：Linked_list/circle_double_linked_list

### stack_queue
* 基于单向链表的栈与基于双向链表的队列
#### stack
* 栈实现路径：\stack_queue\stack
#### queue
* 队列实现路径：\stack_queue\queue

### 树
#### binary_search_tree
* 二叉查找树的路径：\tree\binary_search_tree
* 二叉树- 平衡树的实现: \tree\AVL_tree
* 前缀树- 称字典树:\tree\Trie