use std::cmp;

pub struct Solution {}

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        if height.len() == 2 {
            return cmp::min(height[0], height[1]);
        }

        let mut left: usize = 0;
        let mut right: usize = height.len() - 1;
        let mut max_volume = cmp::min(height[left], height[right]) * ((right - left) as i32);
        while left < right {
            if height[left] > height[right] {
                right -= 1;
            } else if height[left] <= height[right] {
                left += 1;
            }
            let tmp_volume = cmp::min(height[left], height[right]) * ((right - left) as i32);
            if tmp_volume > max_volume {
                max_volume = tmp_volume;
            }
        }
        max_volume
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

    #[test]
    fn test_4() {
        let result = Solution::max_area(vec![8, 7, 2, 1]);
        assert_eq!(result, 7);
    }
}
