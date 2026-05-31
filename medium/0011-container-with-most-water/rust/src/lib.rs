use std::cmp;

pub struct Solution {}

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        if height.len() == 2 {
            return cmp::min(height[0], height[1]);
        }
        0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_1() {
        let result = Solution::max_area(vec![1, 2]);
        assert_eq!(result, 1);
    }

    #[test]
    fn test_2() {
        let result = Solution::max_area(vec![1, 1]);
        assert_eq!(result, 1);
    }

    #[test]
    fn test_3() {
        let result = Solution::max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]);
        assert_eq!(result, 49);
    }
}
