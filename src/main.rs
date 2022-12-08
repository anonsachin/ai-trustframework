#[macro_use]
extern crate serde_derive;

use std::fs::read_to_string;
use std::env::args;
use serde::{Deserialize, Serialize};
use serde_json;

fn main() -> Result<(),ExecutionError> {
    match args().nth(1) {
        // Handling the comandline input
        Some(location) => {
            // Reading the file passed in through the command line
            match read_to_string(location) {
                // parse the file read in
                Ok(file) => {
                    // parsing the string
                    let json_input:Result<serde_json::Value, serde_json::Error> = serde_json::from_str(&file);

                    match json_input {
                        Ok(input) => {
                            match extract_inputs(input) {
                                Ok(inputs) => {
                                    println!("{:?}",inputs);
                                    Ok(())
                                }
                                Err(err) =>{ Err(err) }
                            }
                            
                        },
                        Err(err) => {
                            Err(ExecutionError::ParseError(err.to_string()))
                        }
                    }
                }
                Err(err) => {
                    Err(ExecutionError::InvalidArgError(err.to_string()))
                }
            }
        },
        None => {
            Err(ExecutionError::InvalidArgError("Args not provided.".to_string()))
        }
    }
}

// The error values
#[derive(Debug)]
enum ExecutionError {
    InvalidArgError(String),
    ParseError(String)
}


#[derive(Debug,Deserialize,Serialize)]
struct Inputs {
    image: String,
    directory: String
}

fn extract_inputs(values: serde_json::Value) -> Result<Inputs, ExecutionError> {
    // Initialize empty inputs
    let mut input = Inputs{
        directory: "".to_string(),
        image: "".to_string()
    };

    // Getting the image location
    let image = values["image"].to_owned();
    match image {
         serde_json::Value::String(location) => {
            input.image = location;
         }
        serde_json::Value::Null => {
          return  Err(ExecutionError::ParseError("Unable to get image feild.".to_string()));
        }
        _ => {
            return  Err(ExecutionError::ParseError("Unexpected value type for image feild, expecting string.".to_string()));
          }
    }

    // Getting the image location
    let dir = values["output_directory"].to_owned();
    match dir {
         serde_json::Value::String(location) => {
            input.directory = location;
         }
        serde_json::Value::Null => {
          return  Err(ExecutionError::ParseError("Unable to get output_directory feild.".to_string()));
        }
        _ => {
            return  Err(ExecutionError::ParseError("Unexpected value type for output_directory feild, expecting string.".to_string()));
          }
    }

    Ok(input)
}