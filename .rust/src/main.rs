#[macro_use]
extern crate clap;
use clap::App;
use std::path::Path;

fn main() {
    let yaml = load_yaml!("cli.yaml");
    let matches = App::from_yaml(yaml).get_matches();

    let input = matches.value_of("INPUT").unwrap();

    println!("{}", Path::new(input).exists());
}
