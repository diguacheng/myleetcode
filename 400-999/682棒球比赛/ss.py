def calPoints(ops):
    li=[]
    for item in ops:
        if item=="+":
            li.append(li[-1]+li[-2])
        elif  item=="C":
            li.pop(-1)
        elif item=="D":
            li.append(li[-1]*2)
        else:
            li.append(int(item))
    return sum(li)
        
ops=["5","2","C","D","+"]
print(calPoints(ops))