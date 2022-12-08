use std::{fs::File, io::Write};

use crate::errors::ExecutionError;
use serde_json::{json, to_string};


pub fn output_json_string(prediction: u32, output_directory: &str, output_filename: &str) -> Result<String, ExecutionError> {
    let output = json!({
        "prediction_type": "int",
        "prediction": prediction,
        "explanation_location": format!("{}/{}", output_directory, output_filename)
    });

   match to_string(&output) {
    Ok(output) => {
        Ok(output)
    }
    Err(err) => {
        Err(ExecutionError::ParseError(format!("Unable to convert json to string: {}",err.to_string())))
    }
   }
}

pub fn create_and_write_bytes_to_file (data: &[u8], output_directory: &str, output_filename: &str) -> Result<(),String> {
    let mut out_file = match File::create(format!("{}/{}", output_directory, output_filename)) {
        Ok(val) => {
            val
        }
        Err(err) => {
            return  Err(err.to_string());
        }
    };

    match out_file.write_all(data) {
        Ok(_) => {
            Ok(())
        }
        Err(err) => {
            Err(err.to_string())
        }
    }
}