use std::env::args;

fn main() -> Result<(), WasiError>  {
    // get the args
    let arg: Vec<String> = args().collect();

  match arg.get(0) {
    Some(v) => {
      // prints the frist arg to std out
      println!("{:?}",v);
    }
    None => {
      // if no args the following message goes to stderr
      return Err(WasiError::Args("There are no args".to_string()))
    }
  }

  Ok(())
}

#[derive(Debug)]
enum WasiError {
  Args(String)
}