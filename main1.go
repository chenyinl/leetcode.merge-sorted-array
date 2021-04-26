func merge(nums1 []int, m int, nums2 []int, n int)  {
    j:=0;
    i:=0;
    c:=0; // count for nums1
    for i=0; i<m+n && j<n && c<m;i++ {
        if(nums1[i]>=nums2[j]){
            copy(nums1[i+1:], nums1[i:])
            nums1[i] = nums2[j];
            j++
        }else{
            c++
        }
        
    }
    fmt.Println(i, j);
    copy(nums1[i:], nums2[j:])
    
}
