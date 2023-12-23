+++
title = 'Data Structure and Algorithms in Java'
description = 'My notes on DS in java with included practice problems'
date = 2023-11-15T15:49:08-05:00
draft = false
ShowToc = true
+++

## Intro to DS & Algo with Java

> "The difference between a bad programmer and a good one is whether he considers his code or his data structures more important. Bad programmers worry about the code. Good programmers worry about data structures and their relationships." {{< line_break >}} **Linus Torvalds, 2006**

<!-- DS cover image -->

{{< figure align=center alt="Kazi Hossain picture of a boy working on the computer with wires everywhere" src="/ds/main.jpg" >}}

<!-- DS cover image -->


A data structure is a collection of data items stored in memory in an organized manner, which makes data storing, editing and retrieving faster with consuming less space. Where an algorithm is a set of detailed logical step by step instructions used to solve a given problem. 



## Lists 


## Stacks 


## Queues


## Recursion


## Merge Sort




## Trees
A tree is a non-linear data structure where nodes are organized in a hierarchy. A real life example of thinking about the Trees data structure is like a family tree. Where there is a root from whom there are children of their childrens etc. 

{{< figure align=center alt="Drawining representation of a numerical tree" src="/ds/tree_example.png" width="40%" caption="Example of a tree" >}}

**Root** : The top node of the tree is known as the root node or parent node. {{< line_break >}}
**Leaf** : Any nodes at the bottom of the tree is known as the leaf node. They don't have any children. {{< line_break >}}
**Branch** : Any nodes that is a child node and have children known as a Brench. They have both incoming and outgoing edges. {{< line_break >}}
**Subtree** : If you take any node and cut of its edges or brances that connects to its previous nodes it becomes a subtree. There are as many subtree in a tree as there are nodes.

{{< figure align=center alt="Drawining representation of subtrees" src="/ds/sub-tree.png" width="40%" caption="Example of subtrees in a tree" >}}

**Size :** The size of a tree is the number of nodes in that tree including the root. In the above example the size of the tree is 10.

**Depth :** The levels of branches starting from the root which is at level zero. The figure below in blue showcases the depth in a tree.

{{< figure align=center alt="Drawining representation depth in a tree" src="/ds/depth_tree.png" width="40%" caption="Example of depths in a tree" >}}

**Height :** The number of edges above from the furthest leaf node. (Bottom to top)

{{< figure align=center alt="Drawining representation of heights in a tree" src="/ds/height_tree.png" width="50%" caption="Example of height in a tree" >}}



## Binary Search Trees




## Balanced Binary Search Trees



## Priority Queues


## Heaps & Heapsort



## Hash Tables


## Graphs 

Example code
```java
public class hello {
    public static void main (String [] args){
        System.out.println ("Hello");
    }
}

```