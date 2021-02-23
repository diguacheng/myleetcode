
def intersection(nums1, nums2):
    return list(set(nums1)&set(nums2))


nums1 = [1,2,2,1]
nums2 = [2,2]
print(intersection(nums1,nums2))