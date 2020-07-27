
def twoSum(nums, target):
    d={}
    n = len(nums)
    for i in range(n):
        a=target-nums[i]
        if a in d.keys() and d[a]!=i:
            res=[d[a],i]
            break
        d[nums[i]] =i
    return res
        




nums = [3,3]
target = 6
print(twoSum(nums, target))
