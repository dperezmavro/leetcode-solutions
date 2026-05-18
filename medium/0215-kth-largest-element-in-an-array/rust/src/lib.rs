use std::cmp::Reverse;
use std::collections::BinaryHeap;

pub struct Solution {}

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut min_heap = BinaryHeap::new();
        for n in nums {
            min_heap.push(Reverse(n));
            if min_heap.len() > k as usize {
                min_heap.pop();
            }
        }
        min_heap.peek().unwrap().0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_1() {
        let result = Solution::find_kth_largest(vec![3, 2, 1, 5, 6, 4], 2);
        assert_eq!(result, 5);
    }

    #[test]
    fn test_2() {
        let result = Solution::find_kth_largest(vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4);
        assert_eq!(result, 4);
    }
}
