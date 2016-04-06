#!/usr/bin/env python
# -*- coding: utf-8 -*-


# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def isSameTree(self, p, q):
        """
        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        pi = self.inorder(p)
        pp = self.preorder(p)
        qi = self.inorder(q)
        qp = self.preorder(q)
        if pi == qi and pp == qp:
            return True
        return False

    def inorder(self, root):
        if root is None:
            return '#'
        return self.inorder(root.left)+str(root.val)+self.inorder(root.right)

    def preorder(self, root):
        if root is None:
            return '#'
        return str(root.val)+self.preorder(root.left)+self.preorder(root.right)
