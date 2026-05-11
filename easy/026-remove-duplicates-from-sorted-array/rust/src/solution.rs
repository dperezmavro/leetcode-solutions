pub struct Solution {}

impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let mut k: usize = 0;
        // TODO: this clone looks wrong
        for (i, val) in nums.clone().iter().enumerate() {
            if i == 0 {
                k += 1;
                continue;
            }

            if *val == nums[k - 1] {
                nums[i] = -1;
            } else {
                nums[k] = *val;
                k += 1;
                nums[k] = -1;
            }
        }

        k.try_into().unwrap()
    }
}
