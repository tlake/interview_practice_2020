package main

func main() {}

func removeDuplicates(nums []int) int {
    anchor := 0
    scout := 1
    
    if len(nums) == 0 {
        return 0
    }
    
    if nums[0] == nums[len(nums)-1] {
        return 1
    }
    
    for scout < len(nums) {
        if nums[scout] != nums[anchor] {
            anchor++
            nums[anchor] = nums[scout]
        }
        
        scout++
    }
    
    return anchor+1
}