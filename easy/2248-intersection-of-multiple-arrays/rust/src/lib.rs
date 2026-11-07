use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn intersection(nums: Vec<Vec<i32>>) -> Vec<i32> {
        let target_count = nums.len();
        let mut counter: HashMap<i32, usize> = HashMap::new();
        for vec in nums {
            for n in vec {
                *counter.entry(n).or_insert(0) += 1;
            }
        }

        let mut res: Vec<i32> = counter
            .iter()
            .filter(|e| *e.1 == target_count)
            .map(|e| *e.0)
            .collect::<Vec<i32>>();

        res.sort();
        res
    }
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_empty() {
        let s = Solution::intersection(vec![vec![0], vec![0], vec![0]]);
        let want = vec![0];
        assert_eq!(want, s);
    }

    #[test]
    fn test_1() {
        let s = Solution::intersection(vec![
            vec![3, 1, 2, 4, 5],
            vec![1, 2, 3, 4],
            vec![3, 4, 5, 6],
        ]);
        let want = vec![3, 4];
        assert_eq!(want, s);
    }

    #[test]
    fn test_2() {
        let s = Solution::intersection(vec![vec![1, 2, 3], vec![4, 5, 6]]);
        let want: Vec<i32> = vec![];
        assert_eq!(want, s);
    }
}
