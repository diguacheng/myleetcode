class Solution:
    def maxProfit(self, prices):
        '''
        pricrs list[int]
        return int
        '''
        max = 0
        for i in range(len(prices)):
            for j in range(i+1, len(prices)):
                if prices[j]-prices[i] > max:
                    max = prices[j]-prices[i]
        return max

    def maxProfit1(self, prices):
        '''
        pricrs list[int]
        return int
        '''
        mmax=0
        minport=1<<31
        for i in range(len(prices)):
            if prices[i]<minport:
                minport=prices[i]
            if prices[i]-minport>mmax:
                mmax=prices[i]-minport
        return mmax


prices=[7,1,5,3,6,4]    
s=Solution()
print(s.maxProfit(prices))
