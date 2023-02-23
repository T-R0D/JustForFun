#[allow(dead_code)]
fn char_to_ascii_value(c: char) -> Result<u8, String> {
    let value = c as u32;
    if value > u8::MAX as u32 {
        return Err(String::from(format!("c ({value}) is outside ASCII range")));
    }

    Ok(value as u8)
}

// To ascii index (lower, upper, combined)
// 
