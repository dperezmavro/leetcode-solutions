use std::collections::HashMap;
pub struct Solution {}

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut number_index: HashMap<char, i32> = HashMap::new();

        number_index.insert('I', 1);
        number_index.insert('V', 5);
        number_index.insert('X', 10);
        number_index.insert('L', 50);
        number_index.insert('C', 100);
        number_index.insert('D', 500);
        number_index.insert('M', 1000);

        let mut total = 0;
        for (idx, val) in s.chars().enumerate() {
            let curr_val = match number_index.get(&val) {
                Some(v) => *v,
                _ => 0,
            };
            if idx != 0 {
                let prev_char: char = match s.chars().nth(idx - 1) {
                    Some(c) => c,
                    _ => '0',
                };
                let prev_val = match number_index.get(&prev_char) {
                    Some(v) => *v,
                    _ => 0,
                };

                if prev_val < curr_val {
                    // undo the previous value being added
                    total -= prev_val;
                    // compensate for always adding at the end of the loop
                    total -= curr_val;
                    // subtract real value
                    total += curr_val - prev_val;
                }
            }
            total += curr_val;
        }
        total
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_01() {
        let input = String::from("III");
        let output = 3;
        assert_eq!(Solution::roman_to_int(input), output);
    }

    #[test]
    fn test_02() {
        let input = String::from("LVIII");
        let output = 58;
        assert_eq!(Solution::roman_to_int(input), output);
    }

    #[test]
    fn test_03() {
        let input = String::from("XIV");
        let output = 14;
        assert_eq!(Solution::roman_to_int(input), output);
    }

    #[test]
    fn test_04() {
        let input = String::from("MCMXCIV");
        let output = 1994;
        assert_eq!(Solution::roman_to_int(input), output);
    }
}
