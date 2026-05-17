pub struct Solution {}

impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let mut final_array: Vec<i32> =
            Vec::with_capacity(nums1.len() + nums2.len());

        let mut i: usize = 0;
        let mut j: usize = 0;
        while final_array.len() < final_array.capacity() {
            if j == nums2.len() {
                while i < nums1.len() {
                    final_array.push(nums1[i]);
                    i += 1;
                }
                break;
            }
            if i == nums1.len() {
                while j < nums2.len() {
                    final_array.push(nums2[j]);
                    j += 1;
                }
                break;
            }

            if nums1[i] < nums2[j] {
                final_array.push(nums1[i]);
                i += 1;
            } else {
                final_array.push(nums2[j]);
                j += 1;
            }
        }

        let answer: f64 = match final_array.len() % 2 {
            0 => {
                ((final_array[final_array.len() / 2] as f64)
                    + (final_array[(final_array.len() / 2) - 1] as f64))
                    / 2.0
            }
            _ => final_array[final_array.len() / 2] as f64,
        };
        answer
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn odd_size() {
        let result = Solution::find_median_sorted_arrays(vec![1, 3], vec![2]);
        assert_eq!(result, 2.0);
    }

    #[test]
    fn even_size() {
        let result =
            Solution::find_median_sorted_arrays(vec![1, 2], vec![3, 4]);
        assert_eq!(result, 2.5);
    }

    #[test]
    fn negative_values() {
        let result =
            Solution::find_median_sorted_arrays(vec![-10, -9, -8], vec![1, 2]);
        assert_eq!(result, -8.0);
    }
}
