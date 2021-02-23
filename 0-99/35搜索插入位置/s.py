def searchInsert( nums,target):
    if target in nums:
        return nums.index(target)
    else:
        if nums[-1]<target:
            return len(nums)
        else:
            for i in range(len(nums)):
                if  target<nums[i]:
                    return i
            


