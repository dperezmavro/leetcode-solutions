use std::collections::VecDeque;

pub struct Solution {}

impl Solution {
    pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
        if nums.len() == 1 {
            return nums;
        }

        let mut result: Vec<i32> = vec![];
        let mut deq: VecDeque<i32> = VecDeque::new();
        let mut curr_window: VecDeque<i32> = VecDeque::new();

        // build the first window
        for n in &nums[0..k as usize] {
            curr_window.push_back(*n);
            if deq.len() == 0 {
                deq.push_front(*n);
                continue;
            }

            while deq.len() > 0 {
                // written here for legibility
                let b = deq.back().unwrap();
                if *b < *n {
                    deq.pop_back();
                }
            }

            deq.push_back(*n);
        }

        result.push(*(deq.front().unwrap()));

        // for n in k..nums.len() as i32 {
        for n in &nums[(k as usize)..] {
            let oldest_num = curr_window.pop_front().unwrap();
            if *(deq.front().unwrap()) == oldest_num {
                deq.pop_front();
            }

            let new_num = *n;
            curr_window.push_back(new_num);
            println!("1 n {} new_num {} curr_w  {:?}", *n, new_num, curr_window);
            while deq.len() > 0 {
                // written here for legibility
                println!("2 n {} new_num {} curr_w  {:?}", *n, new_num, curr_window);

                let b = deq.back().unwrap();
                if *b < new_num {
                    deq.pop_back();
                }
            }

            deq.push_back(new_num);
            result.push(*(deq.front().unwrap()));
        }

        result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_1() {
        let result = Solution::max_sliding_window(vec![1], 1);
        assert_eq!(result, vec![1]);
    }

    #[test]
    fn test_2() {
        let result = Solution::max_sliding_window(vec![1, -1], 1);
        assert_eq!(result, vec![1, -1]);
    }

    #[test]
    fn test_3() {
        let result = Solution::max_sliding_window(vec![7, 2, 4], 2);
        assert_eq!(result, vec![7, 4]);
    }

    // #[test]
    // fn test_4() {
    //     let result = Solution::max_sliding_window(vec![1, 3, -1, -3, 5, 3, 6, 7], 3);
    //     assert_eq!(result, vec![3, 3, 5, 5, 6, 7]);
    // }

    // #[test]
    // fn test_5() {
    //     let result = Solution::max_sliding_window(vec![9, 10, 9, -7, -4, -8, 2, -6], 5);
    //     assert_eq!(result, vec![10, 10, 9, 2]);
    // }

    // #[test]
    // fn test_6() {
    //     let result = Solution::max_sliding_window(vec![9, 8, 9, 8], 3);
    //     assert_eq!(result, vec![9, 9]);
    // }
}
