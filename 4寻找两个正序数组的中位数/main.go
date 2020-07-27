package mian

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ll:=len(nums1)+len(nums2)
	if ll%2==1{
		midindex:=ll/2
		return float64(getKele(nums1,nums2,midindex+1))
	}else{
		micdindex1,midindex2:=ll/2-1,ll/2
		return float64(getKele(nums1,nums2,midindex1+1)+getKele(nums1,nums2,midindex2+1))/2.0
	}
}


func getKele(nums1 []int, nums2 []int,k int)int{
	idx1,idx2 :=0,0
	for {
		if idx1==len(nums1){
			return nums2[index2+k-1]
		}
		if idx2==len(nums2){
			return nums1[index1+k-1]
		}
		if k==1{
			return min(nums1[idx1],nums2[idx2])
		}
		half:=k/2
		newidx1:=min(idx1+half,len(nums1)-1)-1
		newidx2:=min(idx2+half,len(nums2)-1)-1
		pivot1,pivot2:=nums1[newidx1],nums2[newidx2]
		if pivot1<=pivot2{
			k=k-(newidx1-idx1+1)
			idx1=newidx1+1
		} else{
			k=k-(newidx2-idx2+1)
			idx2=newidx2+1
		}
	
	}
	return 0
}

func min(a,b int) int{
	if a<b {
		return a
	}
	return b
}
func mian(){


}