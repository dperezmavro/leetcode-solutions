pub struct Solution {}

impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let mut k: usize = 1;
        for i in 1..nums.len() {
            if nums[i] > nums[k - 1] {
                nums[k] = nums[i];
                k += 1;
            }
        }

        k as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_basic() {
        let mut nums = vec![1, 1, 2];
        assert_eq!(Solution::remove_duplicates(&mut nums), 2);
        assert_eq!(&nums[..2], &[1, 2]);
    }

    #[test]
    fn test_longer() {
        let mut nums = vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
        assert_eq!(Solution::remove_duplicates(&mut nums), 5);
        assert_eq!(&nums[..5], &[0, 1, 2, 3, 4]);
    }

    #[test]
    fn test_no_duplicates() {
        let mut nums = vec![3, 4];
        assert_eq!(Solution::remove_duplicates(&mut nums), 2);
        assert_eq!(&nums[..2], &[3, 4]);
    }

    #[test]
    fn test_all_same() {
        let mut nums = vec![1, 1, 1, 1];
        assert_eq!(Solution::remove_duplicates(&mut nums), 1);
        assert_eq!(&nums[..1], &[1]);
    }

    #[test]
    fn test_single_element() {
        let mut nums = vec![42];
        assert_eq!(Solution::remove_duplicates(&mut nums), 1);
    }
}
