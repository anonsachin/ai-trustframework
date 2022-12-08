#[macro_use]
extern crate serde_derive;

use std::fs::read_to_string;
use std::env::args;
use serde_json;
use crate::inputs::*;
use crate::errors::*;

pub mod inputs;
pub mod errors;

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




