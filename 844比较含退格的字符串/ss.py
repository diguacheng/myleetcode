def backspaceCompare(S:str, T:str):
    def func(S):
        temp=0
        for i in range(len(S)-1,-1,-1):
            if S[i]!="#" and temp>0:
                S=S[:i]+'#'+S[i+1:]
                temp-=1
            elif S[i]=='#':
                temp+=1
            else:
                pass
        S=S.replace('#','')
        return S
    return func(S)==func(T)
S = "xywrrmp"
T = "xywrrmu#p"

print(backspaceCompare(S,T))