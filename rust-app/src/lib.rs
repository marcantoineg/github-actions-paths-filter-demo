use std::fs;

#[allow(dead_code)]
const FILE_PATH: &str = "../super-important-data.txt";

#[allow(dead_code)]
pub fn app() -> String {
    let data = fs::read_to_string(FILE_PATH).unwrap();
    return data + ", BUT IN RUUUUUUUST BABYYYYYYYYY";
}
