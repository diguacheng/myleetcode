package main



func maxSatisfied(customers []int, grumpy []int, X int) int {
    n:=len(customers)
    total:=0 
    for i:=0;i<n;i++{
        if grumpy[i]==0{
            total+=customers[i]
        }
    }
    left:=0 
    right:=X-1
    maxT:=total
    for i:=left;i<=right;i++{
        if grumpy[i]==1{
            maxT+=customers[i]
        }
    }
    right++
    res:=maxT
    for right<n{
        maxT+=grumpy[right]*customers[right]-grumpy[right-X]*customers[right-X]
        res=max(res,maxT)
        right++

    }
    return res
}
func max(a,b int)int{
    if a<b{
        return b
    }
    return a
}