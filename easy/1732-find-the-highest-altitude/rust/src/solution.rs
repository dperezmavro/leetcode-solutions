pub struct Solution {}

impl Solution {
    pub fn largest_altitude(gain: Vec<i32>) -> i32 {
        let mut curr_altitude: i32 = 0;
        let mut max_altitude: i32 = 0;
        for g in gain {
            curr_altitude += g;
            if curr_altitude > max_altitude {
                max_altitude = curr_altitude;
            }
        }
        return max_altitude;
    }
}
