class Solution:
    def combine(self, n: int, k: int):
        # init first combination
        nums = list(range(1, k + 1)) + [n + 1]
        output, j = [], 0
        while j < k:
            # add current combination
            output.append(nums[:k])
            print(nums[:k])
            # increase first nums[j] by one
            # if nums[j] + 1 != nums[j + 1]
            j = 0
            while j < k and nums[j + 1] == nums[j] + 1:
                nums[j] = j + 1
                j += 1
            nums[j] += 1
            
        return output

x=Solution()
x.combine(4,2)