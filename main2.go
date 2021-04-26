func merge(nums1 []int, m int, nums2 []int, n int)  {
  // merge two array into the end of array1
    for i:=m+n-1; i>=0 && n-1>=0; i-- {
        if(m-1>=0 && nums1[m-1]>=nums2[n-1]){
            nums1[i] = nums1[m-1];
            //fmt.Println(">", nums1[m-1])
            m--
        }else{
            nums1[i] = nums2[n-1];
            //fmt.Println(">", nums2[n-1])
            n--
        }
    }
}
