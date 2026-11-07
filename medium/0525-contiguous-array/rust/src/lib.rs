use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn find_max_length(nums: Vec<i32>) -> i32 {
        if nums.len() == 1 {
            return 0;
        }

        let mut max_len: i32 = 0;
        let mut diff: i32 = 0;
        let mut diffs: HashMap<i32, i32> = HashMap::with_capacity(nums.len());
        diffs.insert(0, -1);

        for (i, &n) in nums.iter().enumerate() {
            diff += if n == 1 { 1 } else { -1 };

            if let Some(&prev_idx) = diffs.get(&diff) {
                max_len = max_len.max(i as i32 - prev_idx);
            } else {
                diffs.insert(diff, i as i32);
            }
        }

        max_len
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn case_1() {
        let res = Solution::find_max_length(vec![0, 1]);
        assert_eq!(2, res);
    }

    #[test]
    fn case_2() {
        let res = Solution::find_max_length(vec![0, 1, 0]);
        assert_eq!(2, res);
    }

    #[test]
    fn case_3() {
        let res = Solution::find_max_length(vec![0, 1, 1, 1, 1, 1, 0, 0, 0]);
        assert_eq!(6, res);
    }

    #[test]
    fn case_4() {
        let res = Solution::find_max_length(vec![0, 1, 0, 1, 0, 1]);
        assert_eq!(6, res);
    }

    #[test]
    fn case_5() {
        let res = Solution::find_max_length(vec![1]);
        assert_eq!(0, res);
    }

    #[test]
    fn case_6() {
        let res = Solution::find_max_length(vec![0, 0, 1, 0, 0]);
        assert_eq!(2, res);
    }

    #[test]
    fn case_7() {
        let res = Solution::find_max_length(vec![0, 0, 1, 0, 1]);
        assert_eq!(4, res);
    }
}
