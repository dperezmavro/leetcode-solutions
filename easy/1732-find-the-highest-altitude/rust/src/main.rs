mod solution;

fn main() {
    println!(
        "{}",
        solution::Solution::largest_altitude(vec![-5, 1, 5, 0, -7])
    );
    println!(
        "{}",
        solution::Solution::largest_altitude(vec![-4, -3, -2, -1, 4, 3, 2])
    );
}
