def findRepeatNumber(nums):
    """
    哈希表  时间复杂度O(n) 空间复杂度O(n)
    """
    d={}
    for i in nums:
        if i in d.keys():
            return i
        else:
            d[i]=1


def findRepeatNumber1(nums):
    """
    先将数组排序,然后扫描数组
    排序的时间效率为O(nlogn)
    """
    nums.sort()
    temp=nums[0]
    for i in range(1,len(nums)):
        if nums[i]==temp:
            return temp
        else:
            temp=nums[i]



nums=[2, 3, 1, 0, 2, 5, 3]
print(findRepeatNumber1(nums))