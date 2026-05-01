use find_the_highest_altitude::solution;

struct TestCase {
    result: i32,
    input: Vec<i32>,
}

#[cfg(test)]
#[test]
fn test_largest_altitude() {
    let tests: Vec<TestCase> = vec![
        TestCase {
            result: 1,
            input: vec![-5, 1, 5, 0, -7],
        },
        TestCase {
            result: 0,
            input: vec![-4, -3, -2, -1, 4, 3, 2],
        },
    ];
    for t in tests {
        let got = solution::Solution::largest_altitude(t.input);
        assert_eq!(t.result, got);
    }
}
