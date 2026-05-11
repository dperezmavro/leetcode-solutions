mod solution;

fn main() {
    let mut v: Vec<i32> = vec![1, 1, 2];
    let mut res = solution::Solution::remove_duplicates(&mut v);
    println!("{:?} {}", v, res);

    v = vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
    res = solution::Solution::remove_duplicates(&mut v);
    println!("{:?} {}", v, res);

    v = vec![3, 4];
    res = solution::Solution::remove_duplicates(&mut v);
    println!("{:?} {}", v, res);
}
