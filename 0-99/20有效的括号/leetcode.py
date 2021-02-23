
def isValid(s):
    l=[]
    for i in s:
        if i in ['[','{','(']:
            l.append(i)
        elif i in [']','}',')']:
            if len(l)<=0:
                return False
            else:
                temp=l.pop()
                if i==']' and temp=='[':
                    pass
                elif i=='}' and temp=='{':
                    pass
                elif i==')' and temp=='(':
                    pass
                else :
                    return False
    if len(l)>0:
        return False
    else:
        return True

        
            
s="}"
print(isValid(s))