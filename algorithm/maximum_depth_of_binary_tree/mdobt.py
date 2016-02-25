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
        return travel(root)


def travel(node):
    if node is None:
        return 0
    l_depth = 1
    r_depth = 1
    if node.left:
        l_depth += travel(node.left)
    if node.right:
        r_depth += travel(node.right)
    return max(l_depth, r_depth)
