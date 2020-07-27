class MyStack:
    
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.list1=[]
        self.list2=[]
        self.top=0

        

    def push(self, x: int):
        """
        Push element x onto stack.
        """
        self.list1.append(x)
        self.top=x
        

    def pop(self) -> int:
        """
        Removes the element on top of the stack and returns that element.
        """
        

    def top(self) -> int:
        """
        Get the top element.
        """
        

    def empty(self) -> bool:
        """
        Returns whether the stack is empty.
        """
        
