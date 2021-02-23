def removeDuplicates(S: str):
    a=[]
    for i in range(len(S)):
        if len(a)==0 or S[i]!=a[-1]:
            a.append(S[i])
        else :
            a.pop(-1)
    return "".join(a)
            


S="abbaca"
print(removeDuplicates(S))


        
