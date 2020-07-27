
class TreeNode(object):
     def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None


class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        def rserialize(root, string):
            if root is None:
               string += 'none,'
            else:
                string += str(root.val) + ','
                string=rserialize(root.left, string)
                string = rserialize(root.right, string)
            return string
        
        return rserialize(root,'')
            

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        def rdeserialize(l: list):
            if l[0] == 'none':
                l.pop(0)
                return None
            root = TreeNode(int(l[0]))
            l.pop(0)

            root.left = rdeserialize(l)
            root.right = rdeserialize(l)
            return root
        
        data_list = data.split(',')
        data_list=data_list[:len(data_list)-1]
        root = rdeserialize(data_list)
        return root



codec = Codec()
str1 = '1,2,none,none,3,4,none,none,5,none,none,'
root=codec.deserialize(str1)
s = codec.serialize(root)
print(s)
