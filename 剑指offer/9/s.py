class CQueue:
    
    def __init__(self):
        self.input=[]
        self.output=[]


    def appendTail(self, value):
        self.input.append(value)


    def deleteHead(self):
        if len(self.input)==0 and len(self.output)==0:
            return -1
        elif len(self.output)!=0:
            return self.output.pop()
        else:
            while len(self.input)!=0:
                self.output.append(self.input.pop())
            return self.output.pop()
