
use std::fs::read_to_string;
use std::env::args;
use serde_json;
use ai_trustframework::inputs::*;
use ai_trustframework::errors::*;
use ai_trustframework::output::*;



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
                                    // Creating the output json
                                    let data = match output_json_string(6, &inputs.directory, "white.png") {
                                        Ok(data) => {
                                            data
                                        }
                                        Err(err) => {
                                            return Err(err);
                                        }
                                    };

                                    let white_data = generate_max_u8s(784);
                                    let white_gray_image = match gray_png_image_from_bytes(white_data,28,28) {
                                        Ok(img) => {
                                            img
                                        },
                                        Err(err) => {
                                            return  Err(err);
                                        }
                                    };

                                    // writing it out to file
                                    match create_and_write_bytes_to_file(data.as_bytes(), &inputs.directory, "output.json"){
                                        Ok(()) => {
                                            //save image to output directory and white.png
                                            match save_grey_image_to_location(white_gray_image, format!("{}/{}",inputs.directory,"white.png").as_str()){
                                                Ok(_) => {
                                                    Ok(())
                                                },
                                                Err(err) => {
                                                    return Err(err);
                                                }
                                            }
                                        }
                                        Err(err) => {
                                            Err(ExecutionError::FileError(err))
                                        }
                                    }
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




