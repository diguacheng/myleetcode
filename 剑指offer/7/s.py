
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def buildTree(self, preorder, inorder):
        """
        递归法  
        """
        if len(preorder) == 0:
            return None
        root = TreeNode(preorder[0])
        index=inorder.index(preorder[0])
        root.left=self.buildTree(preorder[1:index+1], inorder[0:index])
        root.right=self.buildTree(preorder[index+1:], inorder[index+1:])
        return root
    def buildTree1(self, preorder, inorder):
        if len(preorder) == 0:
            return None
        root=TreeNode(preorder[0])
        l=[]
        l.append(root)
        index=0
        for i in range(1,len(preorder)):
            node=l[-1]
            if node.val!=inorder[index]:
                node.left=TreeNode(preorder[i])
                l.append(node.left)
            else:
                while len(l)!=0 and l[-1].val==inorder[index]:
                    node=l.pop()
                    index+=1
                node.right=TreeNode(preorder[i])
                l.append(node.right)
        
        return root

preorder=[3,9,8,5,4,10,20,15,7]
inorder=[4,5,8,10,9,3,15,20,7]
s=Solution()
print(s.buildTree1(preorder,inorder))