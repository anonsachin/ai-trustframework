#![cfg_attr(debug_assertions, allow(unused_imports))] // added to remove warnings on Deserialize, Serialize
use crate::errors::*;
use image::{open, DynamicImage, Luma, GenericImage, ImageBuffer};
use serde::{Deserialize, Serialize};
use std::fmt;

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

pub fn get_image_data(location: &str) -> Result<Vec<u8>, ExecutionError>{
  // opening the image from location
  let img = match open(location) {
    Ok(img) => {img},
    Err(err) => {
      return Err(ExecutionError::ImageExtractionError(format!("Error opening: {}", err.to_string())));
    }
  };
  // extracting bytes
  Ok(img.as_bytes().to_owned())
}

pub fn gray_png_image_from_bytes(data: Vec<u8>, height: u32, witdh: u32) -> Result<image::GrayImage, ExecutionError> {
  let mut grey_image: image::GrayImage = ImageBuffer::new(witdh,height);

  for pixel in grey_image.enumerate_pixels_mut() {
    let index: usize = pixel.0 as usize + pixel.1 as usize;
    *pixel.2 = match data.get(index){
      Some(data) => {
        Luma([*data])
      }
      None => {
        return Err(ExecutionError::ImageError(format!("Uanble to get data for image, array element not found.")));
      }
    };
  }

  Ok(grey_image)
}

pub fn generate_max_u8s(len: usize) -> Vec<u8> {
  let mut v: Vec<u8> = Vec::new();

  for _i in 1..=len {
    v.push(255);
  }

  v
}