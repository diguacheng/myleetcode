class MinStack:
    
    def __init__(self):
        """
        不同步栈
        """
        self.data=[]# 数据栈
        self.list=[]# 辅助栈 
        

    def push(self, x: int):
        self.data.append(x)
        if  len(self.list)==0 or x<=self.list[-1]:
            self.list.append(x)
        else:
            pass
        

    def pop(self):
        if self.data :
            if self.list[-1]==self.data[-1]:
                self.list.pop()
            return self.data.pop()

        

    def top(self):
        if self.data:
            return self.data[-1]
        
        

    def getMin(self):
        if self.list:
            return self.list[-1]

    

obj=MinStack()
obj.push(-2)
obj.push(0)
obj.push(-3)
print(obj.getMin())
print(obj.pop())
print(obj.top())
print(obj.getMin())
        