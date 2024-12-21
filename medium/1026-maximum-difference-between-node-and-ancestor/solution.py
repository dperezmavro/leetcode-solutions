from typing import Optional, List

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def maxAncestorDiff(self, root: Optional[TreeNode]) -> int:
        if self.isLeaf(root):
            return 0
        return self.maxAncestorDiffInternal(root, [root.val])
    
    def isLeaf(self, root: TreeNode) -> bool:
        return root and not (root.left or root.right)
    
    def maxAncestorDiffInternal(self, root: Optional[TreeNode], ancestors:List[int]) -> int:
        self_max = max([abs(root.val - x) for x in ancestors])
        if self.isLeaf(root):
            return self_max
        
        diff_l = -1
        diff_r = -1

        if root.left:
            ancestors_new = ancestors.copy()
            ancestors_new.append(root.val)
            child_max = self.maxAncestorDiffInternal(root.left, ancestors_new)
            diff_l = max(self_max, child_max, abs(root.val-root.left.val))
        
        if root.right:
            ancestors_new = ancestors.copy()
            ancestors_new.append(root.val)
            child_max = self.maxAncestorDiffInternal(root.right, ancestors_new)
            diff_r = max(self_max, child_max, abs(root.val-root.right.val))

        return max(diff_l, diff_r)