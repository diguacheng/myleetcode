
def isValid(s):
    l=[]
    for i in s:
        if i in ['[','{','(']:
            l.append(i)
        elif i in [']','}',')'] and len(l) > 0:
            temp=l.pop()
            if i==']' and temp=='[' or i=='}' and temp=='{' or i==')' and temp=='(':
                pass
            else :
                return False
        else :
            return False
    return l==[]
 
        
            
s="}"
print(isValid(s))