
def fraction(cont):
    res=[cont[len(cont)-1],1]
    for i in range(len(cont)-2,-1,-1):
        temp=res[1]
        res[1] = res[0]
        res[0]=res[0]*cont[i]+temp
    return res


if __name__=="__main__":
    cont=[i for i in range(1,11)]
    res=fraction(cont)
    print(res)


