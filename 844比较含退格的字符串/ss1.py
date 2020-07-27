def backspaceCompare1(S, T):
    def func(S):
        s = []
        for i in S:
            if i == '#':
                if len(s) > 0:
                    s.pop()
                else:
                    pass
            else:
                s.append(i)
        return s
    return func(S)==func(T)

def backspaceCompare2(S, T):
    s = []
    t = []
    for i in S:
        if i == '#':
            if len(s) > 0:
                s.pop()
            else:
                pass
        else:
            s.append(i)
    for i in T:
        if i == '#':
            if len(t) > 0:
                t.pop()
            else:
                pass
        else:
            t.append(i)
    return s==t
S = "ab##"
T = "c#d#"
print(backspaceCompare(S,T))
# 时间复杂度O(len(S)+Len(T))
# 空间复杂读OO(len(S)+Len(T))