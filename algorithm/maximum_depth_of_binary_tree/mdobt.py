# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def maxDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        def travel(node):
            if node is None:
                return 0
            return max(1+travel(node.left), 1+travel(node.right))

        return travel(root)
