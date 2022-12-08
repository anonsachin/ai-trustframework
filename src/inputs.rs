#![cfg_attr(debug_assertions, allow(unused_imports))] // added to remove warnings on Deserialize, Serialize
use crate::errors::*;
use serde::{Deserialize, Serialize};

#[derive(Debug,Deserialize,Serialize)]
pub struct Inputs {
    pub image: String,
    pub directory: String
}

pub fn extract_inputs(values: serde_json::Value) -> Result<Inputs, ExecutionError> {
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