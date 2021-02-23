
def removeElement(nums, val):
    if len(nums)==1:
        if nums[0]==val:
            nums.pop()
            return 0
        else:
            return 1
    i=0
    j=len(nums)-1
    while i<j:
        while nums[i]!=val:
            i+=1
            if i>=len(nums):
                return 0
        while nums[j]==val:
            j-=1
            if j==-1:
                return 0 
        if i<j:
            temp=nums[j]
            nums[j]=nums[i]
            nums[i]=temp
    return i

def removeElement1(nums, val):
    i=0
    for j in range(len(nums)):
        if nums[j]!=val:
            nums[i]=nums[j]
            i+=1
    return i 



def removeElement3(nums, val):
    i=0
    n=len(nums)
    while i<n:
        if nums[i]==val:
            nums[i]=nums[n]
            n-=1
        else:
            i+=1
    return n

nums = [3,2,2,3]
val = 3

print(removeElement(nums,val))
print(nums)
        