use std::collections::{HashMap, VecDeque};

pub struct Solution {}

impl Solution {
    pub fn find_max_length(nums: Vec<i32>) -> i32 {
        if nums.len() == 1 {
            return 0;
        }

        let mut left = 0;
        let mut right = 0;
        let mut max_len: i32 = 0;
        let mut current_window: VecDeque<i32> = VecDeque::new();
        let mut current_counts: HashMap<i32, usize> = HashMap::with_capacity(2);
        current_counts.insert(0, 0);
        current_counts.insert(1, 0);

        while right < nums.len() || left < right {
            if right < nums.len() {
                let n = nums[right];

                current_window.push_back(n);
                current_counts.entry(n).and_modify(|counter| *counter += 1);

                let zero_count =
                    current_counts.get(&0).expect("0 should be here");
                let one_count =
                    current_counts.get(&1).expect("1 should be here");

                if zero_count == one_count {
                    max_len =
                        Solution::max(max_len, current_window.len() as i32);
                }
                right += 1;
            } else if left < right {
                while left < right {
                    let n =
                        current_window.pop_front().expect("pop_front errored");
                    current_counts.entry(n).and_modify(|counter| *counter -= 1);
                    let zero_count =
                        current_counts.get(&0).expect("0 should be here");
                    let one_count =
                        current_counts.get(&1).expect("1 should be here");

                    if zero_count == one_count {
                        max_len =
                            Solution::max(max_len, current_window.len() as i32);
                    }
                    left += 1;
                }
            }
        }

        max_len
    }

    fn max(a: i32, b: i32) -> i32 {
        if a > b {
            return a;
        }
        b
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
