/* use genertic parameter to write shorter code */

// Find the largest element of a list
use std::cmp::PartialOrd;
pub fn largest<T: PartialOrd + Copy>(list: &[T]) -> T {
    // The type which impls PartialOrd and Copy trait can use
    // the function
    let mut largest = list[0];
    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }
    largest
}

// Use a lifttime annotions to get longest str
pub fn str_longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

//function woth a lifttime ignore
pub fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i]
        }
    }

    &s[..]
}
