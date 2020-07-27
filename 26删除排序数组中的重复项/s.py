def removeDuplicates(nums):
    """
    考虑到数组是有序的，每两个为一组比较，若相同则删除一个
    """
    i=0
    while i<len(nums)-1:
        if nums[i+1]==nums[i]:
            nums.pop(i+1)
        else:
            i+=1
    return len(nums)

def removeDuplicates1(nums):
    """
    双指针
    i为慢指针 j为快指针 
    j往后运行，遇到num[j]!=nums[i],则将i增加1，且此时的nums[i]置为num是[j]的值，否则，j加1
    """
    if len(nums)==0:
        return 0
    i=0
    for j in range(1,len(nums)):
        if nums[j]!=nums[i]:
            i+=1
            nums[i]=nums[j]
    return i+1


    
nums=[0,0,1,1,1,2,2,3,3,4]
print(removeDuplicates1(nums))
print(nums)

